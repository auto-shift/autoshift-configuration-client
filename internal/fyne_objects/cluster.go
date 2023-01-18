package fyne_objects

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func clusterSettings() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItem("Cluster", widget.NewLabel("See TODO ")),
	)

	return tabs
}
