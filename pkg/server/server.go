package server

import (
	"context"
	"net/http"

	"github.com/ereOn/gotyouthere/pkg/messages"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

type HTTPHandler struct {
	mux.Router
	logger  log.Logger
	service Service
}

func NewHTTPHandler(service Service, logger log.Logger) http.Handler {
	handler := &HTTPHandler{
		Router:  *mux.NewRouter(),
		logger:  logger,
		service: service,
	}

	ctx := context.Background()

	handler.Methods("POST").Path("/mark").Handler(httptransport.NewServer(
		ctx,
		makeMarkEndpoint(service),
		messages.DecodeMarkRequest,
		messages.EncodeGenericResponse,
	))

	return handler
}
