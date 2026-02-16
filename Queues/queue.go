package queues

type Node struct {
	Val       int
	Neighbors []*Node
}

type Queue []*Node

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Add(node *Node) {
	*q = append(*q, node)
}

func (q *Queue) Remove() *Node {
	if q.IsEmpty() {
		return nil
	}

	node := (*q)[0]
	*q = (*q)[1:]

	return node
}

func (q *Queue) Front() *Node {
	if q.IsEmpty() {
		return nil
	}
	return (*q)[0]
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Size() int {
	return len(*q)
}
