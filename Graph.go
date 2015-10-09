package graph

type Graph struct {
	nodes []Node
}

type Node struct {
	ids map[int64]bool
}

func New()(*Graph) {
	g := new Graph
	g.nodes = make([]Node)
	return g
}

func (g *Graph) addEdge(start int64, end int64) {
	nodes := g.nodes
	connectedNodes := nodes[start] 
	connectedNodes[end] = true
}

func (g *Graph) removeEdge(start int64, end int64) {
	nodes := g.nodes
	connectedNodes := nodes[start]
	connectedNodes[end] = false
}

