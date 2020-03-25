package client

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/travelist/aoj-cli/client/request"
	"github.com/travelist/aoj-cli/client/response"
	"net/http"
)

// Check token validity
func (client *AOJClient) IsAuthenticated() bool {
	if len(client.Token) == 0 {
		return false
	}

	if _, e := client.Session(context.Background()); e != nil {
		fmt.Printf("%v\n", e)
		return false
	}

	return true
}

// Login
// This function updates current client session
func (client *AOJClient) Login(ctx context.Context, body request.LoginRequest) (
	*response.LoginResponse, error) {

	path := "/session"
	jsonBody, e := encodeBody(body)
	if e != nil {
		return nil, e
	}

	req, e := client.newAPIRequest(ctx, http.MethodPost, path, jsonBody)
	if e != nil {
		return nil, e
	}

	res, e := client.httpClient.Do(req)
	if e != nil {
		return nil, e
	}

	if !IsSuccess(res.StatusCode) {
		return nil, newAOJClientError(res)
	}

	for _, v := range res.Cookies() {
		if v.Name == "JSESSIONID" {
			client.Token = v.Value
			break
		}
	}

	var result response.LoginResponse
	if e := decodeBody(res, &result); e != nil {
		return nil, errors.Wrap(e, "cannot parse response body")
	}

	return &result, e
}

// Fetch current session information
// This method returns INVALID_REFRESH_TOKEN_ERROR when the current session is expired
func (client *AOJClient) Session(ctx context.Context) (
	*response.SessionResponse, error) {
	path := "/self"

	req, e := client.newAPIRequest(ctx, http.MethodGet, path, nil)
	if e != nil {
		return nil, e
	}

	var res response.SessionResponse
	if e := client.send(req, &res); e != nil {
		return nil, e
	}

	return &res, e
}
