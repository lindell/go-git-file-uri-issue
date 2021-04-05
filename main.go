package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

func main() {
	wd, err := os.Getwd()
	handleErr(err)
	// Deletes repositories to clone if already existing
	cleanup(wd)

	firstRepoLocalPath, err := ioutil.TempDir(os.TempDir(), "go-git-first-clone")
	handleErr(err)

	remoteURL := "https://github.com/go-git/go-git"
	fmt.Printf("Cloning from remote URL: %s\n", remoteURL)
	_, err = git.PlainClone(firstRepoLocalPath, false, &git.CloneOptions{
		URL:      remoteURL,
		Progress: os.Stdout,
	})
	handleErr(err)

	secondRepoLocalPath := filepath.Join(wd, "go-git-second-clone")
	firstRepoFileURI := fmt.Sprintf(`file://"%s"`, filepath.ToSlash(firstRepoLocalPath))
	// In case there are spaces in the file URI
	firstRepoFileURI = strings.ReplaceAll(firstRepoFileURI, " ", "%20")
	fmt.Printf("Cloning from local URI: %s\n", firstRepoFileURI)
	_, err = git.PlainClone(secondRepoLocalPath, false, &git.CloneOptions{
		URL:      firstRepoFileURI,
		Progress: os.Stdout,
	})
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func cleanup(wd string) {
	os.RemoveAll(filepath.Join(wd, "go-git-first-clone"))
	os.RemoveAll(filepath.Join(wd, "go-git-second-clone"))
}
