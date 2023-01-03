package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	git "github.com/go-git/go-git/v5"
	http "github.com/go-git/go-git/v5/plumbing/transport/http"
)

//vars
// structs
type GitVars struct {
	GitDir  string `yaml:"gitDir"`
	GitUrl  string `yaml:"gitUrl"`
	GitUser string `yaml:"gitUser"`
	GitPass string `yaml:"gitPass"`
}

//Methods for interacting with a git repository
func GitClone(gitUser, gitPass, gitDir, gitUrl string) {

	r, err := git.PlainClone(gitDir, false, &git.CloneOptions{
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth: &http.BasicAuth{
			Username: gitUser, // yes, this can be anything except an empty string
			Password: string(gitPass),
		},
		URL:      gitUrl,
		Progress: os.Stdout,
	})
	CheckIfError(err)

	ref, err := r.Head()
	CheckIfError(err)

	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	if os.IsNotExist(err) {
		println(err)
	}
	fmt.Println(commit)

}

func GitPull(repo, branch string) {

}

func GitPush(repo, branch string) {

}

// read git configs
func ReadGitConfigs() GitVars {

	yfile, err := ioutil.ReadFile("../../configs/vars.yml")
	if err != nil {

		log.Fatal(err)
	}

	var gitVars GitVars

	err2 := yaml.Unmarshal(yfile, &gitVars)
	if err2 != nil {
		panic(err2)
	}
	return gitVars

}

func WriteGitConfigs(gitEdits GitVars) {
	yEdits, err := yaml.Marshal(gitEdits)
	if err != nil {
		log.Println(err)
	}
	os.WriteFile("../../configs/vars.yml", yEdits, 0644)

}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
