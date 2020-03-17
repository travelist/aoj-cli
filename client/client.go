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
	EndpointURL *url.URL
	HTTPClient  *http.Client
	username    string
	password    string
}

func NewClient(endpointURL string, httpClient *http.Client) (*AOJClient, error) {
	parsedURL, e := url.ParseRequestURI(endpointURL)

	if e != nil {
		return nil, fmt.Errorf("failed to parse url: %s\n", endpointURL)
	}

	client := &AOJClient{
		EndpointURL: parsedURL,
		HTTPClient:  httpClient,
	}

	return client, nil
}

// --- private utilities ---

func (client *AOJClient) newRequest(
	ctx context.Context,
	method string,
	urlPath string,
	body io.Reader) (*http.Request, error) {

	endpointURL := *client.EndpointURL
	endpointURL.Path = path.Join(client.EndpointURL.Path, urlPath)

	req, e := http.NewRequest(method, endpointURL.String(), body)
	if e != nil {
		return nil, e
	}

	req = req.WithContext(ctx)
	return req, nil
}

func (client *AOJClient) send(request *http.Request, result interface{}) error {
	res, e := client.HTTPClient.Do(request)
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
