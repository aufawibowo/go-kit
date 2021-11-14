package makecount

import (
	"context"
	"github.com/aufawibowo/go-kit/service"
	"github.com/go-kit/kit/endpoint"
)

func Endpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}
