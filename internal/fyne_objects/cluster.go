package fyne_objects

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/utils"
)

// sets the logger output file

func init() {
	log.SetOutput(utils.LogFile)
}

//AWS Instance Types
var awsInstanceTypes = []string{"i3", "m4", "m5", "m5a", "m6i", "c4", "c5", "c5a", "r4", "r5", "r5a", "t3", "t3a"}
var awsInstanceSizes = map[string][]string{
	"i3":  {"large"},
	"m4":  {"large", "xlarge", "2xlarge", "4xlarge", "10xlarge", "16xlarge"},
	"m5":  {"large", "xlarge", "2xlarge", "4xlarge", "8xlarge", "12xlarge", "16xlarge"},
	"m5a": {"large", "xlarge", "2xlarge", "4xlarge", "8xlarge", "10xlarge", "16xlarge"},
	"m6i": {"xlarge", "2xlarge", "4xlarge", "8xlarge", "16xlarge"},
	"c4":  {"2xlarge", "4xlarge", "8xlarge"},
	"c5":  {"xlarge", "2xlarge", "4xlarge", "9xlarge", "12xlarge", "18xlarge", "24xlarge"},
	"c5a": {"xlarge", "2xlarge", "4xlarge", "8xlarge", "12xlarge", "16xlarge", "24xlarge"},
	"r4":  {"large", "xlarge", "2xlarge", "4xlarge", "8xlarge", "16xlarge"},
	"r5":  {"large", "xlarge", "2xlarge", "4xlarge", "8xlarge", "12xlarge", "16xlarge", "24xlarge"},
	"r5a": {"large", "xlarge", "2xlarge", "4xlarge", "8xlarge", "12xlarge", "16xlarge", "24xlarge"},
	"t3":  {"large", "xlarge", "2xlarge"},
	"t3a": {"large", "xlarge", "2xlarge"},
}

var envType = []string{"Bare Metal", "AWS", "Azure", "Open Stack", "RHV", "VSphere"}

/// Cluster Main UI
var env = map[string]bool{}

var envOpts widget.CheckGroup

var envTypeOpt string

var envTabs container.AppTabs

func clusterSettings() fyne.CanvasObject {

	envEntry := widget.NewEntry()
	newEnvForm := container.New(
		layout.NewFormLayout(),
		widget.NewLabel("Environment Name: "),
		envEntry,
	)

	envBtn := widget.NewButton("add environment", func() {
		if envEntry.Text != "" {
			showConfirmDialog(envEntry.Text)
			envEntry.SetText("")
		}
	})
	envOpts.Horizontal = true

	envRadioGrp := widget.NewRadioGroup(envType, func(s string) {
		envTypeOpt = s
		log.Println("Cluster is being installed on " + s)
	})

	envRadioGrp.Horizontal = true
	envRadioGrp.SetSelected("Bare Metal")
	envRadioGrp.Required = true

	envTabs.Append(
		container.NewTabItem("Global", container.New(
			layout.NewVBoxLayout(),
			widget.NewLabel("Environment Type:"),
			envRadioGrp,
			widget.NewLabel("Create a new environment:"),
			newEnvForm,
			envBtn,
			widget.NewLabel("Check to enable an environment"),
			container.NewHBox(&envOpts),
			nodeCards(),
			widget.NewButton("Save", func() {}),
		)),
	)
	initEnvOpts()
	initTabList()
	return &envTabs
}

//UI functions
func initEnvOpts() {
	selected := []string{}
	for k, v := range env {
		envOpts.Append(k)
		if v {
			selected = append(selected, k)
		}
	}
	envOpts.SetSelected(selected)
}

func updateEnvOpts(e string) {
	envOpts.Append(e)
	envOpts.SetSelected([]string{e})
}

func initTabList() {
	for k, v := range env {
		if v {
			envTabs.Append(container.NewTabItem(k, container.New(
				layout.NewVBoxLayout(),
				widget.NewLabel(k),
				nodeCards(),
				widget.NewButton("Save", func() {}),
			)))
		}
	}
}

func updateTabList(e string) {
	envTabs.Append(container.NewTabItem(e, container.New(
		layout.NewVBoxLayout(),
		widget.NewLabel(e),
		nodeCards(),
		widget.NewButton("Save", func() {}),
	)))
}

func showConfirmDialog(envName string) {
	dialog := dialog.NewConfirm("Confirm New Environment", "Would You like to create Enviroment: "+envName+"?", func(b bool) {
		if b {
			env[envName] = true
			updateEnvOpts(envName)
			updateTabList(envName)
		}
	}, utils.MainWin)
	dialog.Show()
}

func nodeCards() fyne.CanvasObject {

	masterInstanceSize := widget.NewSelectEntry(awsInstanceSizes[""])
	masterInstanceSize.SetPlaceHolder("Choose a size")
	masterInstanceType := widget.NewSelect(awsInstanceTypes, func(s string) {
		masterInstanceSize.SetOptions(awsInstanceSizes[s])
		masterInstanceSize.CreateRenderer().Refresh()
	})

	infraInstanceSize := widget.NewSelectEntry(awsInstanceSizes[""])
	infraInstanceSize.SetPlaceHolder("Choose a size")
	infraInstanceType := widget.NewSelect(awsInstanceTypes, func(s string) {
		infraInstanceSize.SetOptions(awsInstanceSizes[s])
		infraInstanceSize.CreateRenderer().Refresh()
	})

	storageInstanceSize := widget.NewSelectEntry(awsInstanceSizes[""])
	storageInstanceSize.SetPlaceHolder("Choose a size")
	storageInstanceType := widget.NewSelect(awsInstanceTypes, func(s string) {
		storageInstanceSize.SetOptions(awsInstanceSizes[s])
		storageInstanceSize.CreateRenderer().Refresh()
	})

	workerInstanceSize := widget.NewSelectEntry(awsInstanceSizes[""])
	workerInstanceSize.SetPlaceHolder("Choose a size")
	workerInstanceType := widget.NewSelect(awsInstanceTypes, func(s string) {
		workerInstanceSize.SetOptions(awsInstanceSizes[s])
		workerInstanceSize.CreateRenderer().Refresh()
	})

	cards := container.New(
		layout.NewGridLayout(2),
		widget.NewCard("Master Nodes", "",
			container.New(
				layout.NewFormLayout(),
				widget.NewLabel("Instance Type"),
				masterInstanceType,
				widget.NewLabel("Instance Size"),
				masterInstanceSize,
			),
		),
		widget.NewCard("Infra Nodes", "",
			container.New(
				layout.NewFormLayout(),
				widget.NewLabel("Instance Type"),
				infraInstanceType,
				widget.NewLabel("Instance Size"),
				infraInstanceSize,
			),
		),
		widget.NewCard("Storage Nodes", "",
			container.New(
				layout.NewFormLayout(),
				widget.NewLabel("Instance Type"),
				storageInstanceType,
				widget.NewLabel("Instance Size"),
				storageInstanceSize,
			),
		),
		widget.NewCard("Worker Nodes", "",
			container.New(
				layout.NewFormLayout(),
				widget.NewLabel("Instance Type"),
				workerInstanceType,
				widget.NewLabel("Instance Size"),
				workerInstanceSize,
			),
		),
	)
	return cards
}

// func nodeCards(){
//     cards := container.New(
// 		layout.NewGridLayout(2),
// 		widget.NewCard("Master","",
// 		    container.NewForm(
// 				widget.NewLabel("Instance Type:"),
// 				selectInstanceType,
// 				widget.NewLabel("Instance Size:"),
//                 selectInstanceSize,
// 			),
// 		),
// 	)
// }
