package graph

//Very simple graph implementation without removal using adjacency list.
//Represnting nodes as IDs -> needs to be used in conjunction with mapping to show what node values are
//Node list index is id of each node (which is why removal is not implemented)

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

func (g *Graph) addNode() {
	newId := len(g.nodes) - 1
	g.nodes = append(g.nodes, newId)
}

func (g *Graph) hasNode(id int64) bool {
	nodes := g.nodes
	if (id < len(nodes)) {
		if (nodes[id] === nil) {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

func (g *Graph) hasEdge(start int64, end int64) {
	if (g.hasNode(start) && g.hasNode(end)) {
		nodes := g.nodes
		connectedNodes := nodes[start]
		_, ok := connectedNodes[end]
		return ok
	} else {
		return false
	}
}

func (g *Graph) addEdge(start int64, end int64) bool {
	if (g.hasNode(start) && g.hasNode(end)) {
		nodes := g.nodes
		connectedNodes := nodes[start] 
		connectedNodes[end] = true
		return true
	} else {
		return false
	}
}



