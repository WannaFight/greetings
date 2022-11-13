package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello from go module, %s!", name)
	return message
}
