package main

import (
	"fmt"
	"strings"
)

type URLInfo struct {
	Code    []string
	URLLine []int
	Found   bool
}

func getURLRedirectionCommands() []string {
	// Return a slice containing URL Redirections to search for
	return []string{"http_redirect", "header", "HttpMessage::setResponeCode", "HttpMessage::setHeaders"}
}

func (p *URLInfo) searchURLRedirections() {
	// Get list of all SQLStatements to search for
	urlredirections := getURLRedirectionCommands()

	// Search input source code for SQLStatements
	for codeLine := range p.Code {
		for statement := range urlredirections {
			if strings.Contains(p.Code[codeLine], urlredirections[statement]) {
				p.URLLine = append(p.URLLine, codeLine+1)
				if !p.Found {
					p.Found = true
				}
			}
		}
	}
}

func (p *URLInfo) printResults() {
	printURLRedirection()
	// Loop through found results and output to screen
	for line := range p.URLLine {
		results := fmt.Sprintf("%s%d%s%s", "Line ", p.URLLine[line]-1, " - ", strings.TrimSpace(p.Code[p.URLLine[line]-1]))
		fmt.Println(results)
	}
}

// Main entry point from calling function
func searchURLRedirection(code []string) {
	// Initialise a new URLInfo struct
	urlInfo := new(URLInfo)
	urlInfo.Code = code

	// Search for URL redirections
	urlInfo.searchURLRedirections()
	// Print results if found
	if urlInfo.Found {
		urlInfo.printResults()
	}

}
