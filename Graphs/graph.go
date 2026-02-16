package graphs

import "fmt"

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

	// remove node from all its neighbors
	for neighbor := range g.nodes[node] {
		delete(g.nodes[neighbor], neighbor)
	}

	delete(g.nodes, node)
}

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

func (g *Graph) Print() {
	for node, neighbors := range g.nodes {
		fmt.Printf("%d:", node)
		for neighbor := range neighbors {
			fmt.Printf("%d ", neighbor)
		}
		fmt.Println()
	}
}
