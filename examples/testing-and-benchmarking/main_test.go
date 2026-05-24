// Unit testing is an important part of writing
// principled Go programs. The `testing` package
// provides the tools we need to write unit tests
// and the `go test` command runs tests.

// For the sake of demonstration, this code is in package
// `main`, but it could be any package. Testing code
// typically lives in the same package as the code it tests.
package main

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

// We'll be testing this simple implementation of an
// integer minimum. Typically, the code we're testing
// would be in a source file named something like
// `intutils.go`, and the test file for it would then
// be named `intutils_test.go`.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// A test is created by writing a function with a name
// beginning with `Test`.
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// `t.Error*` will report test failures but continue
		// executing the test. `t.Fatal*` will report test
		// failures and stop the test immediately.
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Writing tests can be repetitive, so it's idiomatic to
// use a *table-driven style*, where test inputs and
// expected outputs are listed in a table and a single loop
// walks over them and performs the test logic.
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		// `t.Run` enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmark tests typically go in `_test.go` files and are
// named beginning with `Benchmark`.
// Any code that's required for the benchmark to run but should
// not be measured goes before this loop.
func BenchmarkIntMin(b *testing.B) {
	for b.Loop() {
		// The benchmark runner will automatically execute this loop
		// body many times to determine a reasonable estimate of the
		// run-time of a single iteration.
		IntMin(1, 2)
	}
}

// Now lets make a more realistic function

// atoiWrapper is just a wrapper that calls `strconv.Atoi()` what
// we care about is the fact that it returns an int, and error,
// but the error returned is a custom struct `strconv.NumError`
func atoiWrapper(in string) (int, error) {
	return strconv.Atoi(in)
}

func TestAtoiWrapper(t *testing.T) {
	var numbers = []struct {
		num  string
		nInt int
		err  error
	}{
		{"2", 2, nil},
		{"-3", -3, nil},
		// thi do not work
		// {"b", 0, &strconv.NumError{Func: "Atoi", Num: "b", Err: strconv.ErrSyntax}},
		{"b", 0, strconv.ErrSyntax},
	}

	for _, tt := range numbers {
		t.Run(tt.num, func(t *testing.T) {
			n, er := atoiWrapper(tt.num)
			if n != tt.nInt || !errors.Is(er, tt.err) {
				t.Errorf("got '%d, %v'; want '%d, %v'", n, er, tt.nInt, tt.err)
			}
		})
	}
}

var errCustom = errors.New("not a number")

func toInt(in string) (int, error) {
	n, err := strconv.Atoi(in)
	if err != nil {
		return 0, errCustom
	}
	return n, nil
}

func TestToInt(t *testing.T) {
	var numbers = []struct {
		number string
		nInt   int
		err    error
	}{
		{"3", 3, nil},
		{"-5", -5, nil},
		{"b", 0, errCustom},
	}

	for _, tt := range numbers {
		t.Run(tt.number, func(t *testing.T) {
			n, er := toInt(tt.number)
			if n != tt.nInt || er != tt.err {
				t.Errorf("got '%d, %v'; want '%d, %v'", n, er, tt.nInt, tt.err)
			}
		})
	}
}
