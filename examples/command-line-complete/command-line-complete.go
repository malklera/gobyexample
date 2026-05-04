// Some command-line tools, like the `go` tool or `git`
// have many *subcommands*, each with its own set of
// flags. For example, `go build` and `go get` are two
// different subcommands of the `go` tool.
// The `flag` package lets us easily define simple
// subcommands that have their own flags.
// But things get complicated when you mix subcommands
// and flags like `git help` and `git -v` better to
// have `cmd subcmd flags` or `cmd flags`or use a
// third party library.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Declare a flag for `./command-line-complete`
	mixFlag := ""
	flag.StringVar(&mixFlag, "mix", "", "the mix flag")

	// We declare a subcommand using the `NewFlagSet`
	// function, and proceed to define new flags specific
	// for this subcommand.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// For a different subcommand we can define different
	// supported flags.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Has to check if any argument was passed, otherwise
	// the `os.Args[1]` below panics.
	if len(os.Args) < 2 {
		// Better to define a custom `Usage()` function
		// to show the usage of the command and all
		// subcommands depending on what it is called
		// e.g. if calling `subCmd.Usage()` do not show
		// the `Usage()` of the parent cmd
		fmt.Println("Available flags.")
		flag.Usage()
		fmt.Println("Available subcommands.")
		fooCmd.Usage()
		barCmd.Usage()
		os.Exit(1)
	}

	// Check which subcommand is invoked.
	switch os.Args[1] {
	// For every subcommand, we parse its own flags and
	// have access to trailing positional arguments.
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		flag.Parse()
		if flag.NFlag() == 0 {
			fmt.Println("Available flags.")
			flag.Usage()
			fmt.Println("Available subcommands.")
			fooCmd.Usage()
			barCmd.Usage()
			os.Exit(1)
		}
		fmt.Println("mix:", mixFlag)
		fmt.Println("tail:", flag.Args())
	}
}
