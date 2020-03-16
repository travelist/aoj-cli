package common

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

var LanguageToDefaultConfigParam = map[string]ConfigFileParam{
	"c":     DefaultConfigFileParamC,
	"cpp":   DefaultConfigFileParamCpp,
	"cpp11": DefaultConfigFileParamCpp11,
	"cpp14": DefaultConfigFileParamCpp14,
	"java":  DefaultConfigFileParamJava,
	"py":    DefaultConfigFileParamPy,
	"py3":   DefaultConfigFileParamPy3,
}

var DefaultTemplateForC = `#include <stdio.h>

int main(){
	return 0;
}
`

var DefaultTemplateForCpp = `#include <stdio.h>
using namespace std;


int main(){
	return 0;
}
`

var DefaultTemplateForCpp11 = `#include <stdio.h>
using namespace std;

int main(){
	return 0;
}
`

var DefaultTemplateForCpp14 = `#include <stdio.h>
using namespace std;

int main(){
	return 0;
}
`

var DefaultTemplateForJava = `import java.util.*;
import java.lang.*;

class Main {
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
	}
}
`

var DefaultTemplateForPy = ``

var DefaultTemplateForPy3 = ``

var LanguageToDefaultTemplate = map[string]string{
	"c":     DefaultTemplateForC,
	"cpp":   DefaultTemplateForCpp,
	"cpp11": DefaultTemplateForCpp11,
	"cpp14": DefaultTemplateForCpp14,
	"java":  DefaultTemplateForJava,
	"py":    DefaultTemplateForPy,
	"py3":   DefaultTemplateForPy3,
}

var MetadataFileTemplate = `problem_id: {{.ProblemId}}`

type MetadataFileParam struct {
	ProblemId string
}
