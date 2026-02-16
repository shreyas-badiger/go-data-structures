package graphs

import (
	"sort"
	"testing"
)

func sorted(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)
	sort.Ints(c)
	return c
}

func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sa, sb := sorted(a), sorted(b)
	for i := range sa {
		if sa[i] != sb[i] {
			return false
		}
	}
	return true
}

func TestNewGraph(t *testing.T) {
	g := NewGraph()
	if g == nil {
		t.Fatal("NewGraph() returned nil")
	}
	if n := g.GetNeighbors(1); len(n) != 0 {
		t.Errorf("new graph should have no neighbors for 1, got %v", n)
	}
}

func TestAddUndirectedEdge(t *testing.T) {
	g := NewGraph()
	g.AddUndirectedEdge(1, 2)
	g.AddUndirectedEdge(1, 3)
	if got := g.GetNeighbors(1); !eq(got, []int{2, 3}) {
		t.Errorf("neighbors of 1: got %v, want [2 3]", got)
	}
	if got := g.GetNeighbors(2); !eq(got, []int{1}) {
		t.Errorf("neighbors of 2: got %v, want [1]", got)
	}
}

func TestAddDirectedEdge(t *testing.T) {
	g := NewGraph()
	g.AddDirectedEdge(1, 2)
	g.AddDirectedEdge(1, 3)
	if got := g.GetNeighbors(1); !eq(got, []int{2, 3}) {
		t.Errorf("neighbors of 1: got %v, want [2 3]", got)
	}
	if got := g.GetNeighbors(2); len(got) != 0 {
		t.Errorf("directed: neighbors of 2 should be empty, got %v", got)
	}
}

func TestDeleteNode(t *testing.T) {
	g := NewGraph()
	g.AddUndirectedEdge(1, 2)
	g.AddUndirectedEdge(1, 3)
	g.AddUndirectedEdge(2, 3)
	g.DeleteNode(1)
	if got := g.GetNeighbors(2); !eq(got, []int{3}) {
		t.Errorf("after deleting 1, neighbors of 2: got %v, want [3]", got)
	}
	if got := g.GetNeighbors(3); !eq(got, []int{2}) {
		t.Errorf("after deleting 1, neighbors of 3: got %v, want [2]", got)
	}
}

func TestDeleteEdge(t *testing.T) {
	g := NewGraph()
	g.AddUndirectedEdge(1, 2)
	g.AddUndirectedEdge(1, 3)
	g.DeleteEdge(1, 2)
	if got := g.GetNeighbors(1); !eq(got, []int{3}) {
		t.Errorf("after DeleteEdge(1,2), neighbors of 1: got %v, want [3]", got)
	}
	if got := g.GetNeighbors(2); len(got) != 0 {
		t.Errorf("after DeleteEdge(1,2), neighbors of 2: got %v, want []", got)
	}
}

func TestDeleteDirectedEdge(t *testing.T) {
	g := NewGraph()
	g.AddDirectedEdge(1, 2)
	g.AddDirectedEdge(2, 1)
	g.DeleteDirectedEdge(1, 2)
	if got := g.GetNeighbors(1); len(got) != 0 {
		t.Errorf("after DeleteDirectedEdge(1,2), neighbors of 1: got %v", got)
	}
	if got := g.GetNeighbors(2); !eq(got, []int{1}) {
		t.Errorf("after DeleteDirectedEdge(1,2), neighbors of 2: got %v, want [1]", got)
	}
}
