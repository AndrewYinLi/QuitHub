package main

import (
	"os"
	"path"
)

// getProgramPath returns the path to where the src code for `QuitHub` is located
func getSourcePath() string{
	sourcePath := os.Getenv("GOPATH") // Init to GOPATH
	dirSlice := []string{"src", "github.com", "AndrewYinLi", "QuitHub"} // Path to QuitHub src code
	for _,dir := range dirSlice{
		sourcePath = path.Join(sourcePath, dir)
	}
	return sourcePath
}

func main() {
	
}
