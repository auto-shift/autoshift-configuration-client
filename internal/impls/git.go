package impls

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
}

//Methods for interacting with a git repository
func GitClone(gitUser, gitPass, gitDir, gitUrl string) {
	// var resp []string

	path := gitDir + "/autoShift"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0775)
		// TODO: handle error
		fmt.Println(err)
	}

	r, err := git.PlainClone(path, false, &git.CloneOptions{
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth: &http.BasicAuth{
			Username: gitUser, // yes, this can be anything except an empty string
			Password: gitPass,
		},
		URL:      gitUrl,
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println("err1")
		fmt.Println(err)
	}

	ref, err := r.Head()
	if err != nil {
		fmt.Println("err2")
		fmt.Println(err)
	}

	err = r.Storer.SetReference(ref)
	CheckIfError(err)

	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	if os.IsNotExist(err) {

		// resp[0] = "Clone Failed"
		// resp[1] = fmt.Sprintln(err)
		// return resp
		fmt.Println(err)
	}

	// resp[0] = "Clone Successful"
	// resp[1] = fmt.Sprintln(commit)
	// return resp
	fmt.Println(commit)

}

// func gitBranch() {

// }

// func gitCheckout() {

// }

// func GitPull(repo, branch string) {

// }

// func GitPush(repo, branch string) {

// }

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
