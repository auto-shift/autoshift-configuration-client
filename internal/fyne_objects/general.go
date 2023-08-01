package fyne_objects

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/utils"
)

func init() {
	log.SetOutput(utils.LogFile)
}

func generalSettings(win fyne.Window) fyne.CanvasObject {

	kubeAdmminCheck := widget.NewCheck("Remove Kubeadmin", func(b bool) {
		if b {
			remove_kubeadmin = true
		} else {
			remove_kubeadmin = false
		}
	})

	cicdRadio := widget.NewRadioGroup([]string{"Tekton", "Jenkins"}, func(value string) {
		log.Println("Radio set to", value)
	})

	cicdRadio.Horizontal = true

	settings := container.New(
		layout.NewVBoxLayout(),
		widget.NewCard("Git Settings", "",
			GitSettings(win)),
		widget.NewCard("CICD", "",
			cicdRadio,
		),
		widget.NewCard("Installation Settings", "",
			container.NewGridWithColumns(3,
				kubeAdmminCheck,
			),
		),
	)
	return settings
}
