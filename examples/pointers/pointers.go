// Go supports <em><a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)">pointers</a></em>,
// allowing you to pass references to values and records
// within your program.

package main

import "fmt"

// We'll show how pointers work in contrast to values with
// 2 functions: `zeroval` and `zeroptr`. `zeroval` has an
// `int` parameter, so arguments will be passed to it by
// value. `zeroval` will get a copy of `ival` distinct
// from the one in the calling function.
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` in contrast has an `*int` parameter, meaning
// that it takes an `int` pointer. The `*iptr` code in the
// function body then _dereferences_ the pointer from its
// memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the
// value at the referenced address.
func zeroptr(iptr *int) {
	*iptr = 0
}

type myStruct struct {
	m bool
	n int
	o string
	p [3]int
	q []string
}

func zeroValAllTypes(aVal bool, bVal int, cVal string, dVal [3]int, eVal []string,
	fVal map[string]int, gVal myStruct) {
	aVal = false
	bVal = 5
	cVal = "five"
	dVal = [...]int{4, 5, 6}
	// WARN: this changed the original value
	eVal[0] = "four"
	fVal["two"] = 4
	gVal.q[0] = "seven"
}

func zeroPtrAllTypes(aVal *bool, bVal *int, cVal *string, dVal *[3]int, eVal *[]string,
	fVal *map[string]int, gVal *myStruct) {
	*aVal = false
	*bVal = 7
	*cVal = "six"
	*dVal = [...]int{7, 8, 9}
	(*eVal)[1] = "five"
	(*fVal)["three"] = 5
	(*gVal).q[2] = "four"
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval(i):", i)

	// The `&i` syntax gives the memory address of `i`,
	// i.e. a pointer to `i`.
	zeroptr(&i)
	fmt.Println("zeroptr(&i):", i)

	// Pointers can be printed too.
	fmt.Println("pointer &i:", &i)

	// A new pointer to a value can be created with the
	// builtin function `new`.
	p := new(42)
	fmt.Println("*p:", *p)
	zeroptr(p)
	fmt.Println("zeroptr(p):", *p)

	fmt.Println()
	fmt.Println()

	name := "leopoldo" // declare a variable
	fmt.Println("name := \"leopoldo\"")
	fmt.Println("showValue(&name)")
	showValue(&name) // pass a pointer to a variable
	// the following give error
	// showValue(name)
	// showValue(*name)
	fmt.Println()
	fmt.Println()

	name2 := new(string)
	*name2 = "juan"
	fmt.Println("*name2 = \"juan\"")
	fmt.Println("printValue(name2, \"pepe\")")
	printValue(name2, "pepe") // pass a pointer
	// name = "teo"		this give error because i try to asign string to adress
	*name2 = "jhon"
	printValue(name2, "jakob")
	// the following give error
	// printValue(*name, "mujica")
	// printValue(&name, "fedor")

	testSlice := []string{"a", "b", "c", "d"}
	fmt.Println("testSlice[0]:", testSlice[0])
	showSlice(&testSlice)
	fmt.Println("testSlice[0]:", testSlice[0])
	fmt.Println()
	fmt.Println()
	fmt.Println()



	var a bool = true
	var b int = 10
	var c string = "ten"
	var d = [3]int{1, 2, 3}
	var e = []string{"one", "two", "three"}
	var f = map[string]int{"one": 1, "two": 2, "three": 3}
	var g = myStruct{
		m: true,
		n: 10,
		o: "ten",
		p: [3]int{1, 2, 3},
		q: []string{"one", "two", "three"},
	}
	fmt.Println("initial a:", a)
	fmt.Println("initial b:", b)
	fmt.Println("initial c:", c)
	fmt.Println("initial d:", d)
	fmt.Println("initial e:", e)
	fmt.Println("initial f:", f)
	fmt.Println("initial g:", g)

	fmt.Println()
	fmt.Println("calling zeroVal")
	fmt.Println()
	zeroValAllTypes(a, b, c, d, e, f, g)

	fmt.Println("after zeroVal a:", a)
	fmt.Println("after zeroVal b:", b)
	fmt.Println("after zeroVal c:", c)
	fmt.Println("after zeroVal d:", d)
	fmt.Println("Notice how the next values change, even when i do not passed a pointer, it is because of the type.")
	fmt.Println("after zeroVal e:", e)
	fmt.Println("after zeroVal f:", f)
	fmt.Println("after zeroVal g:", g)

	fmt.Println()
	fmt.Println("calling zeroPtr")
	fmt.Println()
	zeroPtrAllTypes(&a, &b, &c, &d, &e, &f, &g)

	fmt.Println("after zeroPtr a:", a)
	fmt.Println("after zeroPtr b:", b)
	fmt.Println("after zeroPtr c:", c)
	fmt.Println("after zeroPtr d:", d)
	fmt.Println("after zeroPtr e:", e)
	fmt.Println("after zeroPtr f:", f)
	fmt.Println("after zeroPtr g:", g)
}

func showSlice(testSlice *[]string) {
	// func showSlice(testSlice []*string) {
	// func showSlice(testSlice *([]string)) {
	// func showSlice(testSlice *string) {
	fmt.Println("showSlice")
	// fmt.Println("testSlice[0]:", testSlice[0])
	fmt.Println("testSlice:", testSlice)
	fmt.Println("*testSlice:", *testSlice)
	fmt.Println("(*testSlice)[0]:", (*testSlice)[0])
	// fmt.Println("&testSlice[0]:", &testSlice[0])
	(*testSlice)[0] = "e"
}

func showValue(name *string) { // recibe a pointer
	fmt.Println()
	fmt.Println("showValue(name *string)")
	fmt.Println("name:", name)   // pointer address
	fmt.Println("*name:", *name) // value on that adress

	fmt.Println()
}

func printValue(name *string, newName string) {
	fmt.Println()
	fmt.Println("printValue(name *string, newName string)")
	fmt.Println("name:", name)
	fmt.Println("*name:", *name)

	fmt.Println("*name = newName")
	*name = newName
	// the folowing give error
	// name = newName
	// &name = newName

	fmt.Println("name:", name)
	fmt.Println("*name:", *name)
	fmt.Println()
}


