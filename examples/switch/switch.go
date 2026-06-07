// _Switch statements_ express conditionals across many
// branches.

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	// Here's a basic `switch`.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	fmt.Println()

	// You can use commas to separate multiple expressions
	// in the same `case` statement. We use the optional
	// `default` case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	fmt.Println()

	// `switch` without an expression is an alternate way
	// to express if/else logic. Here we also show how the
	// `case` expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	fmt.Println()

	// A type `switch` compares types instead of values.  You
	// can use this to discover the type of an interface
	// value.  In this example, the variable `t` will have the
	// type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	fmt.Print("whatAmI(true): ")
	whatAmI(true)
	fmt.Print("whatAmI(1): ")
	whatAmI(1)
	fmt.Print("whatAmI(hey): ")
	whatAmI("hey")
	fmt.Println()

	// This evaluate each case till it hit a true or default, stoping at the first
	// true, all the cases are excluyent to each other
	someString := "ab"
	switch {
	case someString == "a":
		fmt.Println("someString == a")
	case strings.EqualFold(someString, "a"):
		fmt.Println("strings.EqualFold(someString, a)")
	default:
		fmt.Println("default someString := ab")
	}
	fmt.Println()

	// If testing the same type for different values, you can declare it inline
	switch otherString := "cd"; otherString {
	case "c":
		fmt.Println("otherString == c")
	case "cd":
		fmt.Println("otherString == cd")
	default:
		fmt.Println("default otherString := cd")
	}
	fmt.Println()

	// or use one declared in another place.
	thirdString := "e"
	switch thirdString {
	case "e":
		fmt.Println("thirdString == e")
	case "ef":
		fmt.Println("thirdString == ef")
	default:
		fmt.Println("default thirdString := ef")
	}
}
