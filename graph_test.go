package graph

import (
	"testing"
	"fmt"
	)

type TestGraph struct {
	nodes []int
	edges []TestEdge
}

type TestEdge struct {
	start int
	end int
}

func TestGraphCreation(t *testing.T) {
	g := New()
	nVertices := g.NumberOfVertices()
	nEdges := g.NumberOfEdges()
	if nVertices != 0 {
		t.Errorf("Started with %d vertices", nVertices)
	}
	if nEdges != 0 {
		t.Errorf("Started with %d edges", nEdges)
	}
}


func TestAddSameEdges(t *testing.T) {
	nodesToAdd := []struct {
			label string
		}{
			{"1"},
			{"2"},
		}

	g := New()
	for _, node := range nodesToAdd {
		g.AddNode(node.label)
	}

	g.AddEdge(0,1)
	g.AddEdge(0,1)
	g.AddEdge(0,1)

	nEdges := g.NumberOfEdges()
	if nEdges != 1 {
		t.Errorf("Number of edges is one, as expected")
	}

}

func TestAddEdge(t *testing.T) {
	nodesToAdd := []struct {
			label string
		}{
			{"1"},
			{"2"},
			{"3"},
			{"4"},
			{"5"},
			{"6"},
		}
	g := New()
	for _, node := range nodesToAdd {
		g.AddNode(node.label)
	}
	for i , _ := range nodesToAdd {
		for j, _ := range nodesToAdd {
			g.AddEdge(i, j)
			j++
		} 
		i++
	}
	for i, _ := range nodesToAdd {
		for j, _ := range nodesToAdd {
			hasEdge, err := g.HasEdge(i,j)
			if err != nil {
				fmt.Println(err)
			}
			if hasEdge == false {
				t.Errorf("Graph does not have edge with start id %d and end id %d", i, j)
			}
		}
	}
	nEdges := g.NumberOfEdges()
	expectedNEdges := len(nodesToAdd)*len(nodesToAdd)
	if nEdges !=  expectedNEdges {
		t.Errorf("Graph does not have edge for every node - number of edges is %d, when %d is expected", nEdges, expectedNEdges)
	}

	for i, _ := range nodesToAdd {
		degree := g.Degree(i)
		if degree != len(nodesToAdd) {
			t.Errorf("Edge with id %d does not have expected degree %d", i, 6)
		}
	}
}

func TestAddAndGetNode(t *testing.T) {
	cases :=[]struct {
			label string
			expectedSize int
		}{
			{"1", 1},
			{"2", 2},
			{"3", 3},
			{"4", 4},
			{"5", 5},
			{"6", 6},
		}
		g := New()
		for _, c := range cases {
			g.AddNode(c.label)
			numberOfVertices := g.NumberOfVertices()
			if numberOfVertices != c.expectedSize {
				t.Errorf("Graph has %d vertices when %d were inserted", numberOfVertices, c.expectedSize)
			}
		}

		for i, c := range cases {
			node, err := g.GetNode(i)
			if err != nil {
				fmt.Println(err)
			}
			if node.Id != i && node.Label != c.label {
				t.Errorf("Graph has node with id %d and label %s at id %d", node.Id, node.Label, i)
			}
		}
}

