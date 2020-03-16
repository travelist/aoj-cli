package request

type SubmitRequest struct {
	ProblemID  string `json:"problemId"`
	Language   string `json:"language"`
	SourceCode string `json:"sourceCode"`
}
