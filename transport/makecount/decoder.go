package makecount

import (
	"context"
	"encoding/json"
	"net/http"
)

func Decoder(_ context.Context, r *http.Request) func() (interface{}, error) {
	return func() (interface{}, error) {
		var request countRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, err
		}
		return request, nil
	}
}
