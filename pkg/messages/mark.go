package messages

import (
	"context"
	"encoding/json"
	"net/http"
)

type MarkRequest struct {
	Identifier string `json:"id"`
}

func DecodeMarkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request MarkRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}
