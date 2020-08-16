package auth

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// Make sure that the file provider can actually read the two files
func TestFileProvider(t *testing.T) {
	clientID := "oRNWbHFXtAECmhnZmEndcjLIaSKbRMVE"
	clientSecret := "aQHPOnmYkPZNgeRziPnQyyOJYytUbcFBVJBvbMKoDdpPqaZbaOiLUTWzPAkpPsZFZbJHrcoltdgpZolyNcgvvBaKcmkqFjucFzXhDONTsPAtHHyccQlLUZpkOuywMiOycDWcCySFsgpDiyGnCWCZJkNTtVdPxbSUTWVIFQiUxaPDYDXRQAVVTbSVZArAZkaLDLOoOvPzxSdhnkkJWzlQDkqsXNKfAIgAldrmyfROSyCGMCfvzdQdUQEaYZTPEoA"

	// Setup temp dir
	dirName, err := ioutil.TempDir(os.TempDir(), "test-")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		_ = os.Remove(dirName)
	}()
	// Setup the files
	ioutil.WriteFile(filepath.Join(dirName, "clientId"), []byte(clientID), 0644)
	ioutil.WriteFile(filepath.Join(dirName, "clientSecret"), []byte(clientSecret), 0644)

	fp := FileProvider{Dir: dirName, ClientIDFile: "clientId", ClientSecretFile: "clientSecret"}
	testProvider(t, fp, clientID, clientSecret)
}

func testProvider(t *testing.T, provider FileProvider, clientID string, clientSecret string) {
	if cID, err := provider.ClientID(); err != nil || cID != clientID {
		t.Error("File provider does not read clientId file properly. ", err)
	}
	if cSecret, err := provider.ClientID(); err != nil || cSecret != clientID {
		t.Error("File provider does not read clientId file properly. ", err)
	}
}
