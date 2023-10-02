package impls

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/utils"
)

func init() {
	logFile := utils.LogFile

	log.SetOutput(logFile)
}

func ReadVars(envNames []string) {

	if IsLocalRepo() {
		accVars := GitConfs.GetDir() + "/vars/accVars"

		if _, err := os.Stat(accVars); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(accVars, os.ModePerm)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("Vars directory created at " + accVars)
			}
		}

		for _, k := range envNames {
			_, err := os.Create(accVars + "/" + k + ".yml")
			if err != nil {
				fmt.Println(err)
			} else {
				log.Println(k + ".yml " + "created")
			}
		}

	}
}
