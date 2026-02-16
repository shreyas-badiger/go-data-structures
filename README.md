# go-data-structures

A small Go module with a few common data structures: graph (adjacency list), min-heap, queue, and stack. Use it as a dependency or copy the packages you need.

## Data structures

- **Graph** (`graph`) – Adjacency list graph. Supports directed and undirected edges, add/delete nodes and edges, get neighbors.
- **Heap** (`heap`) – Min-heap (array-based). Push and pop the minimum element.
- **Queue** (`queues`) – FIFO queue of nodes (struct with `Val` and `Neighbors`). Add, Remove, Front, IsEmpty, Size.
- **Stack** (`stack`) – LIFO stack of integers. Push, Pop, Size, IsEmpty.

## Usage

Add the module to your project:

```bash
go get github.com/shreyas-badiger/go-data-structures
```

Then import the packages you need (import path and package name are both lowercase):

```go
import (
    "github.com/shreyas-badiger/go-data-structures/graph"
    "github.com/shreyas-badiger/go-data-structures/heap"
    "github.com/shreyas-badiger/go-data-structures/queues"
    "github.com/shreyas-badiger/go-data-structures/stack"
)

func main() {
    // Stack
    s := stack.New()
    s.Push(1)
    s.Push(2)
    v, ok := s.Pop()  // 2, true

    // Min-heap
    h := heap.NewHeap()
    h.Push(5)
    h.Push(2)
    min := h.Pop()  // 2

    // Graph
    g := graph.NewGraph()
    g.AddUndirectedEdge(1, 2)
    g.AddDirectedEdge(1, 3)
    neighbors := g.GetNeighbors(1)

    // Queue (holds *queues.Node)
    q := queues.NewQueue()
    q.Add(&queues.Node{Val: 42})
    node := q.Remove()
}
```

## Tests

From the repo root:

```bash
go test ./...
```

## Module path

If you fork or rename the repo, update the module in `go.mod` (e.g. `module github.com/yourname/go-data-structures`) and use the same path in your imports.
