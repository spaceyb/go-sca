package main

import (
	"fmt"
	"strings"
)

type MiscInfo struct {
	Code           []string
	CommandLine    []int
	FoundInjection bool
}

func getMiscCommands() []string {
	// Return a slice containing Injections to search for
	return []string{"allow_url_fopen", "allow_url_include", "display_errors", "file_uploads", "upload_tmp_dir"}
}

func (p *MiscInfo) searchMiscInfo() {
	// Get list of all InjectionStatements to search for
	misccommands := getMiscCommands()

	// Search input source code for Injection Commands
	for codeLine := range p.Code {
		for command := range misccommands {
			if strings.Contains(p.Code[codeLine], misccommands[command]) {
				p.CommandLine = append(p.CommandLine, codeLine+1)
				if !p.FoundInjection {
					p.FoundInjection = true
				}
			}
		}
	}
}

func (p *MiscInfo) printResults() {
	printMiscInfoBanner()
	// Loop through found results and output to screen
	for line := range p.CommandLine {
		results := fmt.Sprintf("%s%d%s%s", "Line ", p.CommandLine[line], " - ", strings.TrimSpace(p.Code[p.CommandLine[line]-1]))
		fmt.Println(results)
	}
}

func searchMiscOptions(code []string) {
	// Initialise a new Info struct and set the sourcecode
	var miscinfo = new(MiscInfo)
	miscinfo.Code = code

	// Search for Command Injection
	miscinfo.searchMiscInfo()
	// Print results if any are found
	if miscinfo.FoundInjection {
		miscinfo.printResults()
	}

}
