package impls

import (
	"log"
	"os"

	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/structs"
	"gopkg.in/yaml.v3"
)

var Tabs = structs.TabVars{}

func readClusterInfo(dir string) {
	yfile, err := os.ReadFile(dir + "/autoShift/vars/all/main.yml")
	handleErr(err)
	readClusterVars(yfile)

}

func readClusterVars(file []byte) {
	var e structs.Env
	err := yaml.Unmarshal(file, &e)
	handleErr(err)
}

func readNodes(file []byte) {
	var master structs.MasterNode
	err := yaml.Unmarshal(file, &master)
	handleErr(err)

	var infra structs.InfraNode
	err1 := yaml.Unmarshal(file, &infra)
	handleErr(err1)

	var worker structs.WorkerNode
	err2 := yaml.Unmarshal(file, &worker)
	handleErr(err2)

	var storage structs.StorageNode
	err3 := yaml.Unmarshal(file, &storage)
	handleErr(err3)

}

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
