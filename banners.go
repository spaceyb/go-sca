package main

import (
	"fmt"
)

const line = "================================"
const infotext = "PHP Source Code Analyser\nWritten in Go"
const missing_files_warning = "Source code file required as a parameter.\nExample ./sca <file>.php "
const sql = "Possible SQL Errors in the following"
const command = "Possible Command Injection in the following"
const urlredirection = "Possible URL redirection in the following"
const socket = "Possible Socket operations in the following"
const miscinfo = "Possible misc info in the following"
const fileaccess = "Possible file access in the following"
const userinfo = "Possible User supplied info in the following"

func printInfoBanner() {
	fmt.Println(line)
	fmt.Println(infotext)
	fmt.Println(line)
}

func printSQLBanner() {
	fmt.Println(line)
	fmt.Println(sql)
	fmt.Println(line)
}

func printCommandInjectionBanner() {
	fmt.Println(line)
	fmt.Println(command)
	fmt.Println(line)
}

func printURLRedirection() {
	fmt.Println(line)
	fmt.Println(urlredirection)
	fmt.Println(line)
}

func printSocketBanner() {
	fmt.Println(line)
	fmt.Println(socket)
	fmt.Println(line)
}

func printMiscInfoBanner() {
	fmt.Println(line)
	fmt.Println(miscinfo)
	fmt.Println(line)
}

func printFileInfoBanner() {
	fmt.Println(line)
	fmt.Println(fileaccess)
	fmt.Println(line)
}

func printUserInfoBanner() {
	fmt.Println(line)
	fmt.Println(userinfo)
	fmt.Println(line)
}

func usage() {
	fmt.Println(missing_files_warning)
}
