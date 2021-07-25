package main

import (
	"fmt"
	"os"
)

func banner() {
	fmt.Println("================================")
	fmt.Println("PHP Source Code Analyser\nWritten in Go")
	fmt.Println("================================")

}

func main() {
	// Print Banner
	banner()

	// Get command line arguments
	args := os.Args[1:]

	fmt.Println(args)
}
