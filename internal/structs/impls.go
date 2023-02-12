package structs

import (
	"io/ioutil"
	"log"

	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/impls"
	"gopkg.in/yaml.v3"
)

func ReadArgoVars() Gitops_Vars {
	gitDir := impls.ReadGitConfigs().GitDir
	yfile, err := ioutil.ReadFile(gitDir + "/autoShift/vars/all/main.yml")
	if err != nil {

		log.Fatal(err)
	}

	var gitOpsVars Gitops_Vars

	err2 := yaml.Unmarshal(yfile, &gitOpsVars)
	if err2 != nil {
		panic(err2)
	}
	return gitOpsVars

}
