package main

import (
	"fmt"
	"strings"
)

type InjInfo struct {
	Code           []string
	InjectionLine  []int
	FoundInjection bool
}

func getInjectionCommands() []string {
	// Return a slice containing Injections to search for
	return []string{"exec", "popen", "proc_close", "proc_open", "proc_get_status", "proc_nice", "proc_terminate",
		"shell_exec", "system", "eval", "assert", "call_user_func", "call_user_method", "create_function"}
}

func (p *InjInfo) searchInjection() {
	// Get list of all InjectionStatements to search for
	injectioncommands := getInjectionCommands()

	// Search input source code for Injection Commands
	for codeLine := range p.Code {
		for command := range injectioncommands {
			if strings.Contains(p.Code[codeLine], injectioncommands[command]) {
				p.InjectionLine = append(p.InjectionLine, codeLine+1)
				if !p.FoundInjection {
					p.FoundInjection = true
				}
			}
		}
	}
}

func (p *InjInfo) printResults() {
	printCommandInjectionBanner()
	// Loop through found results and output to screen
	for line := range p.InjectionLine {
		results := fmt.Sprintf("%s%d%s%s", "Line ", p.InjectionLine[line], " - ", strings.TrimSpace(p.Code[p.InjectionLine[line]-1]))
		fmt.Println(results)
	}
}

func searchCommandInjection(code []string) {
	// Initialise a new Info struct and set the sourcecode
	var injinfo = new(InjInfo)
	injinfo.Code = code

	// Search for Command Injection
	injinfo.searchInjection()
	// Print results if any are found
	if injinfo.FoundInjection {
		injinfo.printResults()
	}

}
