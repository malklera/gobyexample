package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	switch e {
	case nil:
		break
	case io.EOF:
		break
	default:
		fmt.Println(e)
		os.Exit(1)
	}
}

var content = `2026/05/29 20:38:24 Success.
2026/05/30 17:04:26 Success.
2026/05/31 08:45:46 Success.
2026/05/31 17:59:32 Error.
2026/05/31 19:07:57 Success.
2026/06/01 14:00:13 Success.
2026/06/01 21:20:56 Success.
`

func main() {
	path := "example.log"
	file, err := os.Create(path)
	check(err)
	defer os.Remove(path)
	_, err = file.WriteString(content)
	check(err)
	file.Close()
	// The idea of this two examples is searching for the last "Something"
	fmt.Println("usingScanner")
	fileScanner, err := os.Open(path)
	check(err)
	lineScanner := usingScanner(fileScanner)
	fileScanner.Close()
	fmt.Println("lineScanner:", lineScanner)

	// Is this a good idea? test it.
	fmt.Println("usingSeek")
	fileSeek, err := os.Open(path)
	check(err)

	lineSeek := usingSeek(fileSeek)
	fileSeek.Close()
	fmt.Println("lineSeek:", lineSeek)

	// NOTE: is not a good idea in general, it only made sense using Seek when you hit
	// the 10000 lines of logs. 
	// Below 5000 lines logs, Scanner is faster
	// Around 5000 lines logs, They are equal
	// Above 5000 lines logs, Seek is faster
	// Around 10.000 lines logs, Seek is twice as fast as Scanner
}

// usingScanner read the whole file into memory, then loop over it from the end
// backwards to search for the last error that was logged.
func usingScanner(file *os.File) string{
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	for n := 1; n < len(lines); n++ {
		if strings.Contains(lines[len(lines)-n], "Error") {
			return lines[len(lines)-n]
		}
	}
	return ""
}

// usingSeek read the file from the end backwards one byte at a time, in each new
// line it will read a complete line and check for the string "Error", what we care
// about, changing "Error" for some other know string we search for the last
// ocurrence of it.
func usingSeek(file *os.File) string {
	info, err := file.Stat()
	check(err)
	size := info.Size()
	// two bytes for the last `\n`
	for n := int64(2); n < size; n++ {
		// Seek from the end backwards
		_, err = file.Seek(-n, io.SeekEnd)
		check(err)
		// Read one byte
		newLine := make([]byte, 1)
		_, err := file.Read(newLine)
		check(err)
		// Till you find a new line `\n`
		if bytes.Equal(newLine, []byte("\n")) {
			reader := bufio.NewReader(file)
			// Read a string up to and including `\n`
			// This is for when you logs are newline separated.
			line, err := reader.ReadString('\n')
			check(err)
			if strings.Contains(line, "Error") {
				return line
			}
		}

		// Beginning of the file.
		if n == size-1 {
			reader := bufio.NewReader(file)
			// Read a string up to and including `\n`
			// This is for when you logs are newline separated.
			line, err := reader.ReadString('\n')
			check(err)
			if strings.Contains(line, "Error") {
				return line
			}
		}
	}
	return ""
}
