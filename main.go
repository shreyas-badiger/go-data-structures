package main

import (
	"fmt"

	"github.com/shreyas-badiger/go-data-structures/graph"
	"github.com/shreyas-badiger/go-data-structures/heap"
	"github.com/shreyas-badiger/go-data-structures/queue"
	"github.com/shreyas-badiger/go-data-structures/stack"
)

func main() {
	fmt.Println("\n\nStack:")
	s := stack.New()
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.PrettyPrint()

	fmt.Println("\n\nHeap:")
	h := heap.NewHeap()
	h.Push(5)
	h.Push(2)
	h.Push(8)
	h.PrettyPrint()

	fmt.Println("\n\nGraph:")
	g := graph.NewGraph()
	g.AddUndirectedEdge(1, 2)
	g.AddUndirectedEdge(1, 3)
	g.PrettyPrint()

	fmt.Println("\n\nQueue:")
	q := queue.NewQueue()
	q.Add(&queue.Node{Val: 1, Neighbors: nil})
	q.Add(&queue.Node{Val: 2, Neighbors: nil})
	q.PrettyPrint()
}
