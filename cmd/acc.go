package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
	"github.com/auto-shift/autoshift-configuration-client/cmd/internal/fyne_objects"
	"github.com/auto-shift/autoshift-configuration-client/cmd/internal/utils"
)

func main() {
	// TODO: Function to create or read vars file. File will only contain information used locally

	accApp := app.New()

	accWindow := accApp.NewWindow("AutoShift Configuration Client")
	accWindow.Resize(fyne.NewSize(600, 600))

	accWindow.SetMainMenu(utils.MakeMenu(accApp, accWindow))
	accWindow.SetMaster()

	top := widget.NewLabel("")
	right := widget.NewLabel("")
	bottom := canvas.NewText("For more information visit github.com/auto-shift", color.Black)
	bottom.Alignment = fyne.TextAlignCenter
	bottom.TextStyle = fyne.TextStyle{Italic: true}

	middle := container.NewAppTabs(
		//TODO: Tab should allow configuration of cluster environment variables
		container.NewTabItem("Cluster", widget.NewLabel("See TODO ")),
		//TODO: Tab should allow configuration git credentials
		container.NewTabItem("Gitops", fyne_objects.Gitops(accWindow)),
		//TODO: Tab should allow configuration available CICD options
		container.NewTabItem("CICD", widget.NewLabel("see TODO")),
		//TODO: Tab should allow create a top oriented AppTabs object with a tab for each service.
		container.NewTabItem("Services", widget.NewLabel("See TODO")),
	)
	middle.SetTabLocation(container.TabLocationLeading)

	content := container.NewBorder(top, bottom, nil, right, middle)
	accWindow.SetContent(content)
	accWindow.ShowAndRun()

}
