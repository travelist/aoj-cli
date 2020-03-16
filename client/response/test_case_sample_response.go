package response

type TestCaseSampleResponse struct {
	ProblemID string `json:"problemId"`
	Serial    int    `json:"serial"`
	In        string `json:"in"`
	Out       string `json:"out"`
}

type TestCaseSampleListResponse []TestCaseSampleResponse
