package main

import  "strings"

func main () {
	// bjc-qtt-base-preview-1  -> qtt-base-preview
	ServerName := "bjc-qtt-base-preview-1"

	firstIndex := strings.Index(ServerName, "-") + 1
	println(firstIndex)
	lastIndex := strings.LastIndex(ServerName, "-")
	println(lastIndex)

	if firstIndex >= lastIndex {
		firstIndex = 0
		lastIndex = len(ServerName)
	}
	SimpleName := ServerName[firstIndex:lastIndex]
	println(SimpleName)
}
