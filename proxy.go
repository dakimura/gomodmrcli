package gomodmrcli

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/mod/modfile"
	"golang.org/x/xerrors"
)

type ProxyClient struct {
	hc *http.Client
}

func NewProxyClient(hc *http.Client) *ProxyClient {
	return &ProxyClient{
		hc: hc,
	}
}

func (c *ProxyClient) Mod(modulePath string, version string, disableModuleFetch bool) (*modfile.File, error) {
	// --- build URL
	// e.g. "https://proxy.golang.org/golang.org/x/text/@v/v0.3.2.mod"
	url := fmt.Sprintf("https://proxy.golang.org/%s/@v/%s.mod", modulePath, version)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create a http request to %s: %w", url, err)
	}

	// --- set Header
	if disableModuleFetch {
		req.Header.Set("Disable-Module-Fetch", "true")
	}

	// --- execute http Get
	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request")
	}
	defer resp.Body.Close()

	// parse http response
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read http response: %w", err)
	}

	mf, err := modfile.ParseLax("go.mod", r, nil)
	if err != nil {
		return nil, xerrors.Errorf("parse go.mod file: %w", err)
	}

	return mf, nil
}
