package client

import (
	"context"
	"fmt"
	"github.com/travelist/aoj-cli/client/response"
	"net/http"
)

// Data API

// Fetch test cases of the problem
func (client *AOJClient) FindByProblemIdSamples(ctx context.Context, problemId string) (
	response.TestCaseSampleListResponse, error) {

	path := fmt.Sprintf("/testcases/samples/%s", problemId)
	req, e := client.newDataRequest(ctx, http.MethodGet, path, nil)
	if e != nil {
		return nil, e
	}

	var result response.TestCaseSampleListResponse
	if e := client.send(req, &result); e != nil {
		return nil, e
	}

	return result, nil
}

func (client *AOJClient) FindByProblemIdTestcaseHeader(ctx context.Context, problemId string) (
	*response.TestCaseHeaderResponse, error) {

	path := fmt.Sprintf("/testcases/%s/header", problemId)
	req, e := client.newDataRequest(ctx, http.MethodGet, path, nil)
	if e != nil {
		return nil, e
	}

	var res response.TestCaseHeaderResponse
	if e := client.send(req, &res); e != nil {
		return nil, e
	}
	return &res, e
}
