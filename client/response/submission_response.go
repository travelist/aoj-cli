package response

type SubmissionRecordResponse struct {
	JudgeID        int         `json:"judgeId"`
	JudgeType      int         `json:"judgeType"`
	UserID         string      `json:"userId"`
	ProblemID      string      `json:"problemId"`
	SubmissionDate int64       `json:"submissionDate"`
	Language       string      `json:"language"`
	Status         int         `json:"status"`
	CPUTime        int         `json:"cpuTime"`
	Memory         int         `json:"memory"`
	CodeSize       int         `json:"codeSize"`
	Accuracy       string      `json:"accuracy"`
	JudgeDate      int64       `json:"judgeDate"`
	Score          interface{} `json:"score"`
	ProblemTitle   string      `json:"problemTitle"`
	Token          string      `json:"token"`
}

type SubmissionRecordsResponse []SubmissionRecordResponse
