package impls

import (
	"fmt"
	"log"
	"os"

	// "github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/utils"

	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/structs"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/utils"
	git "github.com/go-git/go-git/v5"
	http "github.com/go-git/go-git/v5/plumbing/transport/http"
)

func init() {
	logFile := utils.LogFile
	log.SetOutput(logFile)
	GitConfs = structs.GitVars{}.ReadGitConfigs()
}

// func ReadGitConfigs() structs.GVars {

// 	confPath, err := filepath.Abs("../../configs/vars.yaml")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	yfile, err := os.ReadFile(confPath)
// 	if err != nil {
// 		fmt.Println(err)
// 		log.Fatal(err)
// 	}
// 	fmt.Println(yfile)

// 	var gTest map[string]structs.GVars
// 	err2 := yaml.Unmarshal([]byte(yfile), &gTest)
// 	if err2 != nil {
// 		fmt.Println("err2:")
// 		fmt.Println(err2)
// 	}
// 	fmt.Println("test: ")
// 	fmt.Println(gTest)

// 	return gTest["gitVars"]
// }

// vars
var GitConfs structs.GVars

// structs

// Methods for interacting with a git repository
func GitClone(gitUser, gitPass, gitDir, gitUrl string) {
	// var resp []string

	log.Println("cloning " + gitUrl)

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
			Password: gitPass, // Requires access token, not password
		},
		URL:      gitUrl,
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println("err1")
		fmt.Println(err)
		log.Println(err)
	} else {
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
		log.Println(commit)
	}

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
// func ReadGitConfigs() {

// 	yfile, err := os.ReadFile("../../configs/vars.yml")
// 	if err != nil {

// 		log.Fatal(err)
// 	}

// 	var gconfs structs.GitVars
// 	err2 := yaml.Unmarshal(yfile, &gconfs)
// 	if err2 != nil {
// 		panic(err2)
// 	}
// 	fmt.Println("gconfs:")
// 	fmt.Println(gconfs)
// }

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func IsLocalRepo() bool {
	_, err := os.Stat(GitConfs.GetDir() + "/autoShift")
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
