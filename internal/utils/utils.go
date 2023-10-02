package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2"
)

var LogFileLoc string
var LogFile = new(os.File)
var MainWin fyne.Window

func init() {

	currDateTime := time.Now().String()
	accPath := "./logs"
	if _, err := os.Stat(accPath); os.IsNotExist(err) {
		err := os.Mkdir(accPath, 0777)
		fmt.Println("creating log directory")
		log.Println(err)
	}

	logLoc, err := filepath.Abs(accPath)
	if err != nil {
		fmt.Println("log path error: ")
		fmt.Println(err)
	}

	LogFileLoc = logLoc + "/" + currDateTime

	LogFile, err = os.OpenFile(LogFileLoc, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("LogFile error: ")
		fmt.Println(err)
	}

	log.SetOutput(LogFile)

	log.Println("Welcome to AutoShift Configuration Client")
}

func GetLogs() []string {
	//yfile, err := ioutil.ReadFile(LogFileLoc)

	file, err := os.Open(LogFileLoc)
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	logs := []string{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}

	return logs

}

// func GetString(bs binding.String) string {
// 	logs, err := bs.Get()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return logs
// }
