package api

import (
	"time"

	"github.com/durotimicodes/natwest-clone/user-service/config"
	"github.com/hashicorp/go-retryablehttp"
)

type Client struct {
	BaseURL string
	APIKey  string
	HTTP    *retryablehttp.Client
}

// NewClient initializes a new API Client
func NewClient(cfg *config.Config) *Client {
	client := retryablehttp.NewClient()
	client.RetryMax = 3
	client.RetryWaitMax = 5 * time.Second //Max wait time before retry
	client.RetryWaitMin = 1 * time.Second //Min wait time before retry

	return &Client{
		BaseURL: cfg.APIBaseURL,
		APIKey:  cfg.APIKey,
		HTTP:    client,
	}
}
