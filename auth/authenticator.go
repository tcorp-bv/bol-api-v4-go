package auth

import (
	"context"

	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	defaultAuthURL = "https://login.bol.com"
	tokenEndpoint  = "token"
)

// defaultTokenEndpoint returns the default OAuth 2.0 token endpoint: 'https://login.bol.com/token'.
func defaultTokenEndpoint() string {
	return fmt.Sprintf("%s/%s", defaultAuthURL, tokenEndpoint)
}

// Authenticator retrieves OAuth 2.0 tokens for the bol API.
type Authenticator interface {
	// Token returns the OAuth 2.0 token.
	Token() (*oauth2.Token, error)
	// Client generates an http client that automatically contains the appropriate OAuth2.0 token.
	Client(ctx context.Context) *http.Client
}

// New creates a new instance of the Authenticator using the provider for the credentials.
func New(provider CredentialProvider) (Authenticator, error) {
	config, err := getConfig(provider)
	if err != nil {
		return nil, err
	}
	return &authenticator{config: config}, nil
}

// authenticator in an Authenticator implementation using golang's oauth2 client.
type authenticator struct {
	// The OAuth2.0 credentials
	config clientcredentials.Config
}

// Token Fetches the OAuth 2.0 token using the clientcredentials.
func (a *authenticator) Token() (*oauth2.Token, error) {
	return a.config.Token(context.Background())
}

// Client generates an http client that automatically contains the appropriate OAuth2.0 token.
func (a *authenticator) Client(ctx context.Context) *http.Client {
	return a.config.Client(ctx)
}

// getConfig loads the credentials from the CredentialProvider into a golang oauth2 Config.
func getConfig(provider CredentialProvider) (clientcredentials.Config, error) {
	// Get the credentials from the provider
	clientID, err := provider.ClientID()
	if err != nil {
		return clientcredentials.Config{}, err
	}

	clientSecret, err := provider.ClientSecret()
	if err != nil {
		return clientcredentials.Config{}, err
	}

	// Setup the config
	return clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     defaultTokenEndpoint(),
	}, nil
}
