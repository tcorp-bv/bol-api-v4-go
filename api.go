package bolapiv4

import (
	"context"
	"net/http"
	"time"

	"github.com/tcorp-bv/bol-api-v4-go/client"
	"github.com/tcorp-bv/bol-api-v4-go/http/backoff"
	"github.com/tcorp-bv/bol-api-v4-go/http/middleware"

	"github.com/tcorp-bv/bol-api-v4-go/auth"
)

const (
	defaultHost  = "api.bol.com"
	startBackoff = 1 * time.Second
	maxBackoff   = 20 * time.Second
	maxTries     = 10
)

var (
	// Error codes for which API requests should be retried.
	// Note that 429 will be retried after the period indicated in the header, the rest will retry according to the backoff.
	retryCodes = map[int]bool{
		429: true, // rate limit reached
		500: true, // internal server error
		504: true, // gateway timeout error
		408: true, // request timeout error
	}
)

// BolAPI is the core interface to retrieve the bol.com v3 client.
// The client contains middleware to retry on rate limits and backoff on server timeouts.
// Requests can thus sometimes take multiple minutes to complete.
type BolAPI interface {
	// GetClient gets the bol.com API client
	GetClient() *client.APIClient
}

// Wrapper implemmentation of the BolAPI interface
type bolAPI struct {
	HTTPClient *http.Client
}

// GetClient gets the bol.com API client
func (api *bolAPI) GetClient() *client.APIClient {
	cfg := client.NewConfiguration()
	cfg.HTTPClient = api.HTTPClient

	// transport.Consumers["application/vnd.retailer.v3+json"] = runtime.JSONConsumer()
	// transport.Consumers["application/problem+json"] = runtime.JSONConsumer()
	// transport.Producers["application/vnd.retailer.v3+json"] = runtime.JSONProducer()
	return client.NewAPIClient(cfg)
}

// NewWithHost creates a new API client with the host (default is "api.bol.com")
func NewWithHost(provider auth.CredentialProvider, host string, basepath string) (BolAPI, error) {
	return newBolAPI(provider, host, basepath, nil)
}

// NewWithTransport creates a new API client with the transportProvider as injected middleware
func NewWithTransport(provider auth.CredentialProvider, host string, basepath string, transportProvider func(tripper http.RoundTripper) http.RoundTripper) (BolAPI, error) {
	// create the original BolAPI
	bolAPI, err := newBolAPI(provider, host, basepath, transportProvider)
	if err != nil {
		return nil, err
	}
	// wrap the transport
	bolAPI.HTTPClient.Transport = transportProvider(bolAPI.HTTPClient.Transport)
	return bolAPI, nil
}

func newBolAPI(provider auth.CredentialProvider, host string, basepath string, transportProvider func(tripper http.RoundTripper) http.RoundTripper) (*bolAPI, error) {
	if host == "" {
		host = defaultHost
	}
	bolAuth, err := auth.New(provider)
	if err != nil {
		return nil, err
	}
	// Gets an http client with the oauth automatically added as middleware
	client := bolAuth.Client(context.Background())
	// Add the retrying middleware
	client.Transport = middleware.Middleware{
		MaxTries:   maxTries,
		Backoff:    backoff.NewExponentialBackoff(startBackoff, maxBackoff),
		Transport:  client.Transport,
		RetryCodes: retryCodes,
	}
	// Add the client middleware
	if transportProvider != nil {
		client.Transport = transportProvider(client.Transport)
	}

	return &bolAPI{HTTPClient: client}, nil
}

// New creates a new API client with the default host (default is "api.bol.com") and the provided CredentialProvider.
func New(provider auth.CredentialProvider) (BolAPI, error) {
	return NewWithHost(provider, defaultHost, "")
}
