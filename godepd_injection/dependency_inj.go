package godepd_injection

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Printf send the whole content to os.stdout writer and it implements the writer interface func and print whole thing
// In our case we will send the content to the writer buffer

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, name)
}

// to make our test pass send the whole greeting msg
func NewGreet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s, welcome", name)
}

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	NewGreet(w, "world")
}

func main() {
	// to take os.stdout argument we will io.writer as input to printf that will be implemented by both bytes.Buffer and os.stdout
	//Greet(os.Stdout, "Akshay")
	NewGreet(os.Stdout, "akki")
}
