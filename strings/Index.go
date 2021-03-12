package main

import (
	"fmt"
	"strings"
)

func main() {
	//bjc-base-preview-1  -> base-preview
	serverName := "bjc-base-preview-1"
	firstIndex := strings.Index(serverName, "-") + 1
	lastIndex := strings.LastIndex(serverName, "-")
	fmt.Printf("firstIndex: %+v\n", firstIndex)
	fmt.Printf("lastIndex: %+v\n", lastIndex)
	// 这种情况：base-preview
	if firstIndex >= lastIndex {
		firstIndex = 0
		lastIndex = len(serverName)
	}
	serverSimpleName := serverName[firstIndex:lastIndex]
	fmt.Printf("%+v\n", serverSimpleName)
}
