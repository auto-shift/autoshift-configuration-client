package fyne_objects

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func AppMain(win fyne.Window) fyne.CanvasObject {

	top := widget.NewLabel("")
	right := widget.NewLabel("")
	bottom := canvas.NewText("For more information visit github.com/auto-shift", color.Black)
	bottom.Alignment = fyne.TextAlignCenter
	bottom.TextStyle = fyne.TextStyle{Italic: true}

	middle := container.NewAppTabs(
		//TODO: Tab should allow configuration git credentials
		container.NewTabItem("Gitops", GitSettings(win)),
		//TODO: Tab should allow configuration of cluster environment variables
		container.NewTabItem("Cluster", widget.NewLabel("See TODO ")),
		//TODO: Tab should allow configuration available CICD options
		container.NewTabItem("CICD", widget.NewLabel("see TODO")),
		//TODO: Tab should allow create a top oriented AppTabs object with a tab for each service.
		container.NewTabItem("Services", widget.NewLabel("See TODO")),
	)
	middle.SetTabLocation(container.TabLocationLeading)

	content := container.NewBorder(top, bottom, nil, right, middle)

	return content
}
