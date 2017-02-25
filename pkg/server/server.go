package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Handler interface {
	http.Handler
}

type HTTPHandler struct {
	http.ServeMux
	logger log.Logger
}

func endpoint(ctx context.Context, request interface{}) (interface{}, error) {
	//req := request.(uppercaseRequest)
	//v, err := svc.Uppercase(req.S)
	//if err != nil {
	//	return uppercaseResponse{v, err.Error()}, nil
	//}
	//return uppercaseResponse{v, ""}, nil
	return nil, nil
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	//var request uppercaseRequest
	//if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//	return nil, err
	//}
	//return request, nil
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func NewHTTPHandler(logger log.Logger) *HTTPHandler {
	handler := &HTTPHandler{
		ServeMux: *http.NewServeMux(),
		logger:   logger,
	}

	ctx := context.Background()

	handler.Handle("/mark", httptransport.NewServer(
		ctx,
		endpoint,
		decodeRequest,
		encodeResponse,
	))

	return handler
}
