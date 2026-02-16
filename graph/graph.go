package graph

import (
	"fmt"
	"sort"
)

type Graph struct {
	nodes map[int]map[int]bool
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]map[int]bool),
	}
}

func (g *Graph) NewNode(val int) {
	if _, exists := g.nodes[val]; !exists {
		g.nodes[val] = make(map[int]bool)
	}
}

func (g *Graph) AddDirectedEdge(node1 int, node2 int) {
	if _, exists := g.nodes[node1]; !exists {
		g.NewNode(node1)
	}
	if _, exists := g.nodes[node2]; !exists {
		g.NewNode(node2)
	}
	g.nodes[node1][node2] = true
}

func (g *Graph) AddUndirectedEdge(node1 int, node2 int) {
	if _, exists := g.nodes[node1]; !exists {
		g.NewNode(node1)
	}
	if _, exists := g.nodes[node2]; !exists {
		g.NewNode(node2)
	}
	g.nodes[node1][node2] = true
	g.nodes[node2][node1] = true
}

func (g *Graph) GetNeighbors(node int) []int {
	neighborsList := []int{}
	for neighbor := range g.nodes[node] {
		neighborsList = append(neighborsList, neighbor)
	}
	return neighborsList
}

func (g *Graph) DeleteNode(node int) {
	if _, exists := g.nodes[node]; !exists {
		return
	}

	// remove this node from all its neighbors' adjacency lists
	for neighbor := range g.nodes[node] {
		delete(g.nodes[neighbor], node)
	}

	delete(g.nodes, node)
}

// DeleteEdge removes the edge between node1 and node2 (both directions for undirected).
func (g *Graph) DeleteEdge(node1, node2 int) {
	if _, exists := g.nodes[node1]; !exists {
		return
	}
	if _, exists := g.nodes[node2]; !exists {
		return
	}
	delete(g.nodes[node1], node2)
	delete(g.nodes[node2], node1)
}

// DeleteDirectedEdge removes only the edge from node1 to node2.
func (g *Graph) DeleteDirectedEdge(node1, node2 int) {
	if _, exists := g.nodes[node1]; !exists {
		return
	}
	delete(g.nodes[node1], node2)
}

func (g *Graph) Print() {
	for node, neighbors := range g.nodes {
		fmt.Printf("%d:", node)
		for neighbor := range neighbors {
			fmt.Printf("%d ", neighbor)
		}
		fmt.Println()
	}
}

func (g *Graph) PrettyPrint() {
	if len(g.nodes) == 0 {
		fmt.Println("(empty graph)")
		return
	}
	nodes := make([]int, 0, len(g.nodes))
	for n := range g.nodes {
		nodes = append(nodes, n)
	}
	sort.Ints(nodes)
	fmt.Println("Nodes:", nodes)
	seenUndirected := make(map[string]bool)
	fmt.Print("Edges: ")
	first := true
	for _, from := range nodes {
		for to := range g.nodes[from] {
			if from == to {
				continue
			}
			back := g.nodes[to][from]
			if back {
				key := fmt.Sprintf("%d-%d", minInt(from, to), maxInt(from, to))
				if seenUndirected[key] {
					continue
				}
				seenUndirected[key] = true
				if !first {
					fmt.Print(", ")
				}
				fmt.Printf("%d -- %d", minInt(from, to), maxInt(from, to))
			} else {
				if !first {
					fmt.Print(", ")
				}
				fmt.Printf("%d -> %d", from, to)
			}
			first = false
		}
	}
	if first {
		fmt.Print("(none)")
	}
	fmt.Println()
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
