package fyne_objects

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/impls"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/utils"
)

var remove_kubeadmin = false

// sets the logger output file
func init() {
	log.SetOutput(utils.LogFile)
}

// main app content
func AppMain(win fyne.Window) fyne.CanvasObject {
	// logList := utils.GetLogs()
	// logs := binding.BindStringList(&(utils.GetLogs()))
	utils.MainWin = win

	logs := binding.NewStringList()
	logs.Set(utils.GetLogs())

	list := widget.NewListWithData(logs,
		func() fyne.CanvasObject {
			logLabel := widget.NewLabel("                                                                                                                            ")

			return logLabel
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	sep := widget.NewSeparator()

	right := container.New(layout.NewHBoxLayout(), sep, container.NewVScroll(list))
	right.Resize(fyne.Size{Height: 600})
	top := widget.NewLabel("")

	buildOpts := widget.NewRadioGroup([]string{"New", "Rebaseline"}, func(s string) {
		if s == "New" {
			log.Println("Creating a new cluster")
		}
		if s == "Rebaseline" {
			log.Println("Rebase existing cluster")
		}
	})

	buildOpts.Horizontal = true
	buildOpts.SetSelected("New")

	bottom := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		buildOpts,
		widget.NewLabel(" "),
		widget.NewButton("Build", func() {}),
	)

	middle := container.NewAppTabs(
		//TODO: Tab should allow configuration git credentials
		container.NewTabItem("General", generalSettings(win)),
		//TODO: Tab should allow configuration of cluster environment variables
		container.NewTabItem("Cluster", clusterSettings()),
		//TODO: Tab should allow create a top oriented AppTabs object with a tab for each service.
		container.NewTabItem("Apps", AppSettings(win)),
	)
	if !impls.IsLocalRepo() {
		middle.DisableIndex(1)
		middle.DisableIndex(2)
		middle.DisableIndex(3)
	}

	middle.SetTabLocation(container.TabLocationLeading)

	content := container.NewBorder(top, bottom, nil, right, middle)

	go func() {
		for range time.Tick(time.Second) {
			logs.Set(utils.GetLogs())
		}
	}()

	return content
}

// get/set installation variables

/// Menu functions

func MakeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("New", nil)
	checkedItem := fyne.NewMenuItem("Checked", nil)
	checkedItem.Checked = true
	disabledItem := fyne.NewMenuItem("Disabled", nil)
	disabledItem.Disabled = true
	otherItem := fyne.NewMenuItem("Other", nil)
	mailItem := fyne.NewMenuItem("Mail", func() { fmt.Println("Menu New->Other->Mail") })
	mailItem.Icon = theme.MailComposeIcon()
	otherItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Project", func() { fmt.Println("Menu New->Other->Project") }),
		mailItem,
	)
	fileItem := fyne.NewMenuItem("File", func() { fmt.Println("Menu New->File") })
	fileItem.Icon = theme.FileIcon()
	dirItem := fyne.NewMenuItem("Directory", func() { fmt.Println("Menu New->Directory") })
	dirItem.Icon = theme.FolderIcon()
	newItem.ChildMenu = fyne.NewMenu("",
		fileItem,
		dirItem,
		otherItem,
	)

	openSettings := func() {
		w := a.NewWindow("Fyne Settings")
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	}
	settingsItem := fyne.NewMenuItem("Settings", openSettings)
	settingsShortcut := &desktop.CustomShortcut{KeyName: fyne.KeyComma, Modifier: fyne.KeyModifierShortcutDefault}
	settingsItem.Shortcut = settingsShortcut
	w.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
		openSettings()
	})

	cutShortcut := &fyne.ShortcutCut{Clipboard: w.Clipboard()}
	cutItem := fyne.NewMenuItem("Cut", func() {
		shortcutFocused(cutShortcut, w)
	})
	cutItem.Shortcut = cutShortcut
	copyShortcut := &fyne.ShortcutCopy{Clipboard: w.Clipboard()}
	copyItem := fyne.NewMenuItem("Copy", func() {
		shortcutFocused(copyShortcut, w)
	})
	copyItem.Shortcut = copyShortcut
	pasteShortcut := &fyne.ShortcutPaste{Clipboard: w.Clipboard()}
	pasteItem := fyne.NewMenuItem("Paste", func() {
		shortcutFocused(pasteShortcut, w)
	})
	pasteItem.Shortcut = pasteShortcut
	performFind := func() { fmt.Println("Menu Find") }
	findItem := fyne.NewMenuItem("Find", performFind)
	findItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyF, Modifier: fyne.KeyModifierShortcutDefault | fyne.KeyModifierAlt | fyne.KeyModifierShift | fyne.KeyModifierControl | fyne.KeyModifierSuper}
	w.Canvas().AddShortcut(findItem.Shortcut, func(shortcut fyne.Shortcut) {
		performFind()
	})

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {
			u, _ := url.Parse("https://developer.fyne.io")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItem("Support", func() {
			u, _ := url.Parse("https://fyne.io/support/")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Sponsor", func() {
			u, _ := url.Parse("https://fyne.io/sponsor/")
			_ = a.OpenURL(u)
		}))

	// a quit item will be appended to our first (File) menu
	file := fyne.NewMenu("File", newItem, checkedItem, disabledItem)
	device := fyne.CurrentDevice()
	if !device.IsMobile() && !device.IsBrowser() {
		file.Items = append(file.Items, fyne.NewMenuItemSeparator(), settingsItem)
	}
	main := fyne.NewMainMenu(
		file,
		fyne.NewMenu("Edit", cutItem, copyItem, pasteItem, fyne.NewMenuItemSeparator(), findItem),
		helpMenu,
	)
	checkedItem.Action = func() {
		checkedItem.Checked = !checkedItem.Checked
		main.Refresh()
	}
	return main
}
func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	switch sh := s.(type) {
	case *fyne.ShortcutCopy:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutCut:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutPaste:
		sh.Clipboard = w.Clipboard()
	}

	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}
