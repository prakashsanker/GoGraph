package graph

import (
	"testing"
	"fmt"
	)

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
			if node.id != i && node.label != c.label {
				t.Errorf("Graph has node with id %d and label %s at id %d", node.id, node.label, i)
			}
		}
}

