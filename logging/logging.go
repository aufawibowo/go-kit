package logging

import (
	"github.com/aufawibowo/go-kit/service"
	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.StringService
}
