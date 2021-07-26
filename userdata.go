package main

import (
	"fmt"
	"strings"
)

type UserInfo struct {
	Code  []string
	Line  []int
	Found bool
}

func getUserCommands() []string {
	// Return a slice containing Injections to search for
	return []string{"$_GET", "$_POST", "$HTTP_GET_VARS", "$_COOKIE", "$HTTP_COOKIE_VARS",
		"$_REQUEST", "$_FILES", "$HTTP_POST_FILES", "$_SERVER['REQUEST_METHOD']",
		"$_SERVER['QUERY_STRING']", "$_SERVER[‘REQUEST_URI’]", "$_SERVER[‘HTTP_ACCEPT’]",
		"$_SERVER[‘HTTP_ACCEPT_CHARSET’]", "$_SERVER[‘HTTP_ACCEPT_ENCODING’]",
		"$_SERVER[‘HTTP_ACCEPT_LANGUAGE’]", "$_SERVER[‘HTTP_CONNECTION’]",
		"$_SERVER[‘HTTP_HOST’]", "$_SERVER[‘HTTP_REFERER’]", "$_SERVER[‘HTTP_USER_AGENT’]",
		"$_SERVER[‘PHP_SELF’]"}
}

func (p *UserInfo) searchUserCommands() {
	// Get list of all InjectionStatements to search for
	usercommands := getUserCommands()

	// Search input source code for Injection Commands
	for codeLine := range p.Code {
		for command := range usercommands {
			if strings.Contains(p.Code[codeLine], usercommands[command]) {
				p.Line = append(p.Line, codeLine+1)
				if !p.Found {
					p.Found = true
				}
			}
		}
	}
}

func (p *UserInfo) printResults() {
	printUserInfoBanner()
	// Loop through found results and output to screen
	for line := range p.Line {
		results := fmt.Sprintf("%s%d%s%s", "Line ", p.Line[line], " - ", strings.TrimSpace(p.Code[p.Line[line]-1]))
		fmt.Println(results)
	}
}

func searchUserData(code []string) {
	// Initialise a new Info struct and set the sourcecode
	var userinfo = new(UserInfo)
	userinfo.Code = code

	// Search for Command Injection
	userinfo.searchUserCommands()
	// Print results if any are found
	if userinfo.Found {
		userinfo.printResults()
	}
}
