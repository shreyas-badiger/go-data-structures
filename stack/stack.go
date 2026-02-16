package stack

import (
	"fmt"
	"strconv"
	"strings"
)

// Stack is a LIFO stack of integers.
type Stack struct {
	items []int
}

// New returns an empty stack.
func New() *Stack {
	return &Stack{}
}

// Push adds n to the top of the stack.
func (s *Stack) Push(n int) {
	s.items = append(s.items, n)
}

// Pop removes and returns the top element. The second return is false if the stack was empty.
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	top := len(s.items) - 1
	val := s.items[top]
	s.items = s.items[:top]
	return val, true
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	return len(s.items)
}

// IsEmpty returns true if the stack has no elements.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Print() {
	fmt.Println(s.items)
}

func (s *Stack) PrettyPrint() {
	if len(s.items) == 0 {
		fmt.Println("+---+")
		fmt.Println("|   |  (empty)")
		fmt.Println("+---+")
		return
	}
	maxW := 1
	for _, n := range s.items {
		w := len(strconv.Itoa(n))
		if w > maxW {
			maxW = w
		}
	}
	if maxW < 2 {
		maxW = 2
	}
	top := "+" + strings.Repeat("-", maxW+2) + "+"
	fmt.Println(top)
	for i := len(s.items) - 1; i >= 0; i-- {
		fmt.Printf("| %*d |", maxW, s.items[i])
		if i == len(s.items)-1 {
			fmt.Print("  <- top")
		}
		fmt.Println()
	}
	fmt.Println(top)
}
