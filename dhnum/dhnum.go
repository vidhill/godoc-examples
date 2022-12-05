// This is a simple example package
package dhnum

// Accepts n number of ints and returns the sum of all
func Add(is ...int) int {
	res := 0
	for _, v := range is {
		res += v
	}
	return res
}
