package main

import (
	"flag"
	"net/http"
	"os"

	redis "gopkg.in/redis.v5"

	"github.com/go-kit/kit/log"

	"github.com/ereOn/gotyouthere/pkg/server"
)

func main() {
	endpoint := flag.String("endpoint", ":9999", "The endpoint to listen on.")

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.NewContext(logger).With("timestamp", log.DefaultTimestampUTC)

	redisClientOptions, err := redis.ParseURL("redis://localhost")

	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(redisClientOptions)
	service, err := server.NewService(redisClient, logger)

	if err != nil {
		panic(err)
	}

	server := server.NewHTTPHandler(service, logger)

	// Start the application and block.
	logger.Log("event", "application started", "endpoint", *endpoint)
	err = http.ListenAndServe(*endpoint, server)

	if err != nil {
		panic(err)
	}
}
