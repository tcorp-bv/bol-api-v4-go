package bolapiv4

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/tcorp-bv/bol-api-v4-go/auth"
)

const (
	demoURL      = "api.bol.com"
	demoBasePath = "/retailer-demo/"
)

// GetTestAPI returns an integration testing instance (thus connected to the actual API) with
// credentials from the environment (BOL_CLIENT_ID and BOL_CLIENT_SECRET).
// If production is set to false, the demon environment will be used.
func GetTestAPI(t *testing.T, verbose bool, production bool) *BolAPI {
	provider := auth.EnvironmentProvider{ClientIDKey: "BOL_CLIENT_ID", ClientSecretKey: "BOL_CLIENT_SECRET"}
	path := demoBasePath
	if production {
		path = ""
	}
	bolAPI, err := NewWithTransport(&provider, demoURL, path, func(t http.RoundTripper) http.RoundTripper {
		return URLRewriter{Proxied: t, Verbose: verbose, Production: production}
	})
	if err != nil {
		t.Fatal(err)
	}
	return &bolAPI
}

// URLRewriter rewrites the url to connect to the demo environment
type URLRewriter struct {
	Proxied    http.RoundTripper
	Production bool
	Verbose    bool
}

// RoundTrip is the middleware implementation to rewrite the url to connect to the demo environment
func (lrt URLRewriter) RoundTrip(req *http.Request) (res *http.Response, err error) {
	if !lrt.Production && req.URL.Path[:15] == "/retailer-demo/" { //fix demo url
		req.URL.Path = strings.Replace(req.URL.Path, "/retailer-demo/retailer/", "/retailer-demo/", 1)
	}
	// Send the request, get the response (or the error)
	res, err = lrt.Proxied.RoundTrip(req)
	//print the response
	if lrt.Verbose {
		buf := bytes.Buffer{}
		buf.ReadFrom(res.Body)
		res.Body.Close()
		res.Body = ioutil.NopCloser(bytes.NewBufferString(buf.String()))
		fmt.Printf("%s (%d): %s\n", req.URL.String(), res.StatusCode, buf.String())
	}
	return res, err
}
