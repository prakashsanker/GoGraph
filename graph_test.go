package graph

import "testing"

func TestAddNode(t *testing.T) {

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
}