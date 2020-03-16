package response

type LoginResponse struct {
	ID                         string      `json:"id"`
	Name                       string      `json:"name"`
	Affiliation                string      `json:"affiliation"`
	RegisterDate               int64       `json:"registerDate"`
	LastSubmitDate             int64       `json:"lastSubmitDate"`
	Policy                     string      `json:"policy"`
	Country                    string      `json:"country"`
	BirthYear                  interface{} `json:"birthYear"`
	DisplayLanguage            string      `json:"displayLanguage"`
	DefaultProgrammingLanguage string      `json:"defaultProgrammingLanguage"`
	Status                     struct {
		Submissions  int `json:"submissions"`
		Solved       int `json:"solved"`
		Accepted     int `json:"accepted"`
		WrongAnswer  int `json:"wrongAnswer"`
		TimeLimit    int `json:"timeLimit"`
		MemoryLimit  int `json:"memoryLimit"`
		OutputLimit  int `json:"outputLimit"`
		CompileError int `json:"compileError"`
		RuntimeError int `json:"runtimeError"`
	} `json:"status"`
	URL interface{} `json:"url"`
}
