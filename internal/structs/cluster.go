package structs

// EnvNode
type EnvNode struct {
	instanceTypes []string
	instanceSizes map[string][]string
}

// Methods
func (en EnvNode) GetInstanceTypes() []string {
	return en.instanceTypes
}

func (en EnvNode) GetInstanceSizes() map[string][]string {
	return en.instanceSizes
}

func NewEnvNode(it []string, is map[string][]string) EnvNode {
	nen := EnvNode{it, is}
	return nen
}

// NodeConfs
type NodeConfs struct {
	envList map[string]EnvNode
}

func (confs NodeConfs) AddEnv(env string, envConfs EnvNode) {
	confs.envList[env] = envConfs
}

func (confs NodeConfs) GetEnvConfigs(env string) EnvNode {
	return confs.envList[env]
}

func NewNodeConfs(el map[string]EnvNode) NodeConfs {
	nc := NodeConfs{el}
	return nc
}

//Node structs

type NodeInfo struct {
	nodeType string
	nodeSize string
	replicas int
}

func (node NodeInfo) GetNodeType() string {
	return node.nodeType
}
func (node NodeInfo) GetNodeNodeSize() string {
	return node.nodeSize
}
func (node NodeInfo) GetReplicas() int {
	return node.replicas
}

func NewNode(nt string, ns string, n int) NodeInfo {
	node := NodeInfo{nt, ns, n}
	return node
}

type AllNodes struct {
	master  NodeInfo
	infra   NodeInfo
	worker  NodeInfo
	storage NodeInfo
}

func (nodes AllNodes) GetMaster() NodeInfo {
	return nodes.master
}
func (nodes AllNodes) GetInfra() NodeInfo {
	return nodes.infra
}
func (nodes AllNodes) GetWorker() NodeInfo {
	return nodes.worker
}
func (nodes AllNodes) GetStorage() NodeInfo {
	return nodes.storage
}

// tab structs
type TabVars struct {
	enabled     bool
	env         string
	clusterName string
	nodes       AllNodes
}

func (vars TabVars) GetEnv() string {
	return vars.env
}

func (vars TabVars) SetEnv(env string) {
	vars.env = env
}

func (vars TabVars) GetNodes() AllNodes {
	return vars.nodes
}

func (vars TabVars) Set(enb bool, env string, cName string, nds AllNodes) {
	vars.enabled = enb
	vars.env = env
	vars.clusterName = cName
	vars.nodes = nds
}

type TabList struct {
	envs []TabVars
}

func (tabs TabList) AddTab(tv TabVars) {
	tabs.envs = append(tabs.envs, tv)
}

func (tabs TabList) GetTabs() []TabVars {
	return tabs.envs
}

func (tabs TabList) SearchTabs(env string) TabVars {
	tab := TabVars{}

	for _, t := range tabs.GetTabs() {
		if tab.env == env {
			tab = t
		}
	}

	return tab

}
