package fyne_objects

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var dayOneAppList = []string{"ArgoCD", "OpenShift Data Foundations"}
var dayTwoAppList = []string{"AMQ", "AAP", "Container Security Operator", "Cert-Manager", "Red Hat Advanced Cluster Manager", "Red Hat Advanced Cluster Security", "API for Data Protection", "Red Hat Code Ready Workspaces", "Red Hat Serverless"}

func DayOneSettings(win fyne.Window) fyne.CanvasObject {

	apps := container.New(
		layout.NewGridLayout(2),
	)
	for _, k := range dayOneAppList {
		apps.Add(appGroup(k, win))
	}

	return container.New(layout.NewVBoxLayout(), widget.NewLabel("Check to enable an application"), apps, layout.NewSpacer())
}

func DayTwoSettings(win fyne.Window) fyne.CanvasObject {

	apps := container.New(
		layout.NewGridLayout(2),
	)
	for _, k := range dayTwoAppList {
		apps.Add(appGroup(k, win))
	}

	return container.New(layout.NewVBoxLayout(), widget.NewLabel("Check to enable an application"), apps, layout.NewSpacer())
}

func appGroup(appName string, win fyne.Window) fyne.CanvasObject {

	items := getArgoCDForm()
	editBtn := widget.NewButton("Edit", func() {

		dialog.ShowForm(appName, "Submit", "Cancel",
			items,
			func(b bool) {
				if !b {
					return
				}

			}, win)
	})
	editBtn.Disable()
	group := container.New(
		layout.NewHBoxLayout(),
		widget.NewCheck(appName, func(b bool) {
			if b {
				editBtn.Enable()
			} else {
				editBtn.Disable()
			}
		}),
		layout.NewSpacer(),
		editBtn,
	)
	return group
}
