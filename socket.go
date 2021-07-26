package main

import (
	"fmt"
	"strings"
)

type SocketInfo struct {
	Code       []string
	SocketLine []int
	Found      bool
}

func getSocketCommands() []string {
	// Return a slice containing Injections to search for
	return []string{"socket_create", "socket_connect", "socket_write", "socket_send", "socket_recv", "fsocketopen", "pfsocketopen"}
}

func (p *SocketInfo) searchSocket() {
	// Get list of all InjectionStatements to search for
	socketcommands := getSocketCommands()

	// Search input source code for Injection Commands
	for codeLine := range p.Code {
		for command := range socketcommands {
			if strings.Contains(p.Code[codeLine], socketcommands[command]) {
				p.SocketLine = append(p.SocketLine, codeLine+1)
				if !p.Found {
					p.Found = true
				}
			}
		}
	}
}

func (p *SocketInfo) printResults() {
	printSocketBanner()
	// Loop through found results and output to screen
	for line := range p.SocketLine {
		results := fmt.Sprintf("%s%d%s%s", "Line ", p.SocketLine[line], " - ", strings.TrimSpace(p.Code[p.SocketLine[line]-1]))
		fmt.Println(results)
	}
}

func searchSocketCommand(code []string) {
	// Initialise a new Info struct and set the sourcecode
	var socketinfo = new(SocketInfo)
	socketinfo.Code = code

	// Search for Command Injection
	socketinfo.searchSocket()
	// Print results if any are found
	if socketinfo.Found {
		socketinfo.printResults()
	}

}
