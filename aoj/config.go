package aoj

type Config struct {
	Username string
	Password string
	Language string

	Gen struct {
		WorkspacePath     string
		DirectoryName     string
		TemplateDirectory string
	}

	Test struct {
		BeforeAll  string
		BeforeEach string
		Run        string
		AfterEach  string
		AfterAll   string
	}

	Submit struct {
		DefaultFileName string
	}
}

func NewConfig() *Config {

}
