package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/utils"
)

// form for viewing and editing git settings
func GitSettings(win fyne.Window) fyne.CanvasObject {

	gitInfo := utils.ReadGitConfigs()

	gitDir := gitInfo.GitDir
	dir := binding.BindString(&gitDir)
	gitDirHBox := container.New(
		layout.NewHBoxLayout(),
		widget.NewLabelWithStyle("Git Directory:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithData(dir),
	)

	gitUser := gitInfo.GitUser
	user := binding.BindString(&gitUser)
	gitUserHBox := container.New(
		layout.NewHBoxLayout(),
		widget.NewLabelWithStyle("Git Username:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithData(user),
	)

	gitUrl := gitInfo.GitUrl
	url := binding.BindString(&gitUrl)
	gitUrlHBox := container.New(
		layout.NewHBoxLayout(),
		widget.NewLabelWithStyle("Git Remote Repository:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithData(url),
	)

	gitPass := widget.NewPasswordEntry()

	// gitBranch := widget.NewEntry()

	items := []*widget.FormItem{
		widget.NewFormItem("Password:", gitPass),
		// widget.NewFormItem("Branch Name:", gitBranch),
	}
	//buttons for performing git commands
	gitops := container.New(
		layout.NewHBoxLayout(),
		widget.NewButton(
			"Clone Remote Repo",
			func() {

				dialog.ShowForm("Please provide:", "Submit", "Cancel",
					items,
					func(b bool) {
						if !b {
							return
						}
						utils.GitClone(gitUser, gitPass.Text, gitDir, gitUrl)
						// gitResp := utils.GitClone(gitUser, gitPass.Text, gitDir, gitUrl)
						// dialog.ShowInformation(gitResp[0], gitResp[1], win)
					}, win)
			},
		),
		widget.NewButton(
			"Push to Remote Repo",
			func() {

				GitConfEditDialog(win)

			},
		),
		widget.NewButton(
			"Edit Git Settings",
			func() {

				GitConfEditDialog(win)

			},
		),
	)

	gitInfoCont := container.New(layout.NewVBoxLayout(), gitDirHBox, gitUserHBox, gitUrlHBox, gitops)

	return gitInfoCont
}

// Pop up dialog for editing git settings
func GitConfEditDialog(win fyne.Window) {

	gitSettings := utils.ReadGitConfigs()
	var gitLoc string

	label := widget.NewLabel("Git Settings")
	label.Alignment = fyne.TextAlignCenter

	gitUser := gitSettings.GitUser
	user := binding.BindString(&gitUser)
	userName := widget.NewEntry()
	userName.Bind(user)

	gitUrl := gitSettings.GitUrl
	url := binding.BindString(&gitUrl)
	gitRepo := widget.NewEntry()
	gitRepo.Bind(url)

	getDirLoc := func(list fyne.ListableURI, err error) {
		if err != nil {
			panic(err)
		}
		if list == nil {
			log.Println("Cancelled")
		}

		gitLoc = list.Path()
	}

	repoDir := widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {
		dialog.ShowFolderOpen(getDirLoc, win)
	})

	items := []*widget.FormItem{
		widget.NewFormItem("UserName:", userName),
		widget.NewFormItem("Repository URL:", gitRepo),
		widget.NewFormItem("choose a repository location:", repoDir),
	}

	form := dialog.NewForm("Edit Git Settings", "Confirm", "Cancel", items, func(b bool) {
		if !b {
			return
		}

		gitVars := utils.GitVars{
			GitDir:  gitLoc,
			GitUrl:  gitRepo.Text,
			GitUser: userName.Text,
		}
		utils.WriteGitConfigs(gitVars)
		win.SetContent(AppMain(win))
	}, win)

	form.Resize(fyne.NewSize(600, 600))

	form.Show()

}
