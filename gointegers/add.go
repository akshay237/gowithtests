package gointegers

import "fmt"

func Add(a, b int) (sum int) {
	sum = a + b
	return
}

func main() {
	sum := Add(1, 2)
	fmt.Println("Sum: ", sum)
}
