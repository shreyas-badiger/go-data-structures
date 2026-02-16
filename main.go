package main

import (
	"fmt"

	"github.com/shreyas-badiger/go-data-structures/Graphs"
	"github.com/shreyas-badiger/go-data-structures/Heap"
	"github.com/shreyas-badiger/go-data-structures/Queues"
	"github.com/shreyas-badiger/go-data-structures/Stack"
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
	for v, ok := h.Pop(); ok; v, ok = h.Pop() {
		fmt.Println(v)
	}

	fmt.Println("Graph:")
	g := graphs.NewGraph()
	g.AddUndirectedEdge(1, 2)
	g.AddUndirectedEdge(1, 3)
	g.Print()

	fmt.Println("Queue:")
	q := queues.NewQueue()
	q.Add(&queues.Node{Val: 1, Neighbors: nil})
	q.Add(&queues.Node{Val: 2, Neighbors: nil})
	fmt.Println(q.Remove().Val, q.Remove().Val)
}
