package aoj


// AOJ Client
type AOJClient struct {
	username string
	password string
}

func NewAPIClient() *AOJClient {
	return &AOJClient{
	}
}

func (client *AOJClient) Login(username string, password string) {
	client.username = username
	client.password = password
	// authenticate
}

func (client *AOJClient) FetchProblem(username string, password string) {

}

// Fetch test cases of the problem
func (client *AOJClient) FetchTestCases(problemId string) (string, error) {

}

func (client *AOJClient) Submit(problemId string, lang string, code string) {
}
