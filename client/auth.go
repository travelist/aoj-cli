package client

import (
	"context"
	"github.com/travelist/aoj-cli/client/request"
	"github.com/travelist/aoj-cli/client/response"
	"net/http"
)

// Login
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

	var res response.LoginResponse
	if e := client.send(req, &res); e != nil {
		return nil, e
	}
	return &res, e
}

// Fetch current session information
// This method returns INVALID_REFRESH_TOKEN_ERROR when the current session is expired
func (client *AOJClient) Session(ctx context.Context) (
	*response.SessionResponse, error) {
	path := "/session"

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
