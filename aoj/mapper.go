package aoj

// Mapping language identifier
var aojLanguageId = map[string]string{
	"c":     "C",
	"cpp14": "C++14",
	"cpp":   "C++",
	"d":     "D",
	"go":    "Go",
	"java":  "java",
	"js":    "JavaScript",
	"py3":   "Python3",
	"ruby":  "Ruby",
	"rust":  "Rust",
}

func ConvertToAOJLanguage(lang string) string {
	return aojLanguageId[lang]
}
