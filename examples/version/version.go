package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"
)

/*
TAG=$(git describe --tags --exact-match 2>/dev/null || echo dev)
go build -ldflags "-X 'main.tag=$TAG'"
*/
var tag = ""

func main() {
	bi, ok := debug.ReadBuildInfo()
	if ok {
		// name of program
		fmt.Print("nameProgram ")
		// version of the program, either the tag or the commit
		fmt.Printf("version %s ", tag)
		vcsTime := ""
		for _, s := range bi.Settings {
			if s.Key == "vcs.time" {
				vcsTime = s.Value
			}
		}
		date, err := time.Parse(time.RFC3339, vcsTime)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing time: %v\n", err)
		}
		// modification time asosiated with the commit
		fmt.Printf("(%v) ", date.Format(time.DateOnly))
		// year beginning project and author
		fmt.Println("copyright [year] [author]")
		fmt.Println()
		fmt.Println("Source", bi.Path)
		fmt.Println()
		fmt.Println("Build info:")
		fmt.Println(bi.GoVersion)
		for _, s := range bi.Settings {
			fmt.Printf("%s=%s\n", s.Key, s.Value)
		}
	} else {
		fmt.Fprintf(os.Stderr, "failed to debug.ReadBuildInfo")
	}

}
