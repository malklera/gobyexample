// Go offers built-in support for JSON encoding and
// decoding, including to and from built-in and custom
// data types.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Only exported fields will be encoded/decoded in JSON.
// Fields must start with capital letters to be exported.
type rspNoTag struct {
	Page   int
	Fruits []string
}

// The `tag`s allow us to define the name of the keys in
// in the json file, if it do not have them, it will be
// the same as the field name.
type rspTag struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// First we'll look at encoding basic data types to
	// JSON strings. Here are some examples for atomic
	// values.
	bolB, _ := json.Marshal(true)
	fmt.Println("string(bolB):", string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println("string(intB):", string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println("string(fltB):", string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println("string(strB):", string(strB))

	// And here are some for slices and maps, which encode
	// to JSON arrays and objects as you'd expect.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println("string(slcB):", string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println("string(mapB):", string(mapB))
	fmt.Println()

	// The JSON package can automatically encode your
	// custom data types. It will only include exported
	// fields in the encoded output and will by default
	// use those names as the JSON keys.
	res1D := &rspNoTag{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println("string(res1B):", string(res1B))

	// You can use tags on struct field declarations
	// to customize the encoded JSON key names. Check the
	// definition of `response2` above to see an example
	// of such tags.
	res2D := &rspTag{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println("string(res2B):", string(res2B))
	fmt.Println()

	// Now let's look at decoding JSON data into Go
	// values. Here's an example for a generic data
	// structure.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// We need to provide a variable where the JSON
	// package can put the decoded data. This
	// `map[string]interface{}` will hold a map of strings
	// to arbitrary data types.
	var dat map[string]interface{}

	// Here's the actual decoding, and a check for
	// associated errors.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println("dat:", dat)

	// In order to use the values in the decoded map,
	// we'll need to convert them to their appropriate type.
	// For example here we convert the value in `num` to
	// the expected `float64` type.
	num := dat["num"].(float64)
	fmt.Println("num := dat[\"num\"].(float64)")
	fmt.Println("num:", num)

	// Accessing nested data requires a series of
	// conversions.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println("str1:", str1)

	// We can also decode JSON into custom data types.
	// This has the advantages of adding additional
	// type-safety to our programs and eliminating the
	// need for type assertions when accessing the decoded
	// data.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := rspTag{}
	if err := json.Unmarshal([]byte(str), &res); err != nil {
		panic(err)
	}
	fmt.Println("res:", res)
	fmt.Println("res.Fruits[0]:", res.Fruits[0])

	// In the examples above we always used bytes and
	// strings as intermediates between the data and
	// JSON representation on standard out. We can also
	// stream JSON encodings directly to `os.Writer`s like
	// `os.Stdout` or even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	if err := enc.Encode(d); err != nil {
		panic(err)
	}

	// Streaming reads from `os.Reader`s like `os.Stdin`
	// or HTTP request bodies is done with `json.Decoder`.
	dec := json.NewDecoder(strings.NewReader(str))
	res1 := rspTag{}
	if err := dec.Decode(&res1); err != nil {
		panic(err)
	}
	fmt.Println(res1)
	fmt.Println()

	// Reading and writing json files.
	// Create a file as part of the setup.
	content := []byte(`{"page": 1, "fruits": ["apple", "peach"]}`)
	path := filepath.Join(os.TempDir(), "test.json")
	if err := os.WriteFile(path, content, 0644); err != nil {
		panic(err)
	}

	// Here is where we actually read from a json file.
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fromPath := rspTag{}
	// Has to pass a pointer to `json.Unmarshal`.
	if err := json.Unmarshal(file, &fromPath); err != nil {
		panic(err)
	}
	fmt.Println("fromPath:", fromPath)
	fmt.Println("fromPath.Fruits[0]:", fromPath.Fruits[0])

	// Now if we want to write to a json file.
	// Option 1: Create a new file, was implemented in the setup above.

	// Option 2: Add data to an existing file.
	// We still have the data from the file in `fromPath` so we take that
	// and create another entry, then marshal both.

	secondContent := rspTag{
		Page:   2,
		Fruits: []string{"banana", "orange"}}

	finalContent := append([]rspTag{fromPath}, secondContent)
	marshalFinalContent, err := json.Marshal(&finalContent)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(path, marshalFinalContent, 0644); err != nil {
		panic(err)
	}

	// Then check the file.
	finalFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("finalFile:", string(finalFile))
}
