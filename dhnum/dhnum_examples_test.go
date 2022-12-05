package dhnum_test

import (
	"fmt"

	"github.com/vidhill/godoc-examples/dhnum"
)

func ExampleAdd() {
	res := dhnum.Add(1, 2)

	fmt.Println(res)
	// Output: 3
}

func ExampleAdd_example2() {
	res := dhnum.Add(4, 2, 1)

	fmt.Println(res)
	// Output: 7
}
