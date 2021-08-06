package gomodmrcli

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type IndexClient struct {
	hc *http.Client
}

func NewIndexClient(hc *http.Client) *IndexClient {
	return &IndexClient{
		hc: hc,
	}
}

type Index struct {
	Path      string    `json:"Path"`
	Version   string    `json:"Version"`
	Timestamp time.Time `json:"Timestamp"`
}

func (c *IndexClient) Index(since time.Time, limit int, disableModuleFetch bool) ([]Index, error) {
	const layout="2006-01-02T15:04:05Z"
	if limit < 0 || limit > 2000 {
		return nil, errors.New("limit must be between 0 and 2000")
	}

	// --- build URL
	url := "https://index.golang.org/index?"
	var queries []string
	if !since.IsZero() {
		queries = append(queries, fmt.Sprintf("since=%s", since.Format(layout)))
	}
	if limit != 0 {
		queries = append(queries, fmt.Sprintf("limit=%d", limit))
	}
	url += strings.Join(queries, "&")

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
	var ret []Index
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		idx, err := parseIndexLine(scanner.Bytes())
		if err != nil {
			return nil, err
		}
		ret = append(ret, *idx)
	}

	return ret, nil
}

func parseIndexLine(line []byte) (*Index, error) {
	var idx Index
	if err := json.Unmarshal(line, &idx); err != nil {
		return nil, fmt.Errorf("parse line as json. line=%s: %w", line, err)
	}
	return &idx, nil
}
