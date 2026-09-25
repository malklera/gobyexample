package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("inkscape", "source.svg", "--export-type=pdf", "--export-filename=document.pdf")
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("inkscape failed: %v\n%s\n", err, out)
	}
}
