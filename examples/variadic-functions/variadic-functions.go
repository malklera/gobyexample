// [_Variadic functions_](https://en.wikipedia.org/wiki/Variadic_function)
// can be called with any number of trailing arguments.
// For example, `fmt.Println` is a common variadic
// function.

package main

import "fmt"

// Here's a function that will take an arbitrary number
// of `int`s as arguments.
func sum(nums ...int) {
	fmt.Print("nums: ", nums, " ")
	total := 0
	// Within the function, the type of `nums` is
	// equivalent to `[]int`. We can call `len(nums)`,
	// iterate over it with `range`, etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println("total:", total)
}

// This do not work, ... can only go with the last argument of the function
// func rest(char ...string, nums ...int) {
// 	...
// }

func rest(char string, nums ...int) {
	fmt.Println("char:", char)
	fmt.Println("nums:", nums)
	total := 0
	for i := range nums {
		total -= i
	}
	fmt.Println("total:", total)
}

func main() {

	// Variadic functions can be called in the usual way
	// with individual arguments.
	fmt.Println("sum(1, 2)")
	sum(1, 2)
	fmt.Println("sum(1, 2, 3)")
	sum(1, 2, 3)
	fmt.Println()

	// If you already have multiple args in a slice,
	// apply them to a variadic function using
	// `func(slice...)` like this.
	nums := []int{1, 2, 3, 4}
	fmt.Println("nums := []int{1, 2, 3, 4}")
	fmt.Println("sum(nums...)")
	sum(nums...)
	fmt.Println()

	fmt.Println("rest(someString, nums...)")
	rest("someString", nums...)
}
