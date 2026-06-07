// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path := filepath.Join(os.TempDir(), "data")
	dataFile, err := os.Create(path)
	check(err)
	_, err = dataFile.WriteString("hello\ngo\n")
	check(err)
	// Perhaps the most basic file reading task is
	// slurping a file's entire contents into memory.
	content, err := os.ReadFile(path)
	check(err)
	fmt.Printf("os.ReadFile(%s)\n", path)
	fmt.Print(string(content))
	fmt.Println()

	// You'll often want more control over how and what
	// parts of a file are read. For these tasks, start
	// by `Open`ing a file to obtain an `os.File` value.
	file, err := os.Open(path)
	check(err)

	// Read some bytes from the beginning of the file.
	// Allow up to 5 to be read but also note how many
	// actually were read.
	byte5 := make([]byte, 5)
	read5, err := file.Read(byte5)
	check(err)
	fmt.Printf("file.Read(make([]byte, 5))\n")
	fmt.Printf("read '%d' bytes: '%s'\n", read5, string(byte5[:read5]))
	fmt.Println()

	// You can also `Seek` to a known location in the file
	// and `Read` from there.
	seek6, err := file.Seek(6, io.SeekStart)
	check(err)
	byte2 := make([]byte, 2)
	read2, err := file.Read(byte2)
	check(err)
	fmt.Println("file.Seek(6, io.SeekStart)")
	fmt.Println("file.Read(make([]byte, 2))")
	fmt.Printf("read '%d' bytes from '%d': '%s'\n", read2, seek6, string(byte2[:read2]))
	fmt.Println()

	// Other methods of seeking are relative to the
	// current cursor position,
	_, err = file.Seek(2, io.SeekCurrent)
	check(err)

	// and relative to the end of the file.
	_, err = file.Seek(-4, io.SeekEnd)
	check(err)

	// TODO: should make two examples that actually show the difference between
	// `Read` and `io.ReadAtLeast`

	// The `io` package provides some functions that may
	// be helpful for file reading. For example, reads
	// like the ones above can be more robustly
	// implemented with `ReadAtLeast`.
	seekAtLeast6, err := file.Seek(6, io.SeekStart)
	check(err)
	byteAtLeast2 := make([]byte, 2)
	readAtLeast2, err := io.ReadAtLeast(file, byteAtLeast2, 2)
	check(err)
	fmt.Println("file.Seek(6, io.SeekStart)")
	fmt.Println("file.Read(make([]byte, 2))")
	fmt.Printf("read '%d' bytes from '%d': '%s'\n", readAtLeast2, seekAtLeast6, string(byteAtLeast2))
	fmt.Println()

	// There is no built-in rewind, but
	// `Seek(0, io.SeekStart)` accomplishes this.
	_, err = file.Seek(0, io.SeekStart)
	check(err)

	// The `bufio` package implements a buffered
	// reader that may be useful both for its efficiency
	// with many small reads and because of the additional
	// reading methods it provides.
	bufioReader := bufio.NewReader(file)
	peek5, err := bufioReader.Peek(5)
	check(err)

	fmt.Println("file.Seek(0, io.SeekStart)")
	fmt.Println("bufio.NewReader(file)")
	fmt.Println("bufioReader.Peek(5)")
	fmt.Printf("read '5' bytes: '%s'\n", string(peek5))
	fmt.Println()

	// Close the file when you're done (usually this would
	// be scheduled immediately after `Open`ing with
	// `defer`).
	file.Close()
}
