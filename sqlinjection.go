package main

import (
	"fmt"
	"strings"
)

type SQLInfo struct {
	Code              []string
	SQLLine           []int
	FoundSQLStatement bool
}

func getSQLStatements() []string {
	// Return slice of SQLStatements to search for
	return []string{"insert", "INSERT", "update", "UPDATE", "delete", "DELETE", "select", "SELECT"}
}

func (p *SQLInfo) searchSQLStatements() {
	// Get list of all SQLStatements to search for
	sqlstatements := getSQLStatements()

	// Search input source code for SQLStatements
	for codeLine := range p.Code {
		for statement := range sqlstatements {
			if strings.Contains(p.Code[codeLine], sqlstatements[statement]) {
				p.SQLLine = append(p.SQLLine, codeLine+1)
				if !p.FoundSQLStatement {
					p.FoundSQLStatement = true
				}
			}
		}
	}
}

func (p *SQLInfo) printResults() {
	printSQLBanner()
	// Loop through found results and output to screen
	for line := range p.SQLLine {
		results := fmt.Sprintf("%s%d%s%s", "Line ", p.SQLLine[line]-1, " - ", strings.TrimSpace(p.Code[p.SQLLine[line]-1]))
		fmt.Println(results)
	}
}

// Main entry point from calling function
func searchSQL(code []string) {
	// Initialise a new SQL struct and set the sourcecode
	var sqlinfo = new(SQLInfo)
	sqlinfo.Code = code

	// Search for SQL statements
	sqlinfo.searchSQLStatements()
	// Print resutls if any are found
	if sqlinfo.FoundSQLStatement {
		sqlinfo.printResults()
	}

}
