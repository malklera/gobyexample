// In Go, an _array_ is a numbered sequence of elements of a
// specific length. In typical Go code, [slices](slices) are
// much more common; arrays are useful in some special
// scenarios.

package main

import "fmt"

func main() {

	// Here we create an array `a5` that will hold exactly
	// 5 `int`s. The type of elements and length are both
	// part of the array's type. By default an array is
	// zero-valued, which for `int`s means `0`s.
	var a5 [5]int
	fmt.Println("a5:", a5)

	// We can set a value at an index using the
	// `array[index] = value` syntax, and get a value with
	// `array[index]`.
	a5[4] = 100
	fmt.Println("a5[4] = 100")
	fmt.Println("a5:", a5)
	fmt.Println("a5[4]:", a5[4])

	// The builtin `len` returns the length of an array.
	fmt.Println("len(a5):", len(a5))
	fmt.Println()

	// Use this syntax to declare and initialize an array
	// in one line.
	b5 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b5 := [5]int{1, 2, 3, 4, 5}")
	fmt.Println("b5:", b5)
	fmt.Println()

	// You can also have the compiler count the number of
	// elements for you with `...`
	compCount := [...]int{1, 2, 3, 4, 5}
	fmt.Println("compCount := [...]int{1, 2, 3, 4, 5}")
	fmt.Println("compCount:", compCount)
	fmt.Println()

	// If you specify the index with `:`, the elements in
	// between will be zeroed.
	compCountIndex := [...]int{100, 3: 400, 500}
	fmt.Println("compCountIndex := [...]int{100, 3: 400, 500}")
	fmt.Println("compCountIndex:", compCountIndex)
	fmt.Println()

	// Array types are one-dimensional, but you can
	// compose types to build multi-dimensional data
	// structures.
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD: ", twoD)
	fmt.Println()

	// You can create and initialize multi-dimensional
	// arrays at once too.
	twoDInit := [2][3]int{
		{0, 1, 2},
		{1, 2, 3},
	}
	fmt.Println("twoDInit: ", twoDInit)
}
