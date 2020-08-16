package auth

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	defaultDir              = "/etc/secrets"
	defaultClientIDFile     = "clientId"
	defaultClientSecretFile = "clientSecret"
)

// CredentialProvider provides the OAuth2.0 clientID and secret for the bol.com account.
type CredentialProvider interface {
	// ClientID returns the OAuth 2.0 ID for the bol.com retailer
	ClientID() (string, error)
	// ClientSecret return the OAuth 2.0 secret for the bol.com retailer
	ClientSecret() (string, error)
}

// FileProvider is a CredentialProvider implementation for a dir containing two files, each containing the secret as a string.
type FileProvider struct {
	// The directory with the files, this is usually the place where you mount the secret volume
	Dir string // Should equal /etc/secrets

	// The filename within Dir containing the ClientId
	ClientIDFile string // Should equal "clientId"

	// The filename within Dir containing the ClientSecret
	ClientSecretFile string // Should equal "clientSecret"
}

// NewFileProvider returns a new FileProvider with that reads the credentials from "/etc/secrets/clientId" and "/etc/secrets/clientSecret"
func NewFileProvider() CredentialProvider {
	return &FileProvider{Dir: defaultDir, ClientIDFile: defaultClientIDFile, ClientSecretFile: defaultClientSecretFile}
}

// ClientID returns the OAuth 2.0 ID for the bol.com retailer, read from the configured file
func (p *FileProvider) ClientID() (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(p.Dir, p.ClientIDFile))
	return string(data), err
}

// ClientSecret return the OAuth 2.0 secret for the bol.com retailer, read from the configured file
func (p *FileProvider) ClientSecret() (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(p.Dir, p.ClientSecretFile))
	return string(data), err
}

// BasicProvider is a CredentialProvider implementation for hardcoded credentials
type BasicProvider struct {
	// The OAuth 2.0 clientID for the bol.com retailer
	ClientIDVal string
	// The OAuth 2.0 secret for the bol.com retailer
	ClientSecretVal string
}

// ClientID returns the configured OAuth 2.0 ID for the bol.com retailer (ClientIDVal)
func (p *BasicProvider) ClientID() (string, error) {
	return p.ClientIDVal, nil
}

// ClientSecret return the configured OAuth 2.0 secret for the bol.com retailer (ClientSecretVal)
func (p *BasicProvider) ClientSecret() (string, error) {
	return p.ClientSecretVal, nil
}

// EnvironmentProvider is a CredentialProvider implementation that fetches the OAuth 2.0 credentials from the configured environment variables.
type EnvironmentProvider struct {
	// ClientIDKey contains the environment variable name for the OAuth 2.0 clientID for the bol.com retailer
	ClientIDKey string
	// ClientSecretKey contains the environment variable name for the OAuth 2.0 secret for the bol.com retailer
	ClientSecretKey string
}

// ClientID returns the OAuth 2.0 ID for the bol.com retailer, read from the environment
func (p *EnvironmentProvider) ClientID() (string, error) {
	clientID, found := os.LookupEnv(p.ClientIDKey)
	if !found {
		return "", fmt.Errorf("environment variable '%s' not found", p.ClientIDKey)
	}
	return clientID, nil
}

// ClientSecret return the OAuth 2.0 secret for the bol.com retailer, read from the environment
func (p *EnvironmentProvider) ClientSecret() (string, error) {
	clientSecret, present := os.LookupEnv(p.ClientSecretKey)
	if !present {
		return "", fmt.Errorf("environment variable '%s' not found", p.ClientSecretKey)
	}
	return clientSecret, nil
}
