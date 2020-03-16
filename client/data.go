package client

import (
	"context"
	"fmt"
	"github.com/travelist/aoj-cli/client/response"
	"net/http"
)

// Data API

// Fetch test cases of the problem
func (client *AOJClient) findByProblemIdSamples(ctx context.Context, problemId string) (
	*response.TestCaseSampleListResponse, error) {

	path := fmt.Sprintf("/testcases/samples/%s", problemId)
	request, e := client.newRequest(ctx, http.MethodGet, path, nil)

	if e != nil {
		return nil, e
	}

	var res response.TestCaseSampleListResponse
	e = client.send(request, &res)
	return &res, e
}

func (client *AOJClient) findByProblemIdTestcaseHeader(ctx context.Context, problemId string) (
	*response.TestCaseHeaderResponse, error) {

	path := fmt.Sprintf("/testcases/%s/header", problemId)
	request, e := client.newRequest(ctx, http.MethodGet, path, nil)
	if e != nil {
		return nil, e
	}

	var res response.TestCaseHeaderResponse
	e = client.send(request, &res)
	return &res, e
}
