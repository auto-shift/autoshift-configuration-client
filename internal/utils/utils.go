package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

var LogFileLoc string
var LogFile = new(os.File)

func init() {
	currDateTime := time.Now().String()
	logLoc, err := filepath.Abs("./logs")
	if err != nil {
		fmt.Println("logloc error: ")
		fmt.Println(err)
	}
	fmt.Println("logloc: ")
	fmt.Println(logLoc)

	LogFileLoc = logLoc + "/" + currDateTime

	fmt.Println(LogFileLoc)

	LogFile, err = os.OpenFile(LogFileLoc, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("LogFile error: ")
		fmt.Println(err)
	}

	log.SetOutput(LogFile)

	log.Println("Welcome to AutoShift Configuration Client")
}

func GetLogs() string {
	yfile, err := ioutil.ReadFile(LogFileLoc)

	if err != nil {

		log.Fatal(err)
	}

	var data string
	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {

		log.Fatal(err2)
	}

	return data

}
