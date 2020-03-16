package tmpl

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
