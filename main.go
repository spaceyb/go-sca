package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	// Print Info Banner
	printInfoBanner()

	// Get command line arguments and load file
	args := os.Args[1:]
	if len(args) != 0 {
		filename := strings.Join(args, "")
		f, err := os.Open(filename)

		if err != nil {
			log.Fatal()
		}

		defer f.Close()
		scanner := bufio.NewScanner(f)

		// Empty slice to store the sourcecode
		var sourcecode = make([]string, 0)

		for scanner.Scan() {
			sourcecode = append(sourcecode, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		// Search for possible vulnerabilities
		searchSQL(sourcecode)
		searchCommandInjection(sourcecode)
		searchURLRedirection(sourcecode)
		searchSocketCommand(sourcecode)
		searchFileAccess(sourcecode)
		searchUserData(sourcecode)
		searchMiscOptions(sourcecode)
	} else {
		usage()
	}
}
