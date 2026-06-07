// _Slices_ are an important data type in Go, giving
// a more powerful interface to sequences than arrays.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Unlike arrays, slices are typed only by the
	// elements they contain (not the number of elements).
	// An uninitialized slice equals to nil and has
	// length 0.
	var unInit []string
	fmt.Println("var unInit []string:", unInit)
	fmt.Println("unInit == nil:", unInit == nil)
	fmt.Println("len(unInit):", len(unInit))
	fmt.Println()

	// To create a slice with non-zero length, use
	// the builtin `make`. Here we make a slice of
	// `string`s of length `3` (initially zero-valued).
	// By default a new slice's capacity is equal to its
	// length; if we know the slice is going to grow ahead
	// of time, it's possible to pass a capacity explicitly
	// as an additional parameter to `make`.
	slice3 := make([]string, 3)
	fmt.Println("slice3 := make([]string, 3):", slice3)
	fmt.Println("len(slice3):", len(slice3))
	fmt.Println("cap(slice3):", cap(slice3))
	fmt.Println()

	// We can set and get just like with arrays.
	slice3[0] = "a"
	slice3[1] = "b"
	slice3[2] = "c"
	fmt.Println("slice3[0] = \"a\"")
	fmt.Println("slice3[1] = \"b\"")
	fmt.Println("slice3[2] = \"c\"")

	fmt.Println("slice3:", slice3)
	fmt.Println("slice3[2]:", slice3[2])

	// `len` returns the length of the slice as expected.
	fmt.Println("len(slice3):", len(slice3))
	fmt.Println("cap(slice3):", cap(slice3))
	fmt.Println()

	// In addition to these basic operations, slices
	// support several more that make them richer than
	// arrays. One is the builtin `append`, which
	// returns a slice containing one or more new values.
	// Note that we need to accept a return value from
	// `append` as we may get a new slice value.
	fmt.Println("slice3:", slice3)
	slice3 = append(slice3, "d")
	fmt.Println("slice3 = append(slice3, \"d\")")
	fmt.Println("slice3:", slice3)
	slice3 = append(slice3, "e", "f")
	fmt.Println("slice3 = append(slice3, \"e\", \"f\")")
	fmt.Println("slice3:", slice3)
	fmt.Println()

	// Slices can also be `copy`'d. Here we create an
	// empty slice `copySlice` of the same length as `s` and copy
	// into `copySlice` from `s`.
	copySlice := make([]string, len(slice3))
	fmt.Println("copySlice:", copySlice)
	fmt.Println("copy(copySlice, slice3)")
	copy(copySlice, slice3)
	fmt.Println("copySlice:", copySlice)
	fmt.Println()

	// Slices support a "slice" operator with the syntax
	// `slice[low:high]`. For example, this gets a slice
	// of the elements `s[2]`, `s[3]`, and `s[4]`.
	sliceLowHigh := slice3[2:5]
	fmt.Println("sliceLowHigh := slice3[2:5]")
	fmt.Println("sliceLowHigh:", sliceLowHigh)
	fmt.Println()

	// This slices up to (but excluding) `s[5]`.
	sliceHigh := slice3[:5]
	fmt.Println("sliceHigh := slice3[:5]")
	fmt.Println("sliceHigh:", sliceHigh)
	fmt.Println()

	// And this slices up from (and including) `s[2]`.
	sliceLow := slice3[2:]
	fmt.Println("sliceLow := slice3[2:]")
	fmt.Println("sliceLow:", sliceLow)
	fmt.Println()

	// We can declare and initialize a variable for slice
	// in a single line as well.
	sliceInit := []string{"g", "h", "i"}
	fmt.Println("sliceInit := []string{\"g\", \"h\", \"i\"}")
	fmt.Println("sliceInit:", sliceInit)
	fmt.Println()

	// The `slices` package contains a number of useful
	// utility functions for slices.
	secondInit := []string{"g", "h", "i"}
	fmt.Println("secondInit := []string{\"g\", \"h\", \"i\"}")
	fmt.Println("slices.Equal(sliceInit, secondInit):", slices.Equal(sliceInit, secondInit))
	fmt.Println()

	// Slices can be composed into multi-dimensional data
	// structures. The length of the inner slices can
	// vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD: ", twoD)
}
