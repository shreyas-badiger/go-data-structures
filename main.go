package main

import (
	"fmt"

	"github.com/shreyas-badiger/go-data-structures/graph"
	"github.com/shreyas-badiger/go-data-structures/heap"
	"github.com/shreyas-badiger/go-data-structures/queues"
	"github.com/shreyas-badiger/go-data-structures/stack"
)

func main() {
	fmt.Println("Stack:")
	s := stack.New()
	s.Push(10)
	s.Push(20)
	s.Push(30)
	for v, ok := s.Pop(); ok; v, ok = s.Pop() {
		fmt.Println(v)
	}

	fmt.Println("Heap:")
	h := heap.NewHeap()
	h.Push(5)
	h.Push(2)
	h.Push(8)
	for h.Size() > 0 {
		fmt.Println(h.Pop())
	}

	fmt.Println("Graph:")
	g := graph.NewGraph()
	g.AddUndirectedEdge(1, 2)
	g.AddUndirectedEdge(1, 3)
	g.Print()

	fmt.Println("Queue:")
	q := queues.NewQueue()
	q.Add(&queues.Node{Val: 1, Neighbors: nil})
	q.Add(&queues.Node{Val: 2, Neighbors: nil})
	fmt.Println(q.Remove().Val, q.Remove().Val)
}
