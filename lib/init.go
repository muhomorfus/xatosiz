package tracing

import (
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/tracing/openapi"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var api *openapi.APIClient
var serviceName string
var logger *zap.SugaredLogger

type Config struct {
	Server      string
	ServiceName string
	Logger      *zap.SugaredLogger
}

func Init(c *Config) {
	api = openapi.NewAPIClient(&openapi.Configuration{
		Servers: []openapi.ServerConfiguration{
			{URL: c.Server},
		},
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	})

	serviceName = c.ServiceName

	if c.Logger != nil {
		logger = c.Logger.WithOptions(zap.Fields(zap.Bool("xatosiz", true)))
	} else {
		logger = zap.NewNop().Sugar()
	}
}
