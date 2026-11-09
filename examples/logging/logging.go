// The Go standard library provides straightforward
// tools for outputting logs from Go programs, with
// the [log](https://pkg.go.dev/log) package for
// free-form output and the
// [log/slog](https://pkg.go.dev/log/slog) package for
// structured output.
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"log/slog"
)

// setLogs takes a path and set the different output for logs, `envDir/file`,
// `os.Stderr` and some notification command.
func setLogs(envDir string) (*os.File, io.Writer) {
	logPath := filepath.Join(envDir, "test-logs.log")
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("os.OpenFile(%s, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) : %v", logPath, err)
		return nil, os.Stderr
	}

	return logFile, io.MultiWriter(os.Stderr, logFile)
}

func main() {
	// Simply invoking functions like `Println` from the
	// `log` package uses the _standard_ logger, which
	// is already pre-configured for reasonable logging
	// output to `os.Stderr`. Additional methods like
	// `Fatal*` or `Panic*` will exit the program after
	// logging.
	log.Println("standard logger")

	// Loggers can be configured with _flags_ to set
	// their output format. By default, the standard
	// logger has the `log.Ldate` and `log.Ltime` flags
	// set, and these are collected in `log.LstdFlags`.
	// We can change its flags to emit time with
	// microsecond accuracy, for example.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("'log.SetFlags(log.LstdFlags | log.Lmicroseconds)'")

	// It also supports emitting the file name and
	// line from which the `log` function is called.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("'log.SetFlags(log.LstdFlags | log.Lshortfile)'")

	// It may be useful to create a custom logger and
	// pass it around. When creating a new logger, we
	// can set a _prefix_ to distinguish its output
	// from other loggers.
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from 'mylog := log.New(os.Stdout, \"my:\", log.LstdFlags)'")

	// We can set the prefix
	// on existing loggers (including the standard one)
	// with the `SetPrefix` method.
	mylog.SetPrefix("ohmy:")
	mylog.Println("from 'mylog.SetPrefix(\"ohmy:\")'")

	// Loggers can have custom output targets;
	// any `io.Writer` works.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// This call writes the log output into `buf`.
	buflog.Println("hello")

	// This will actually show it on standard output.
	fmt.Print("from buflog:", buf.String())
	fmt.Println()

	// The `slog` package provides
	// _structured_ log output. For example, logging
	// in JSON format is straightforward.
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// In addition to the message, `slog` output can
	// contain an arbitrary number of key=value
	// pairs.
	myslog.Info("hello again", "key", "val", "age", 25)

	fmt.Println()

	// Some times we want to send logs to multiple
	// outputs, in this example to os.Stderr and a file
	fmt.Println("Output logs to os.Stderr and test-logs.log")
	logFile, multiWriter := setLogs(".")
	defer logFile.Close()

	log.SetOutput(multiWriter)

	log.Printf("test with printf1")
}
