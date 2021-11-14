package transport

import (
	"context"
	"encoding/json"
	"net/http"
)

func Encode(_ context.Context, w http.ResponseWriter, response interface{}) func() error {
	return func() error{
		return json.NewEncoder(w).Encode(response)
	}

}
