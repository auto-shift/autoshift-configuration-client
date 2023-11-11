package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/data"
)

var homeTabs container.AppTabs

func AppMain(win fyne.Window) fyne.CanvasObject {
	for k, v := range data.ClusterList.GetAll() {
		tab := &container.TabItem{Text: k, Icon: theme.StorageIcon(), Content: widget.NewLabel(v.GetDomain())}
		homeTabs.Append(tab)
	}
	top := container.New(
		layout.NewVBoxLayout(),
		cluster_interface(win))
	homeTabs.SetTabLocation(container.TabLocationLeading)

	content := container.NewBorder(top, nil, nil, nil, &homeTabs)

	return content
}
