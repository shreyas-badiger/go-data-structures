package stack

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
