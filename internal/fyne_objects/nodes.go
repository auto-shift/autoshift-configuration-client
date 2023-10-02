package fyne_objects

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/data"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/structs"
)

func nodeCards(env string, nodes structs.AllNodes) fyne.CanvasObject {

	confs := data.Confs.GetEnvConfigs(env)

	//master
	masterInstanceSize := widget.NewSelectEntry(confs.GetInstanceTypes())
	masterInstanceSize.SetPlaceHolder("Choose a size")
	masterInstanceType := widget.NewSelect(confs.GetInstanceTypes(), func(s string) {
		masterInstanceSize.SetOptions(confs.GetInstanceSizes()[s])
		masterInstanceSize.CreateRenderer().Refresh()
	})
	masterReplicas := widget.NewEntry()

	//infra
	infraInstanceSize := widget.NewSelectEntry(confs.GetInstanceTypes())
	infraInstanceSize.SetPlaceHolder("Choose a size")
	infraInstanceType := widget.NewSelect(confs.GetInstanceTypes(), func(s string) {
		infraInstanceSize.SetOptions(confs.GetInstanceSizes()[s])
		infraInstanceSize.CreateRenderer().Refresh()
	})
	infraReplicas := widget.NewEntry()

	//infra
	storageInstanceSize := widget.NewSelectEntry(confs.GetInstanceTypes())
	storageInstanceSize.SetPlaceHolder("Choose a size")
	storageInstanceType := widget.NewSelect(confs.GetInstanceTypes(), func(s string) {
		storageInstanceSize.SetOptions(confs.GetInstanceSizes()[s])
		storageInstanceSize.CreateRenderer().Refresh()
	})
	storageReplicas := widget.NewEntry()
	//worker
	workerInstanceSize := widget.NewSelectEntry(confs.GetInstanceTypes())
	workerInstanceSize.SetPlaceHolder("Choose a size")
	workerInstanceType := widget.NewSelect(confs.GetInstanceTypes(), func(s string) {
		workerInstanceSize.SetOptions(confs.GetInstanceSizes()[s])
		workerInstanceSize.CreateRenderer().Refresh()
	})
	workerReplicas := widget.NewEntry()

	cards := container.New(
		layout.NewGridLayout(2),

		widget.NewCard("Master Nodes", "",
			container.New(
				layout.NewFormLayout(),
				widget.NewLabel("Instance Type"),
				masterInstanceType,
				widget.NewLabel("Instance Size"),
				masterInstanceSize,
				masterReplicas,
			),
		),
		widget.NewCard("Infra Nodes", "",
			container.New(
				layout.NewFormLayout(),
				widget.NewLabel("Instance Type"),
				infraInstanceType,
				widget.NewLabel("Instance Size"),
				infraInstanceSize,
				infraReplicas,
			),
		),
		widget.NewCard("Storage Nodes", "",
			container.New(
				layout.NewFormLayout(),
				widget.NewLabel("Instance Type"),
				storageInstanceType,
				widget.NewLabel("Instance Size"),
				storageInstanceSize,
				storageReplicas,
			),
		),
		widget.NewCard("Worker Nodes", "",
			container.New(
				layout.NewFormLayout(),
				widget.NewLabel("Instance Type"),
				workerInstanceType,
				widget.NewLabel("Instance Size"),
				workerInstanceSize,
				workerReplicas,
			),
		),
	)
	return cards
}
