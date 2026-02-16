package main

import "fmt"

type stack struct {
	intStack []int // No pointer needed
}

func (s *stack) push(num int) { // Use pointer receiver to modify the slice
	s.intStack = append(s.intStack, num)
}

func (s *stack) pop() int {
	if len(s.intStack) == 0 {
		return -1
	}
	val := s.intStack[len(s.intStack)-1]
	s.intStack = s.intStack[:len(s.intStack)-1]
	return val
}

func main() {
	s := stack{} // Automatically initializes to empty slice
	s.push(10)
	s.push(20)
	s.push(30)
	fmt.Println(s.pop()) // Output: 30
    fmt.Println(s.pop()) // Output: 20
    fmt.Println(s.pop()) // Output: 10
}