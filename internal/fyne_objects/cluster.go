package fyne_objects

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var instanceTypes = []string{"i3", "m4", "m5", "m5a", "m6i", "c4", "c5", "c5a", "r4", "r5", "r5a", "t3", "t3a"}
var instanceSizes = map[string][]string{
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

var env = []string{"dev ", "sbx ", "infra ", "impl ", "prod ", "mgmt ", "test "}

func clusterSettings() fyne.CanvasObject {

	envTabs := container.NewAppTabs()

	envOpts := widget.NewCheckGroup(env, func(s []string) {

	})

	envOpts.Horizontal = true

	envTabs.Append(
		container.NewTabItem("Global", container.New(
			layout.NewVBoxLayout(),
			widget.NewLabel("Check to enable an environment"),
			envOpts,
			nodeCards(),
			widget.NewButton("Save", func() {}),
		)),
	)

	for _, k := range env {
		envTabs.Append(container.NewTabItem(k, container.New(
			layout.NewVBoxLayout(),
			widget.NewLabel(k),
			nodeCards(),
			widget.NewButton("Save", func() {}),
		)))
	}

	return envTabs
}

func nodeCards() fyne.CanvasObject {

	masterInstanceSize := widget.NewSelectEntry(instanceSizes[""])
	masterInstanceSize.SetPlaceHolder("Choose a size")
	masterInstanceType := widget.NewSelect(instanceTypes, func(s string) {
		masterInstanceSize.SetOptions(instanceSizes[s])
		masterInstanceSize.CreateRenderer().Refresh()
	})

	infraInstanceSize := widget.NewSelectEntry(instanceSizes[""])
	infraInstanceSize.SetPlaceHolder("Choose a size")
	infraInstanceType := widget.NewSelect(instanceTypes, func(s string) {
		infraInstanceSize.SetOptions(instanceSizes[s])
		infraInstanceSize.CreateRenderer().Refresh()
	})

	storageInstanceSize := widget.NewSelectEntry(instanceSizes[""])
	storageInstanceSize.SetPlaceHolder("Choose a size")
	storageInstanceType := widget.NewSelect(instanceTypes, func(s string) {
		storageInstanceSize.SetOptions(instanceSizes[s])
		storageInstanceSize.CreateRenderer().Refresh()
	})

	workerInstanceSize := widget.NewSelectEntry(instanceSizes[""])
	workerInstanceSize.SetPlaceHolder("Choose a size")
	workerInstanceType := widget.NewSelect(instanceTypes, func(s string) {
		workerInstanceSize.SetOptions(instanceSizes[s])
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
