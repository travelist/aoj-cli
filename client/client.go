package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"path"
)

type AOJClient struct {
	// Base API endpoint URL
	APIEndpointURL *url.URL
	// Data API endpoint URL
	DataAPIEndpointURL *url.URL

	// Authentication token
	Token string
	// client
	httpClient *http.Client
}

func NewClient(
	baseEndpointURL string,
	dataEndpointUrl string,
	httpClient *http.Client,
	token string) (*AOJClient, error) {

	parsedBaseURL, e := url.ParseRequestURI(baseEndpointURL)
	if e != nil {
		return nil, fmt.Errorf("failed to parse url: %s\n", baseEndpointURL)
	}
	parsedDataURL, e := url.ParseRequestURI(dataEndpointUrl)
	if e != nil {
		return nil, fmt.Errorf("failed to parse url: %s\n", dataEndpointUrl)
	}

	client := &AOJClient{
		APIEndpointURL:     parsedBaseURL,
		DataAPIEndpointURL: parsedDataURL,
		httpClient:         httpClient,
		Token:              token,
	}

	return client, nil
}

// --- private utilities ---

func (client *AOJClient) newAPIRequest(
	ctx context.Context,
	method string,
	urlPath string,
	body io.Reader) (*http.Request, error) {

	endpointURL := *client.APIEndpointURL
	endpointURL.Path = path.Join(client.APIEndpointURL.Path, urlPath)

	req, e := http.NewRequest(method, endpointURL.String(), body)
	if e != nil {
		return nil, e
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.AddCookie(&http.Cookie{Name: "JSESSIONID", Value: client.Token})
	return req, nil
}

func (client *AOJClient) newDataRequest(
	ctx context.Context,
	method string,
	urlPath string,
	body io.Reader) (*http.Request, error) {

	endpointURL := *client.DataAPIEndpointURL
	endpointURL.Path = path.Join(client.DataAPIEndpointURL.Path, urlPath)

	req, e := http.NewRequest(method, endpointURL.String(), body)
	if e != nil {
		return nil, e
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.AddCookie(&http.Cookie{Name: "JSESSIONID", Value: client.Token})
	return req, nil
}

func (client *AOJClient) send(request *http.Request, result interface{}) error {
	res, e := client.httpClient.Do(request)
	if e != nil {
		return e
	}

	if !IsSuccess(res.StatusCode) {
		return newAOJClientError(res)
	}

	if e := decodeBody(res, result); e != nil {
		return errors.Wrap(e, "cannot parse response body")
	}

	return nil
}

func encodeBody(in interface{}) (io.Reader, error) {
	body, e := json.Marshal(in)
	if e != nil {
		return nil, e
	}

	return bytes.NewReader(body), nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func IsSuccess(code int) bool {
	return 200 <= code && code < 300
}
