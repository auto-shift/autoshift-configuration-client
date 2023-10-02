package data

import (
	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/structs"
)

func init() {
	Confs = structs.NewNodeConfs(envConfs)
}

// creates a public instance of the NodeConf struct
var Confs structs.NodeConfs

var envConfs = map[string]structs.EnvNode{
	"AWS": structs.NewEnvNode(
		[]string{"i3", "m4", "m5", "m5a", "m6i", "c4", "c5", "c5a", "r4", "r5", "r5a", "t3", "t3a"},
		map[string][]string{
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
		},
	),
	"GCP": structs.NewEnvNode(
		[]string{"i3", "m4", "m5", "m5a", "m6i", "c4", "c5", "c5a", "r4", "r5", "r5a", "t3", "t3a"},
		map[string][]string{
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
		},
	),
	"Bare Metal": structs.NewEnvNode(
		[]string{"i3", "m4", "m5", "m5a", "m6i", "c4", "c5", "c5a", "r4", "r5", "r5a", "t3", "t3a"},
		map[string][]string{
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
		},
	),
}
