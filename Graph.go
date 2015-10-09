package graph

import (
	"errors"
)

//Very simple graph implementation without removal using adjacency list.
//Represnting nodes as IDs -> needs to be used in conjunction with mapping to show what node values are
//Node list index is id of each node (which is why removal is not implemented)

type Graph struct {
	nodes []Node
}

type Node struct {
	ids map[int]bool
}

func New()(*Graph) {
	g := new(Graph)
	g.nodes = make([]Node, 5)
	return g
}

func (g *Graph) AddNode() int {
	newId := len(g.nodes) - 1
	newConnectedIds := make(map[int]bool)
	newNode := Node{ids: newConnectedIds}
	g.nodes = append(g.nodes, newNode)
	return newId
}

func (g *Graph) HasNode(id int) (bool, error) {
	nodes := g.nodes
	if id < len(nodes) {
		//kosher because I'm not implementing removal
		return true, nil
	} else {
		return false, errors.New("Node id not in graph")
	}
}

//first few lines could be composed out to a separate function
func (g *Graph) HasEdge(start int, end int) (bool, error) {
	hasStartNode, err := g.HasNode(start)
	if err != nil {
		return false, errors.New("Illegal start node id")
	}
	hasEndNode, err := g.HasNode(end) 
	if err != nil {
		return false, errors.New("Illegal end node id")
	}
	if hasStartNode && hasEndNode {
		nodes := g.nodes
		connectedNodes := nodes[start].ids
		_, ok := connectedNodes[end]
		return ok, nil
	}
	return false, nil
}

func (g *Graph) NumberOfVertices() int{
	return len(g.nodes)
}

func (g *Graph) NumberOfEdges() int{
	edgeCounter := 0
	for _, node := range g.nodes {
		connectedNodes := node.ids
		edgeCounter += len(connectedNodes)
	}
	return edgeCounter
}

func (g *Graph) Degree(id int) int {
	node := g.nodes[id]
	return len(node.ids)
}

//first few lines could be composed out to a separate function
func (g *Graph) AddEdge(start int, end int) (bool, error) {
	hasStartNode, err := g.HasNode(start)
	if err != nil {
		return false, errors.New("Illegal start node id")
	}
	hasEndNode, err := g.HasNode(end) 
	if err != nil{
		return false, errors.New("Illegal end node id")
	}
	if hasEndNode && hasStartNode {
		nodes := g.nodes
		connectedNodes := nodes[start].ids
		connectedNodes[end] = true
		return true, nil
	}
	return false, nil
}



