package boilerplate

const ConfigFileTemplate = `[gen]
template_file = "$HOME/.aoj-cli/template.txt"
destination_file_name = "{{.GenSourceFileName}}"

[test]
before_all="{{.TestBeforeAll}}"
before_each=""
command="{{.TestCommand}}"
after_each=""
after_all=""

[submit]
language = "{{.Language}}"
source_file_name = "{{.SubmitSourceFileName}}"
`

const (
	LangC          = "C"
	LangCPP11      = "C++11"
	LangCPP14      = "C++14"
	LangCPP        = "C++"
	LangJava       = "JAVA"
	LangPython     = "Python"
	LangPython3    = "Python3"
	LangD          = "D"
	LangCSharp     = "C#"
	LangGo         = "Go"
	LangJavascript = "JavaScript"
	LangPHP        = "PHP"
	LangScala      = "Scala"
	LangHaskel     = "Haskell"
	LangOCaml      = "OCaml"
	LangKotlin     = "Kotlin"
	LangRuby       = "Ruby"
	LangRust       = "Rust"
)

var ValidLanguageList = []string{
	LangC,
	LangCPP11,
	LangCPP14,
	LangCPP,
	LangJava,
	LangPython,
	LangPython3,
	LangD,
	LangCSharp,
	LangGo,
	LangJavascript,
	LangPHP,
	LangScala,
	LangHaskel,
	LangOCaml,
	LangKotlin,
	LangRuby,
	LangRust,
}

var ValidLanguageSet = map[string]bool{
	LangC:          true,
	LangCPP11:      true,
	LangCPP14:      true,
	LangCPP:        true,
	LangJava:       true,
	LangPython:     true,
	LangPython3:    true,
	LangD:          true,
	LangCSharp:     true,
	LangGo:         true,
	LangJavascript: true,
	LangPHP:        true,
	LangScala:      true,
	LangHaskel:     true,
	LangOCaml:      true,
	LangKotlin:     true,
	LangRuby:       true,
	LangRust:       true,
}

type ConfigFileParam struct {
	Language             string
	GenSourceFileName    string
	TestBeforeAll        string
	TestCommand          string
	SubmitSourceFileName string
}

var DefaultConfigFileParamC = ConfigFileParam{
	Language:             LangC,
	GenSourceFileName:    "main.c",
	TestBeforeAll:        "gcc main.c",
	TestCommand:          "./a.out",
	SubmitSourceFileName: "main.c",
}

var DefaultConfigFileParamCpp = ConfigFileParam{
	Language:             LangCPP,
	GenSourceFileName:    "main.cpp",
	TestBeforeAll:        "g++ main.cpp -o a.out",
	TestCommand:          "./a.out",
	SubmitSourceFileName: "main.cpp",
}
var DefaultConfigFileParamCpp11 = ConfigFileParam{
	Language:             LangCPP11,
	GenSourceFileName:    "main.cpp",
	TestBeforeAll:        "g++ main.cpp -o a.out",
	TestCommand:          "./a.out",
	SubmitSourceFileName: "main.cpp",
}
var DefaultConfigFileParamCpp14 = ConfigFileParam{
	Language:             LangCPP14,
	GenSourceFileName:    "main.cpp",
	TestBeforeAll:        "g++ main.cpp -o a.out",
	TestCommand:          "./a.out",
	SubmitSourceFileName: "main.cpp",
}
var DefaultConfigFileParamJava = ConfigFileParam{
	Language:             LangJava,
	GenSourceFileName:    "Main.java",
	TestBeforeAll:        "javac Main.java",
	TestCommand:          "java Main",
	SubmitSourceFileName: "Main.java",
}
var DefaultConfigFileParamPy = ConfigFileParam{
	Language:             LangPython,
	GenSourceFileName:    "main.py",
	TestBeforeAll:        "",
	TestCommand:          "python main.py",
	SubmitSourceFileName: "main.py",
}
var DefaultConfigFileParamPy3 = ConfigFileParam{
	Language:             LangPython3,
	GenSourceFileName:    "main.py",
	TestBeforeAll:        "",
	TestCommand:          "python main.py",
	SubmitSourceFileName: "main.py",
}

var DefaultConfigFileParamD = ConfigFileParam{
	Language:             LangD,
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamCSharp = ConfigFileParam{
	Language:             LangCSharp,
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamGo = ConfigFileParam{
	Language:             LangGo,
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamJavaScript = ConfigFileParam{
	Language:             LangJavascript,
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamPhp = ConfigFileParam{
	Language:             LangPHP,
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamScala = ConfigFileParam{
	Language:             LangScala,
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamHaskell = ConfigFileParam{
	Language:             LangHaskel,
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamOCaml = ConfigFileParam{
	Language:             LangOCaml,
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamKotlin = ConfigFileParam{
	Language:             LangKotlin,
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamRuby = ConfigFileParam{
	Language:             LangRuby,
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}
var DefaultConfigFileParamRust = ConfigFileParam{
	Language:             LangRust,
	GenSourceFileName:    "",
	TestBeforeAll:        "",
	TestCommand:          "",
	SubmitSourceFileName: "",
}

var LanguageToDefaultConfigParam = map[string]ConfigFileParam{
	LangC:          DefaultConfigFileParamC,
	LangCPP11:      DefaultConfigFileParamCpp,
	LangCPP14:      DefaultConfigFileParamCpp11,
	LangCPP:        DefaultConfigFileParamCpp14,
	LangJava:       DefaultConfigFileParamJava,
	LangPython:     DefaultConfigFileParamPy,
	LangPython3:    DefaultConfigFileParamPy3,
	LangD:          DefaultConfigFileParamD,
	LangCSharp:     DefaultConfigFileParamCSharp,
	LangGo:         DefaultConfigFileParamGo,
	LangJavascript: DefaultConfigFileParamJavaScript,
	LangPHP:        DefaultConfigFileParamPhp,
	LangScala:      DefaultConfigFileParamScala,
	LangHaskel:     DefaultConfigFileParamHaskell,
	LangOCaml:      DefaultConfigFileParamOCaml,
	LangKotlin:     DefaultConfigFileParamKotlin,
	LangRuby:       DefaultConfigFileParamRuby,
	LangRust:       DefaultConfigFileParamRust,
}
