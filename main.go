package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// getBackupPath gets path to store backups
func getBackupPath(dirName string) string{
	backupPath := os.Getenv("GOPATH") // Init to GOPATH
	dirSlice := []string{"src", "github.com", "AndrewYinLi", "QuitHub"} // Path to QuitHub src code
	for _,dir := range dirSlice{
		backupPath = path.Join(backupPath, dir)
	}
	backupPath = path.Join(backupPath, dirName)
	return backupPath
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
	// Get args
	cd,_ := os.Getwd()
	baseName := filepath.Base(cd)
	commitName := os.Args[1]
	if len(os.Args) == 2{
		commitName = os.Args[2]
	}
	// Determine action
	if os.Args[1] == "commit"{
		commit(baseName, commitName)
	} else if os.Args[1] == "revert"{
		revert(baseName, commitName)
	} else if os.Args[1] == "history"{
		history(baseName)
	} else if os.Args[1] == "delete"{
		delete(baseName, commitName)
	}
}
