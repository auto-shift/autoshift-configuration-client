package fyne_objects

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/impls"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/structs"
)

var gitVars = structs.GitVars{}

// form for viewing and editing git settings
func GitSettings(win fyne.Window) fyne.CanvasObject {

	gitDir := impls.GitConfs.GetDir()
	dir := binding.BindString(&gitDir)
	gitDirHBox := container.New(
		layout.NewHBoxLayout(),
		widget.NewLabelWithStyle("Git Directory:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithData(dir),
	)

	gitUser := impls.GitConfs.GetUser()
	user := binding.BindString(&gitUser)
	gitUserHBox := container.New(
		layout.NewHBoxLayout(),
		widget.NewLabelWithStyle("Git Username:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithData(user),
	)

	gitUrl := impls.GitConfs.GetUrl()
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
						impls.GitClone(gitUser, gitPass.Text, gitDir, gitUrl)
						// gitResp := impls.GitClone(gitUser, gitPass.Text, gitDir, gitUrl)
						// dialog.ShowInformation(gitResp[0], gitResp[1], win)
					}, win)
			},
		),
		widget.NewButton(
			"Push to Remote Repo",
			func() {
				dialog.NewInformation("TODO:", "This button should push commits to remote repository", win)
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

	label := widget.NewLabel("Git Settings")
	label.Alignment = fyne.TextAlignCenter

	gitUser := impls.GitConfs.GetUser()
	user := binding.BindString(&gitUser)
	userName := widget.NewEntry()
	userName.Bind(user)

	gitUrl := impls.GitConfs.GetUrl()
	url := binding.BindString(&gitUrl)
	gitRepo := widget.NewEntry()
	gitRepo.Bind(url)

	getDirLoc := func(list fyne.ListableURI, err error) {
		if err != nil {
			fmt.Println("line 111, fyne_obj/git.go:")
			panic(err)
		}
		if list == nil {
			log.Println("Cancelled")
		}

		impls.GitConfs.SetDir(list.Path())
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

		impls.GitConfs.SetUrl(gitRepo.Text)

		impls.GitConfs.SetUser(userName.Text)
		fmt.Println(impls.GitConfs)
		gitVars.UpdateGitVars(impls.GitConfs)
		fmt.Println("updated gitVars:")
		fmt.Println(gitVars)
		gitVars.WriteGitConfigs()
		win.SetContent(AppMain(win))
	}, win)

	form.Resize(fyne.NewSize(600, 600))

	form.Show()

}
