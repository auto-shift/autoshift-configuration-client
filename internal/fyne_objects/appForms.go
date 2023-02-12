package fyne_objects

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/structs"
)

func init() {

}

var regions = []string{"us-east-1", "us-east-2", "us-west-1", "us-west-2"}

// func getItems(appName string) []*widget.FormItem {

// 	formItems := getArgoCDForm()

// 	return formItems
// }

func getArgoCDForm() []*widget.FormItem {
	argoVars := structs.ReadArgoVars()
	fmt.Println(argoVars)
	formItems := []*widget.FormItem{
		widget.NewFormItem("Region: ", widget.NewSelect(regions, func(s string) {})),
	}
	return formItems
}
