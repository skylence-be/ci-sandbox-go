// Command ci-sandbox-go prints a greeting. It exists so the CI pipeline has
// real code to lint, vet, and test.
package main

import (
	"fmt"

	"github.com/skylence-be/ci-sandbox-go/greet"
)

func main() {
	fmt.Println(greet.Greeting("skylence"))
}
