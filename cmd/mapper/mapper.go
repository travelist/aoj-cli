package mapper

// Mapping language identifier
var aojConfigToSystemLanguageName = map[string]string{
	"c":       "C",
	"cpp11":   "C++11",
	"cpp14":   "C++14",
	"cpp":     "C++",
	"java":    "JAVA",
	"py":      "Python",
	"py3":     "Python3",
	"d":       "D",
	"csharp":  "C#",
	"go":      "Go",
	"js":      "JavaScript",
	"php":     "PHP",
	"scala":   "Scala",
	"haskell": "Haskell",
	"ocaml":   "OCaml",
	"kotlin":  "Kotlin",
	"ruby":    "Ruby",
	"rust":    "Rust",
}

func ConvertToAOJLanguage(lang string) string {
	return aojConfigToSystemLanguageName[lang]
}
