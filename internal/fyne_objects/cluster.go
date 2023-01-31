package fyne_objects

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var instanceTypes = []string{"i3", "m4", "m5", "m5a", "m6i", "c4", "c5", "c5a", "r4", "r5", "r5a", "t3", "t3a"}
var instanceSize = map[string][]string{
	"i3":  {"large"},
	"m4":  {"large", "xlarge", "2xlarge", "4xlarge", "10xlarge", "16xlarge"},
	"m5":  {"large", "xlarge", "2xlarge", "4xlarge", "8xlarge", "12xlarge", "16xlarge"},
	"m5a": {"large", "xlarge", "2xlarge", "4xlarge", "8xlarge", "10xlarge", "16xlarge"},
	"m6i": {"xlarge", "2xlarge", "4xlarge", "8xlarge", "16xlarge"},
	"c4":  {"2xlarge", "4xlarge", "8xlarge"},
	"c5":  {"xlarge", "2xlarge", "4xlarge", "9xlarge", "12xlarge", "18xlarge", "24xlarge"},
	"c5a": {"xlarge", "2xlarge", "4xlarge", "8xlarge", "12xlarge", "16xlarge", "24xlarge"},
	"r4":  {"large", "xlarge", "2xlarge", "4xlarge", "8xlarge", "16xlarge"},
	"r5":  {"large", "xlarge", "2xlarge", "4xlarge", "8xlarge", "12xlarge", "16xlarge", "24xlarge"},
	"r5a": {"large", "xlarge", "2xlarge", "4xlarge", "8xlarge", "12xlarge", "16xlarge", "24xlarge"},
	"t3":  {"large", "xlarge", "2xlarge"},
	"t3a": {"large", "xlarge", "2xlarge"},
}

func clusterSettings() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItem("Cluster", widget.NewLabel("See TODO ")),
	)

	return tabs
}
