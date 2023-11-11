package structs

import (
	"fmt"
	"sync"
)

var cluster_instance *clusters

var lock = &sync.Mutex{}

func Init_clusters() *clusters {

	if cluster_instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if cluster_instance == nil {
			fmt.Println("Creating new instance")
			cluster_instance = new(clusters)
			cluster_instance.all = make(map[string]cluster)
		} else {
			fmt.Println("Already Created 1")
		}
	} else {
		fmt.Println("Already Created 2")
	}
	return cluster_instance
}

type clusters struct {
	all map[string]cluster
}

func (cs *clusters) Add(c cluster) {
	cs.all[c.name] = c
}

func (cs *clusters) AddNew(c_name, c_domain, c_platform string) {
	cs.all[c_name] = cluster{c_name, c_domain, c_platform, nodes{}, apps{}}
}

func (cs *clusters) GetAll() map[string]cluster {
	return cs.all
}
