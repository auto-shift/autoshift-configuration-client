package fyne_objects

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type envState struct {
	devEnabled   bool
	implEnabled  bool
	infraEnabled bool
	mgmtEnabled  bool
	sbxEnabled   bool
	testEnabled  bool
	prodEnabled  bool
}

func EnvSettings() fyne.CanvasObject {

	allTab := container.NewTabItem("Global", globalTab())
	devTab := container.NewTabItem("Dev", allEnvSettings())
	implTab := container.NewTabItem("Impl", allEnvSettings())
	infraTab := container.NewTabItem("Infra", allEnvSettings())
	mgmtTab := container.NewTabItem("Mgmt", allEnvSettings())
	sbxTab := container.NewTabItem("Sbx", allEnvSettings())
	testTab := container.NewTabItem("Test", allEnvSettings())
	prodTab := container.NewTabItem("Prod", allEnvSettings())

	tabs := container.NewAppTabs()
	tabs.SetTabLocation(container.TabLocationTop)
	return tabs
}

func allEnvSettings() fyne.CanvasObject {

	ac := widget.NewAccordion()

	ac.Append(widget.NewAccordionItem("D", &widget.Entry{Text: "Four"}))

	return ac
}

func globalTab() fyne.CanvasObject {
	envs := []string{"Dev", "Impl", "Infra", "Mgmt", "Sbx", "Test", "Prod"}
	var envChecks []fyne.CanvasObject

	for i, e := range envs {
		envChecks[i] = widget.NewCheck("Enable "+e+"?: ", func(b bool) {
			if b {

			}
		})
	}

	vBox := container.New(
		layout.NewVBoxLayout(),
		container.New(
			layout.NewHBoxLayout(),
		),
		widget.NewLabel("Global settings apply to all enabled environments"),
		allEnvSettings(),
	)

	return vBox
}

func envTab(env string) fyne.CanvasObject {

	vBox := container.New(
		layout.NewVBoxLayout(),
		allEnvSettings(),
	)

	return vBox
}
