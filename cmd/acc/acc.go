package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/fyne_objects"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/impls"
)

func main() {

	accApp := app.New()

	mainWin := accApp.NewWindow("AutoShift Configuration Client")
	mainWin.Resize(fyne.NewSize(800, 600))

	mainWin.SetMainMenu(fyne_objects.MakeMenu(accApp, mainWin))
	mainWin.SetMaster()

	mainWin.SetContent(fyne_objects.AppMain(mainWin))

	if impls.ReadGitConfigs().GitDir == "Not Set" {
		fyne_objects.GitConfEditDialog(mainWin)
	}

	mainWin.Show()

	accApp.Run()
}
