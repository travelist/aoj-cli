package mapper

// Mapping language identifier
var aojLanguageId = map[string]string{
	"c":     "C",
	"cpp11": "C++11",
	"cpp14": "C++14",
	"cpp":   "C++",
	"java":  "java",
	"py":    "Python",
	"py3":   "Python3",
	// TODO
	//"d":       "D",
	//"go":      "Go",
	//"js":      "JavaScript",
	//"php":     "Php",
	//"scala":   "Scala",
	//"haskell": "Haskell",
	//"ocaml":   "Ocaml",
	//"kotlin":  "Kotlin",
	//"ruby":    "Ruby",
	//"rust":    "Rust",
}

func ConvertToAOJLanguage(lang string) string {
	return aojLanguageId[lang]
}
