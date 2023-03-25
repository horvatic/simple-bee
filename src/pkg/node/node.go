package node

type node struct {
	ip   string
	name string
	tags []string
}

func NewNode(ip string, name string, tags []string) *node {
	n := node{ip: ip, name: name, tags: tags}
	return &n
}

func (n node) GetIp() string {
	return n.ip
}

func (n node) GetName() string {
	return n.name
}

func (n node) GetTags() []string {
	return n.tags
}
