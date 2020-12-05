package jenkinsapi

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Jenkins struct {
	Client   *http.Client
	Endpoint string
	User     string
	ApiToken string
}

func New(endpoint, user, apiToken string) (*Jenkins, error) {
	config := getDefaultConf()
	transport := &http.Transport{
		MaxIdleConns:          config.HTTPMaxConns.MaxIdleConns,
		MaxIdleConnsPerHost:   config.HTTPMaxConns.MaxIdleConnsPerHost,
		IdleConnTimeout:       config.HTTPTimeout.IdleConnTimeout,
		ResponseHeaderTimeout: config.HTTPTimeout.HeaderTimeout,
	}
	return &Jenkins{
		Client:   &http.Client{Transport: transport},
		Endpoint: endpoint,
		User:     user,
		ApiToken: apiToken,
	}, nil
}

func (j *Jenkins) Do(path string, body []byte) (*http.Response, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/%s", j.Endpoint, path))
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, uri.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, nil
	}
	req.SetBasicAuth(j.User, j.ApiToken)
	date := time.Now().UTC().Format(http.TimeFormat)
	req.Header.Set(HTTPHeaderDate, date)
	req.Header.Set(HTTPHeaderHost, j.Endpoint)
	req.Header.Set(HTTPHeaderUserAgent, userAgent())
	return j.Client.Do(req)
}
