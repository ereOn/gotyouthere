package server

import (
	"time"

	"github.com/go-kit/kit/log"
	"gopkg.in/redis.v5"
)

const (
	presenceKey string = "presence"
)

type Service interface {
	Mark(identifier string) error
}

type service struct {
	redisClient  *redis.Client
	redisScripts map[string]string
	logger       log.Logger
}

func NewService(redisClient *redis.Client, logger log.Logger) (Service, error) {
	service := service{
		redisClient:  redisClient,
		redisScripts: make(map[string]string),
		logger:       logger,
	}

	{
		sha, err := redisClient.ScriptLoad(LockScript).Result()

		if err != nil {
			return nil, err
		}

		service.redisScripts["lock"] = sha
	}
	{
		sha, err := redisClient.ScriptLoad(UnlockScript).Result()

		if err != nil {
			return nil, err
		}

		service.redisScripts["unlock"] = sha
	}

	return service, nil
}

func (s service) Lock(timeout time.Duration, identifier string) error {
	s.logger.Log("event", "lock", "timeout", timeout, "identifier", identifier)

	keys := []string{
		"presence",
		"lock",
		"waits",
		"unblocked_waits",
	}

	pipeline := s.redisClient.Pipeline()
	pipeline.EvalSha(s.redisScripts["lock"], keys, identifier, "1").Result()
	pipeline.BLPop(timeout, "unblocked_waits")
	_, err := pipeline.Exec()

	s.logger.Log("event", "lock", "timeout", timeout, "identifier", identifier, "error", err)

	return err
}

func (s service) Unlock(identifier string) error {
	s.logger.Log("event", "unlock", "identifier", identifier)

	keys := []string{
		"presence",
		"lock",
		"waits",
		"unblocked_waits",
	}

	_, err := s.redisClient.EvalSha(s.redisScripts["unlock"], keys, identifier, "1").Result()

	s.logger.Log("event", "unlock", "identifier", identifier, "error", err)

	return err
}

func (s service) Mark(identifier string) error {
	s.Lock(time.Second*30, identifier)
	s.logger.Log("event", "mark", "identifier", identifier, "action", "waiting...")
	time.Sleep(time.Second * 5)
	s.logger.Log("event", "mark", "identifier", identifier, "action", "waited")
	s.Unlock(identifier)

	return nil
}
