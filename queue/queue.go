package queue

import (
	"fmt"
	"strconv"
	"strings"
)

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

func (q *Queue) Print() {
	for _, node := range *q {
		fmt.Printf("%d ", node.Val)
	}
	fmt.Println()
}

func (q *Queue) PrettyPrint() {
	if len(*q) == 0 {
		fmt.Println("+---+")
		fmt.Println("|   |  (empty)")
		fmt.Println("+---+")
		fmt.Println("  ^")
		fmt.Println("front")
		return
	}
	maxW := 2
	for _, node := range *q {
		w := len(strconv.Itoa(node.Val))
		if w > maxW {
			maxW = w
		}
	}
	cellW := maxW + 2
	var sep string
	for range *q {
		sep += "+" + strings.Repeat("-", cellW)
	}
	sep += "+"
	fmt.Println(sep)
	inner := "|"
	for _, node := range *q {
		inner += fmt.Sprintf(" %*d |", maxW, node.Val)
	}
	fmt.Println(inner)
	fmt.Println(sep)
	fmt.Print("  ^")
	fmt.Print(strings.Repeat(" ", len(inner)-6))
	fmt.Println("^")
	fmt.Print("front")
	fmt.Print(strings.Repeat(" ", len(inner)-8))
	fmt.Println("back")
}
