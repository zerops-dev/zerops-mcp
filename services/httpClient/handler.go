package httpClient

import (
	"net/http"

	"github.com/zerops-dev/di/di/s"
)

func DiScope() *s.XScope {
	return s.Scope(
		s.Service(New),
		s.Config(NewConfig, "httpClient"),
	)
}

func New(config Config) *http.Client {
	return &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   config.Timeout,
	}
}
