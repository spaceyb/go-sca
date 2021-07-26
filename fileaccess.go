package main

import (
	"fmt"
	"strings"
)

type FileInfo struct {
	Code  []string
	Line  []int
	Found bool
}

func getFileCommands() []string {
	// Return a slice containing Injections to search for
	return []string{"fopen", "readfile", "file", "fpassthru", "gzopen",
		"gzfile", "readgzfile", "copy", "rename", "rmdir",
		"mkdir", "unlink", "file_get_contents", "file_put_contents",
		"parse_ini_file"}
}

func (p *FileInfo) searchFileCommands() {
	// Get list of all InjectionStatements to search for
	filecommands := getFileCommands()

	// Search input source code for Injection Commands
	for codeLine := range p.Code {
		for command := range filecommands {
			if strings.Contains(p.Code[codeLine], filecommands[command]) {
				p.Line = append(p.Line, codeLine+1)
				if !p.Found {
					p.Found = true
				}
			}
		}
	}
}

func (p *FileInfo) printResults() {
	printFileInfoBanner()
	// Loop through found results and output to screen
	for line := range p.Line {
		results := fmt.Sprintf("%s%d%s%s", "Line ", p.Line[line], " - ", strings.TrimSpace(p.Code[p.Line[line]-1]))
		fmt.Println(results)
	}
}

func searchFileAccess(code []string) {
	// Initialise a new Info struct and set the sourcecode
	var fileinfo = new(FileInfo)
	fileinfo.Code = code

	// Search for Command Injection
	fileinfo.searchFileCommands()
	// Print results if any are found
	if fileinfo.Found {
		fileinfo.printResults()
	}
}
