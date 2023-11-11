package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/data"
)

func cluster_interface(win fyne.Window) fyne.CanvasObject {
	cluster_card := widget.NewCard("Add a new cluster:", "",
		container.New(
			layout.NewHBoxLayout(),
			widget.NewLabel("List of clusters go here"),
			layout.NewSpacer(),
		),
	)

	new_cluster_dialog_hbox := container.NewHBox(
		layout.NewSpacer(),
		widget.NewButton("Add Cluster", func() { new_cluster_dialog(win) }),
	)

	cluster_box := container.New(
		layout.NewVBoxLayout(),
		cluster_card,
		new_cluster_dialog_hbox,
	)

	return cluster_box
}

func new_cluster_dialog(win fyne.Window) {
	var platform string
	new_cluster_name_entry := widget.NewEntry()
	new_cluster_domain_entry := widget.NewEntry()
	new_cluster_platform_select := widget.NewSelect([]string{"aws", "bare metal", "gcp", "azure"},
		func(sel string) {
			platform = sel
		})

	items := []*widget.FormItem{
		{Text: "Environment:", Widget: new_cluster_name_entry},
		{Text: "Domain:", Widget: new_cluster_domain_entry},
		{Text: "Service Platform:", Widget: new_cluster_platform_select},
	}

	cluster_form := dialog.NewForm(
		"Add new Cluster", "Submit", "Cancel", items,
		func(b bool) {
			if !b {
				return
			}

			data.ClusterList.AddNew(
				new_cluster_name_entry.Text,
				new_cluster_domain_entry.Text,
				platform,
			)

			newCluster := data.ClusterList.GetAll()[new_cluster_name_entry.Text]
			homeTabs.Append(&container.TabItem{Text: newCluster.GetName(), Icon: theme.StorageIcon(), Content: widget.NewLabel(newCluster.GetDomain())})

		},
		win)

	cluster_form.Show()
}
