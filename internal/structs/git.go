package structs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type GitMap map[string]GVars
type GitVars struct {
	GMap GitMap `yaml:"gitVars"`
	//Todo, make it an array of gitMaps so user can choose between multiple repos
}
type GVars struct {
	GitDir  string `yaml:"git-Dir"`
	GitUrl  string `yaml:"git-Url"`
	GitUser string `yaml:"git-User"`
}

// getters
func (vars GVars) GetDir() string {
	return vars.GitDir
}

func (vars GVars) GetUrl() string {
	return vars.GitUrl
}

func (vars GVars) GetUser() string {
	return vars.GitUser
}

// setters
func (vars GitVars) UpdateGitVars(gVars GVars) {
	gMap := make(map[string]GVars)
	gMap["gitVars"] = gVars
	vars.GMap = gMap

}

func (vars GVars) SetDir(dir string) {
	vars.GitDir = dir
}

func (vars GVars) SetUser(user string) {
	vars.GitUser = user
}

func (vars GVars) SetUrl(url string) {
	vars.GitUrl = url
}

// methods
func (gVars GitVars) ReadGitConfigs() GVars {

	confPath, err := filepath.Abs("../../configs/vars.yml")
	if err != nil {
		fmt.Println(err)
	}

	yfile, err := os.ReadFile(confPath)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	err2 := yaml.Unmarshal([]byte(yfile), &gVars.GMap)
	if err2 != nil {
		fmt.Println("err2:")
		fmt.Println(err2)
	}
	fmt.Println("test: ")
	fmt.Println(gVars.GMap)

	return gVars.GMap["gitVars"]
}

func (gVars GitVars) WriteGitConfigs() {
	yEdits, err := yaml.Marshal(gVars.GMap)
	if err != nil {
		log.Println(err)
	}
	os.WriteFile("../../configs/vars.yml", yEdits, 0644)
	fmt.Println(gVars.ReadGitConfigs())
}
