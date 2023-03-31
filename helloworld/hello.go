package helloworld

import "fmt"

func Hello() string {
	return "Hello, world."
}

func Hi(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
