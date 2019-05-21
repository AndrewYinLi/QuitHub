package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// getSourcePath returns the path to where the src code for `QuitHub` is located
func getSourcePath() string{
	sourcePath := os.Getenv("GOPATH") // Init to GOPATH
	dirSlice := []string{"src", "github.com", "AndrewYinLi", "QuitHub"} // Path to QuitHub src code
	for _,dir := range dirSlice{
		sourcePath = path.Join(sourcePath, dir)
	}
	return sourcePath
}

// Commit to the history for dirName a copy of the cwd renamed as commitName
func commit(dirName string, commitName string){

}

// Revert the contents of the cwd to the contents of commitName stored in the history for dirName
func revert(dirName string, commitName string){

}

// Print the history of commits for dirName
func history(dirName string){

}

// Delete the directory commitName committed to dirName
func delete(dirName string, commitName string){

}

func main() {
	cd,_ := os.Getwd()
	base := filepath.Base(cd)
	if os.Args[1] == "commit"{
		if len(os.Args) < 2{ // Error
			// os.Args[2] is the new name
		} else{

		}
	} else if os.Args[1] == "revert"{

	} else if os.Args[1] == "history"{

	} else if os.Args[1] == "delete"{
		if len(os.Args) < 2{ // Error, no name for deletion

		} else{

		}
	}
}
