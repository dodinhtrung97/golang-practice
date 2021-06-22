package assignment

import "fmt"

func newInt() int {
	// Direct
	var i int = 3
	fmt.Println(i)

	// Indirect
	a := 3
	fmt.Println(a)

	return a
}
