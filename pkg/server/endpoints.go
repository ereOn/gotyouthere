package server

import (
	"context"

	"github.com/ereOn/gotyouthere/pkg/messages"
	"github.com/go-kit/kit/endpoint"
)

func makeMarkEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(messages.MarkRequest)
		err := svc.Mark(req.Identifier)

		return nil, err
	}
}
