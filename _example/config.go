package main

import (
	"net/url"
	"strings"
)

// RunConfig contains the existing command-line credentials and endpoint selection.
type RunConfig struct {
	ApiKey     string
	ApiSecret  string
	BaseUrl    string
	UseTestNet bool
}

// NewRunConfig normalizes the endpoint without dereferencing an omitted host.
func NewRunConfig(apiKey string, apiSecret string, hostUsed *string) (*RunConfig, error) {
	config := &RunConfig{
		ApiKey:     apiKey,
		ApiSecret:  apiSecret,
		UseTestNet: false,
	}
	if hostUsed != nil {
		config.BaseUrl = *hostUsed
	}
	if config.BaseUrl == "" {
		config.BaseUrl = "https://api.gateio.ws/api/v4"
	}
	if !strings.HasPrefix(config.BaseUrl, "http") {
		config.BaseUrl = "https://" + config.BaseUrl
	}
	if !strings.HasSuffix(config.BaseUrl, "/api/v4") {
		config.BaseUrl += "/api/v4"
	}
	parsedUrl, err := url.Parse(config.BaseUrl)
	if err != nil {
		return nil, err
	}
	if parsedUrl.Host == "fx-api-testnet.gateio.ws" {
		config.UseTestNet = true
	}
	return config, nil
}
