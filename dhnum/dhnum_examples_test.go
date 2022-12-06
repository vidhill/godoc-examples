package dhnum_test

import (
	"fmt"

	"github.com/vidhill/godoc-examples/dhnum"
)

// Comments above tests are rendered alongside the example code
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

// Multiple lines of results
func ExampleAdd_exampleMulti() {
	for _, i := range []int{1, 2, 3} {
		res := dhnum.Add(4, i)
		fmt.Println(res)
	}

	// Output:
	// 5
	// 6
	// 7
}
