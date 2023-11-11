package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/home/ui"
)

func main() {

	accApp := app.New()

	mainWin := accApp.NewWindow("AutoShift Configuration Client")
	mainWin.SetFullScreen(false)
	mainWin.Resize(fyne.NewSize(1280, 900))

	mainWin.SetMainMenu(ui.MakeMenu(accApp, mainWin))
	mainWin.SetMaster()

	mainWin.SetContent(ui.AppMain(mainWin))

	mainWin.Show()

	accApp.Run()
}
