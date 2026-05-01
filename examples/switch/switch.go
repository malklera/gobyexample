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

	// You can use commas to separate multiple expressions
	// in the same `case` statement. We use the optional
	// `default` case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

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
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	// This evaluate each case till it hit a true or default, stoping at the first
	// true, all the cases are excluyent to each other
	someString := "bc"
	switch {
	case someString == "b":
		fmt.Println("someString == b")
	case strings.EqualFold(someString, "b"):
		fmt.Println("strings.EqualFold(someString == b)")
	default:
		fmt.Println("default someString := bc")
	}

	// If testing the same type for different values, you can declare it inline
	switch otherString := "bc"; otherString {
	case "b":
		fmt.Println("otherString == b")
	case "bc":
		fmt.Println("otherString == bc")
	default:
		fmt.Println("default otherString := bc")
	}

	// or use one declared in another place.
	thirdString := "bc"
	switch thirdString {
	case "b":
		fmt.Println("thirdString == b")
	case "bc":
		fmt.Println("thirdString == bc")
	default:
		fmt.Println("default thirdString := bc")
	}
}
