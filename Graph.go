package graph

import (
	"errors"
	// "fmt"
)

type Graph struct {
	Nodes []Node
}

type Node struct {
	Label string
	Id int
	Visited bool
	TimesVisited int
	Ids map[int]bool
}


func New()(*Graph) {
	g := new(Graph)
	g.Nodes = make([]Node, 0)
	return g
}

var count int


func (g *Graph) AddNode(label string) int {
	newId := len(g.Nodes)
	newConnectedIds := make(map[int]bool)
	newNode := Node{Ids: newConnectedIds, Label: label, Id: newId, Visited: false, TimesVisited: 0}
	g.Nodes = append(g.Nodes, newNode)
	count++
	return newId
}

func (g *Graph) GetNode(id int) (Node, error) {
	hasNode, err := g.HasNode(id)
	var node Node
	if err != nil {
		return node, err
	}

	if (hasNode) {
		node = g.Nodes[id]
	}
	return node, nil
}

func (g *Graph) GetNodes() []Node {
	return g.Nodes;
}

func (g *Graph) HasNode(id int) (bool, error) {
	Nodes := g.Nodes
	if id < len(Nodes) {
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
		Nodes := g.Nodes
		connectedNodes := Nodes[start].Ids
		_, ok := connectedNodes[end]
		return ok, nil
	}
	return false, nil
}

func (g* Graph) GetConnectedVertices(id int) (map[int]bool, error) {
	var connectedNodes map[int]bool
	node, err := g.GetNode(id)
	if err != nil {
		return connectedNodes, err
	}
	return node.Ids, nil
}

func (g *Graph) NumberOfVertices() int{
	return len(g.Nodes)
}

func (g *Graph) NumberOfEdges() int{
	edgeCounter := 0
	for _, node := range g.Nodes {
		connectedNodes := node.Ids
		edgeCounter += len(connectedNodes)
	}
	return edgeCounter
}

func (g *Graph) Degree(id int) int {
	node := g.Nodes[id]
	return len(node.Ids)
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
		Nodes := g.Nodes
		connectedNodes := Nodes[start].Ids
		connectedNodes[end] = true
		return true, nil
	}
	return false, nil
}

var visited map[int]bool
var recursiveStack map[int]bool
var sortedOrder map[int]Node


func (g *Graph)hasCycle(id int, visited map[int]bool, recursiveStack map[int]bool) bool {
	node, _ := g.GetNode(id)
	if visited[node.Id] == false {
		visited[node.Id] = true
		recursiveStack[node.Id] = true
		connected := node.Ids
		g.Nodes[id] = node
		for connectedNodeId, _ := range connected {
			if visited[connectedNodeId] == false && g.hasCycle(connectedNodeId, visited, recursiveStack) {
				return true
			} else if recursiveStack[connectedNodeId] == true {
				return true
			}
		}
	}
	recursiveStack[node.Id] = false
	return false
}

func (g *Graph) HasCycle() bool {
	for _, node := range g.Nodes {
		visited = make(map[int]bool)
		recursiveStack = make(map[int]bool)
		if (g.hasCycle(node.Id, visited, recursiveStack)) {
			return true
		}
	}
	return false
}

func (g *Graph) TopologicalSort() map[int]Node {
	visited = make(map[int]bool)
	sortedOrder = make(map[int]Node)
	for _, vertex := range g.Nodes {
		if visited[vertex.Id] != true {
			g.topologicalSort(vertex.Id, sortedOrder)
		}
	}
	return sortedOrder
}

func (g *Graph) topologicalSort(id int, sortedOrder map[int]Node) {
	node, _ := g.GetNode(id)
	visited[node.Id] = true
	connected := node.Ids
	g.Nodes[id] = node
	for connectedNodeId, _ := range connected {
		connectedNode, _ := g.GetNode(connectedNodeId)
		if visited[connectedNode.Id] == false {
			g.topologicalSort(connectedNodeId, sortedOrder)
		}
	}
	sortedOrder[count] = node
	count--
}

