package auth

import "testing"

const (
	testClientID     = "oRNWbHFXtAECmhnZmEndcjLIaSKbRMVE"
	testClientSecret = "aQHPOnmYkPZNgeRziPnQyyOJYytUbcFBVJBvbMKoDdpPqaZbaOiLUTWzPAkpPsZFZbJHrcoltdgpZolyNcgvvBaKcmkqFjucFzXhDONTsPAtHHyccQlLUZpkOuywMiOycDWcCySFsgpDiyGnCWCZJkNTtVdPxbSUTWVIFQiUxaPDYDXRQAVVTbSVZArAZkaLDLOoOvPzxSdhnkkJWzlQDkqsXNKfAIgAldrmyfROSyCGMCfvzdQdUQEaYZTPEoA"
)

func TestDefaultEndpoint(t *testing.T) {
	if defaultTokenEndpoint() != "https://login.bol.com/token" {
		t.Errorf("Expected https://login.bol.com, got %s as token endpoint", defaultTokenEndpoint())
	}
}

func TestGetConfig(t *testing.T) {
	c, err := getConfig(&BasicProvider{ClientIDVal: testClientID, ClientSecretVal: testClientSecret})
	if err != nil {
		t.Error(err)
	}
	if c.ClientID != testClientID || c.ClientSecret != testClientSecret || c.TokenURL != defaultTokenEndpoint() {
		t.Error("Client credentials not properly parsed")
	}
}
