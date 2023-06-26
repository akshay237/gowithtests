package hello

import "fmt"

const (
	prefixEnglish = "Hello, "
	prefixSpanish = "Hola, "
	prefixFrench  = "Benjour, "
)

func Hello(name, language string) string {
	if len(name) == 0 {
		name = "World"
	}
	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = prefixSpanish
	case "french":
		prefix = prefixFrench
	default:
		prefix = prefixEnglish
	}
	return
}

func Add(a, b int) (sum int) {
	sum = a + b
	return
}

func main() {
	fmt.Println(Hello("akki", ""))
	fmt.Println(Hello("", ""))
}
