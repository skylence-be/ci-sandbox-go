// Package greet builds greetings.
package greet

import "fmt"

// Greeting returns a stable greeting for name; empty names get a fallback.
func Greeting(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("hello, %s", name)
}
