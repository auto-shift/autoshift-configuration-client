package fyne_objects

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/utils"
)

func Gitops(win fyne.Window) fyne.CanvasObject {
	//Read vars file for git data
	utils.GetGitVars()

	gitCard := widget.NewCard("Git Settings", "", makeGitFormEdit(win))

	return gitCard
}

//
func makeGitFormEdit(win fyne.Window) fyne.CanvasObject {

	edit := widget.NewButton(
		"edit",
		func() {

			GitConfDialog(win)

		})

	return edit
}

func GitConfDialog(win fyne.Window) {

	label := widget.NewLabel("Git Settings")
	label.Alignment = fyne.TextAlignCenter

	userName := widget.NewEntry()
	userName.SetPlaceHolder("gituser")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	gitRepo := widget.NewEntry()
	gitRepo.SetPlaceHolder("https://github.com/auto-shift/autoshift.git")
	remember := false

	items := []*widget.FormItem{
		widget.NewFormItem("UserName:", userName),
		widget.NewFormItem("Password:", password),
		widget.NewFormItem("Repository URL:", gitRepo),
	}

	dialog.ShowForm("Edit Git Settings", "Confirm", "Cancel", items, func(b bool) {
		if !b {
			return
		}
		var rememberText string
		if remember {
			rememberText = "and remember this login"
		}

		log.Println("Please Authenticate", userName.Text, password.Text, rememberText)
	}, win)

}
