// _Maps_ are Go's built-in [associative data type](https://en.wikipedia.org/wiki/Associative_array)
// (sometimes called _hashes_ or _dicts_ in other languages).

package main

import (
	"fmt"
	"maps"
)

func main() {

	// To create an empty map, use the builtin `make`:
	// `make(map[key-type]val-type)`.
	makeMap := make(map[string]int)
	fmt.Println("makeMap := make(map[string]int)")
	fmt.Println("makeMap:", makeMap)

	// Set key/value pairs using typical `name[key] = val`
	// syntax.
	makeMap["k1"] = 7
	makeMap["k2"] = 13

	// Printing a map with e.g. `fmt.Println` will show all of
	// its key/value pairs.
	fmt.Println("makeMap[\"k1\"] = 7")
	fmt.Println("makeMap[\"k2\"] = 13")
	fmt.Println("makeMap:", makeMap)

	// Get a value for a key with `name[key]`.
	fmt.Println("makeMap[\"k1\"]:", makeMap["k1"])

	// If the key doesn't exist, the
	// [zero value](https://go.dev/ref/spec#The_zero_value) of the
	// value type is returned.
	fmt.Println("makeMap[\"k3\"]:", makeMap["k3"])

	// The builtin `len` returns the number of key/value
	// pairs when called on a map.
	fmt.Println("len(makeMap):", len(makeMap))
	fmt.Println()

	// The builtin `delete` removes key/value pairs from
	// a map.
	fmt.Println("makeMap:", makeMap)
	delete(makeMap, "k2")
	fmt.Println("delete(makeMap, 'k2')")
	fmt.Println("makeMap:", makeMap)
	fmt.Println()

	// To remove *all* key/value pairs from a map, use
	// the `clear` builtin.
	fmt.Println("makeMap:", makeMap)
	clear(makeMap)
	fmt.Println("clear(makeMap)")
	fmt.Println("makeMap:", makeMap)
	fmt.Println()

	// The optional second return value when getting a
	// value from a map indicates if the key was present
	// in the map. This can be used to disambiguate
	// between missing keys and keys with zero values
	// like `0` or `""`.
	secondKey, exist := makeMap["k2"]
	fmt.Println("secondKey, exist := makeMap[\"k2\"]")
	fmt.Println("secondKey:", secondKey)
	fmt.Println("exist:", exist)
	fmt.Println()

	// You can also declare and initialize a new map in
	// the same line with this syntax.
	mapInit := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("mapInit := map[string]int{\"foo\": 1, \"bar\": 2}")
	fmt.Println("mapInit:", mapInit)
	fmt.Println()

	// The `maps` package contains a number of useful
	// utility functions for maps.
	secondInit := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("secondInit := map[string]int{\"foo\": 1, \"bar\": 2}")
	fmt.Println("maps.Equal(mapInit, secondInit):", maps.Equal(mapInit, secondInit))
}
