package data

import (
	"fmt"

	"github.com/auto-shift/autoshift-configuration-client/cmd/acc/internal/structs"
)

func init() {
	seedClusters()
	fmt.Println(ClusterList)
}

var ClusterList = structs.Init_clusters()

func seedClusters() {
	fmt.Println("Adding seed data")
	ClusterList.AddNew("dev", "dev.ocp4.example.com:6443", "aws")
	ClusterList.AddNew("preprod", "preprod.ocp4.example.com:6443", "aws")
	ClusterList.AddNew("prod", "prod.ocp4.example.com:6443", "aws")
}
