package messages

import (
	"context"
	"encoding/json"
	"net/http"
)

func EncodeGenericResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
