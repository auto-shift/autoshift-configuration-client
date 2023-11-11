package structs

type cluster struct {
	name     string
	domain   string
	platform string
	nodes    nodes
	apps     apps
}

func (c cluster) GetName() string {
	return c.name
}

func (c cluster) GetDomain() string {
	return c.domain
}
func (c cluster) GetPlatform() string {
	return c.platform
}
