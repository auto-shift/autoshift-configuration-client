package utils

import (
	"fmt"

	billy "github.com/go-git/go-billy/v5"
	git "github.com/go-git/go-git/v5"
	http "github.com/go-git/go-git/v5/plumbing/transport/http"
	memory "github.com/go-git/go-git/v5/storage/memory"
)

// git vars

var storer *memory.Storage
var fs billy.Filesystem

//Methods for interacting with a git repository

func GitClone(gitUser, gitPass, gitRepo string) {
	auth := &http.BasicAuth{
		Username: gitUser,
		Password: gitPass,
	}

	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL:  gitRepo,
		Auth: auth,
	})

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	fmt.Println("Repository cloned")

	w, err := r.Worktree()

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

}

func GitPull(repo, branch string) {

}

func GitPush(repo, branch string) {

}

//TODO: reads git variables from cloned repository
func GetGitVars() {

}
