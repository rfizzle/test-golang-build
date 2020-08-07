package test_golang_build

import "fmt"

func main() {
	fmt.Println(defaultText())
}

func defaultText() string {
	return "Hello World"
}

func customText(target string) string {
	return fmt.Sprintf("Hello, %s", target)
}