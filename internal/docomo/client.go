package docomo

import (
	"docomo-bike/internal/libs/logger"

	"github.com/gojektech/heimdall/httpclient"
)

func NewScrappingClient(logger *logger.Logger) *ScrappingClient {
	return &ScrappingClient{
		HTTPClient: httpclient.NewClient(),
		Logger:     logger,
	}
}

type Client interface {
	Login(id string, password string) (string, error)
}

type ScrappingClient struct {
	HTTPClient *httpclient.Client
	Logger     *logger.Logger
}
