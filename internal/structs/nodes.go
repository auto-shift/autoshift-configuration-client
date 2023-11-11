package structs

type nodes struct {
	nodeList []node
}

type node struct {
	node_type string
	replicas  int
}
