package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	// Sprintf substitutes the name parameter's value for the %v format verb.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
