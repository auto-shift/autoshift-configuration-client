package fyne_objects

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/data"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/structs"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/utils"
)

// sets the logger output file

func init() {
	log.SetOutput(utils.LogFile)

}

// vars
var tlDomain = "localhost"
var envTypeSel = "Bare Metal"
var platformList = data.Platform
var envVars = structs.TabVars{}
var envList = map[string]bool{}
var envTypeOpt string
var tabList = structs.TabList{}

// / Cluster Main UI
var envOpts widget.CheckGroup
var envTabs container.AppTabs

// main cluster function
func clusterSettings() fyne.CanvasObject {

	domainBind := binding.NewString()
	domainBind.Set(tlDomain)
	domainEntry := widget.NewEntryWithData(domainBind)
	domainForm := widget.NewCard("Domain: ", "",
		container.New(
			layout.NewFormLayout(),
			widget.NewLabel("Top Level Domain: "),
			domainEntry,
		),
	)

	envEntry := widget.NewEntry()

	envBtn := widget.NewButton("add environment", func() {
		if envEntry.Text != "" {
			showConfirmDialog(envEntry.Text)
			envEntry.SetText("")
		}
	})
	envOpts.Horizontal = true

	envSelect := widget.NewSelect(platformList, func(s string) {
		envTypeOpt = s
		log.Println("Cluster is being installed on " + s)
	})

	newEnvForm := container.New(
		layout.NewFormLayout(),
		widget.NewLabel("Environment Type:"),
		envSelect,
		widget.NewLabel("Environment Name: "),
		envEntry,
	)

	var err error
	envTypeBind := binding.NewString()
	envTypeBind.Set(envTypeSel)
	envSelect.Selected, err = envTypeBind.Get()
	if err != nil {
		log.Println(err)
	}

	newClusterCard := widget.NewCard("Add a new Environment:", "",
		container.New(
			layout.NewVBoxLayout(),
			newEnvForm,
			envBtn,
		),
	)

	envCheckGrpCard := widget.NewCard("Environments", "", container.New(
		layout.NewVBoxLayout(),
		widget.NewLabel("Check to enable an environment"),
		container.NewHBox(&envOpts)),
	)

	line := canvas.NewLine(color.Black)
	line.StrokeWidth = 5

	envTabs.Append(
		container.NewTabItem("Global", container.New(
			layout.NewVBoxLayout(),
			domainForm,
			envCheckGrpCard,
			newClusterCard,
			line,
			widget.NewButton("Save", func() {}),
		),
		))

	initEnvOpts()
	initTabList()
	return &envTabs
}

// UI functions
func initEnvOpts() {
	selected := []string{}
	for k, v := range envList {
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

// Function creates new tabs
func updateTabList(eVars structs.TabVars) {
	env := eVars.GetEnv()

	envTabs.Append(container.NewTabItem(env, container.New(
		layout.NewVBoxLayout(),
		widget.NewLabel(env),
		nodeCards(env, eVars.GetNodes()),
		widget.NewButton("Save", func() {}),
	)))
}

// creates the global tab and tabs for already created environments
func initTabList() {
	for k, v := range envList {
		if v {
			updateTabList(findTabByEnv(k))

		}
	}
}

// confirms new environment tab creation
func showConfirmDialog(envName string) {
	dialog := dialog.NewConfirm("Confirm New Environment", "Would You like to create Enviroment: "+envName+"?", func(b bool) {
		if b {
			envList[envName] = true
			updateEnvOpts(envName)
			//create tab and add to updateTabList

			updateTabList(findTabByEnv(envName))
		}
	}, utils.MainWin)
	dialog.Show()
}

func findTabByEnv(env string) structs.TabVars {
	tab := tabList.SearchTabs(env)
	if (tab == structs.TabVars{}) {
		tab.SetEnv(env)
	}
	return tab
}
