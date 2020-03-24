package client

import (
	"context"
	"github.com/travelist/aoj-cli/client/request"
	"github.com/travelist/aoj-cli/client/response"
	"net/http"
)

func (client *AOJClient) Submit(ctx context.Context, body request.SubmitRequest) (
	*response.SubmitResponse, error) {

	path := "/submissions"
	jsonBody, e := encodeBody(body)
	if e != nil {
		return nil, e
	}

	req, e := client.newAPIRequest(ctx, http.MethodPost, path, jsonBody)
	if e != nil {
		return nil, e
	}

	var res response.SubmitResponse
	if e := client.send(req, &res); e != nil {
		return nil, e
	}
	return &res, e
}
