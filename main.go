package main

import (
	"bufio"
	"fmt"
	"github.com/xyproto/stringpainter"
	"os"
)

const versionString = "Yaloco 1.3.1"

func usage() {
	fmt.Println("Please provide a filename as the first argument, or provide data on stdin.")
}

func main() {
	var scanner *bufio.Scanner
	if len(os.Args) > 1 {
		if os.Args[1] == "-V" || os.Args[1] == "--version" {
			fmt.Println(versionString)
			os.Exit(0)
		} else if os.Args[1] == "-h" || os.Args[1] == "--help" {
			usage()
			os.Exit(0)
		} else if os.Args[1] == "-" {
			scanner = bufio.NewScanner(os.Stdin)
		} else {
			f, err := os.Open(os.Args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			scanner = bufio.NewScanner(f)
		}
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	// Increase the maximum line length to 1MB
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	// Colorize the input data
	for scanner.Scan() {
		fmt.Println(stringpainter.Colorize(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)

	}
}
