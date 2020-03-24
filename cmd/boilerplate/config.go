package boilerplate

var ConfigFileTemplate = `[general]
language = "{{.Language}}"
username = "{{.Username}}"
password = "{{.Password}}"

[gen]
template_file = "$HOME/.aoj-cli/template.txt"
destination_file_name = "{{.GenSourceFileName}}"

[test]
before_all="{{.TestBeforeAll}}"
before_each=""
command="{{.TestCommand}}"
after_each=""
after_all=""

[submit]
source_file_name = "{{.SubmitSourceFileName}}"`

var ValidLanguage = []string{
	"c",
	"cpp",
	"cpp11",
	"cpp14",
	"java",
	"py",
	"py3",
	"d",
	"csharp",
	"go",
	"js",
	"php",
	"scala",
	"haskell",
	"ocaml",
	"kotlin",
	"ruby",
	"rust",
}

type ConfigFileParam struct {
	Language             string
	Username             string
	Password             string
	GenSourceFileName    string
	TestBeforeAll        string
	TestCommand          string
	SubmitSourceFileName string
}

var DefaultConfigFileParamC = ConfigFileParam{
	Language:             "c",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "main.c",
	TestBeforeAll:        "gcc main.c",
	TestCommand:          "./a.out",
	SubmitSourceFileName: "main.c",
}

var DefaultConfigFileParamCpp = ConfigFileParam{
	Language:             "cpp",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "main.cpp",
	TestBeforeAll:        "g++ main.cpp -o a.out",
	TestCommand:          "./a.out",
	SubmitSourceFileName: "main.cpp",
}
var DefaultConfigFileParamCpp11 = ConfigFileParam{
	Language:             "cpp11",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "main.cpp",
	TestBeforeAll:        "g++ main.cpp -o a.out",
	TestCommand:          "./a.out",
	SubmitSourceFileName: "main.cpp",
}
var DefaultConfigFileParamCpp14 = ConfigFileParam{
	Language:             "cpp14",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "main.cpp",
	TestBeforeAll:        "g++ main.cpp -o a.out",
	TestCommand:          "./a.out",
	SubmitSourceFileName: "main.cpp",
}
var DefaultConfigFileParamJava = ConfigFileParam{
	Language:             "java",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "Main.java",
	TestBeforeAll:        "javac Main.java",
	TestCommand:          "java Main",
	SubmitSourceFileName: "Main.java",
}
var DefaultConfigFileParamPy = ConfigFileParam{
	Language:             "py",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "main.py",
	TestBeforeAll:        "",
	TestCommand:          "python main.py",
	SubmitSourceFileName: "main.py",
}
var DefaultConfigFileParamPy3 = ConfigFileParam{
	Language:             "py3",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "main.py",
	TestBeforeAll:        "",
	TestCommand:          "python main.py",
	SubmitSourceFileName: "main.py",
}

var DefaultConfigFileParamD = ConfigFileParam{
	Language:             "d",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamCSharp = ConfigFileParam{
	Language:             "csharp",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamGo = ConfigFileParam{
	Language:             "go",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamJavaScript = ConfigFileParam{
	Language:             "js",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamPhp = ConfigFileParam{
	Language:             "php",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamScala = ConfigFileParam{
	Language:             "scala",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamHaskell = ConfigFileParam{
	Language:             "haskell",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamOCaml = ConfigFileParam{
	Language:             "ocaml",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamKotlin = ConfigFileParam{
	Language:             "kotlin",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamRuby = ConfigFileParam{
	Language:             "ruby",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamRust = ConfigFileParam{
	Language:             "rust",
	Username:             "",
	Password:             "",
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}

var LanguageToDefaultConfigParam = map[string]ConfigFileParam{
	"c":       DefaultConfigFileParamC,
	"cpp":     DefaultConfigFileParamCpp,
	"cpp11":   DefaultConfigFileParamCpp11,
	"cpp14":   DefaultConfigFileParamCpp14,
	"java":    DefaultConfigFileParamJava,
	"py":      DefaultConfigFileParamPy,
	"py3":     DefaultConfigFileParamPy3,
	"d":       DefaultConfigFileParamD,
	"csharp":  DefaultConfigFileParamCSharp,
	"go":      DefaultConfigFileParamGo,
	"js":      DefaultConfigFileParamJavaScript,
	"php":     DefaultConfigFileParamPhp,
	"scala":   DefaultConfigFileParamScala,
	"haskell": DefaultConfigFileParamHaskell,
	"ocaml":   DefaultConfigFileParamOCaml,
	"kotlin":  DefaultConfigFileParamKotlin,
	"ruby":    DefaultConfigFileParamRuby,
	"rust":    DefaultConfigFileParamRust,
}
