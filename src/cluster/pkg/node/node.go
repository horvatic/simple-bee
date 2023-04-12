package node

type Node struct {
	ip   string
	name string
	tags []string
}

func NewNode(ip string, name string, tags []string) *Node {
	n := Node{ip: ip, name: name, tags: tags}
	return &n
}

func (n Node) GetIp() string {
	return n.ip
}

func (n Node) GetName() string {
	return n.name
}

func (n Node) GetTags() []string {
	return n.tags
}
