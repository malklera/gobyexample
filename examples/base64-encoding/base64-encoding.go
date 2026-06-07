// Go provides built-in support for [base64
// encoding/decoding](https://en.wikipedia.org/wiki/Base64).

package main

// This syntax imports the `encoding/base64` package with
// the `b64` name instead of the default `base64`. It'll
// save us some space below.
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// Here's the `string` we'll encode/decode.
	data := "abc123!?$*&()'-=@~"
	fmt.Println("data:", data)
	fmt.Println()

	// Go supports both standard and URL-compatible
	// base64. Here's how to encode using the standard
	// encoder. The encoder requires a `[]byte` so we
	// convert our `string` to that type.
	stdEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("stdEnc:", stdEnc)

	// Decoding may return an error, even if you think
	// the input is well formed, get accustom to always
	// check it.
	stdDec, err := b64.StdEncoding.DecodeString(stdEnc)
	if err != nil {
		fmt.Println("b64.StdEncoding.DecodeString(stdEnc):", err)
	}
	fmt.Println("stdDec:", string(stdDec))
	fmt.Println()

	// This encodes/decodes using a URL-compatible base64
	// format.
	urlEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("urlEnc:", urlEnc)
	urlDec, err := b64.URLEncoding.DecodeString(urlEnc)
	if err != nil {
		fmt.Println("b64.StdEncoding.DecodeString(stdEnc):", err)
	}
	fmt.Println("urlDec", string(urlDec))
}
