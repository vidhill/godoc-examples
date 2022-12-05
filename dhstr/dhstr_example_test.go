// Whole file example
package dhstr_test

import (
	"fmt"

	"github.com/vidhill/godoc-examples/dhstr"
)

type Person struct {
	name string
}

func Example() {
	p := Person{
		name: "David",
	}
	res := dhstr.Prefix(p.name)

	fmt.Println(res)
	// Output: hello_David
}
