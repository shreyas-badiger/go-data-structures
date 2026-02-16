package stack

import (
	"testing"
)

func TestStackPushPop(t *testing.T) {
	s := New()
	s.Push(10)
	s.Push(20)
	s.Push(30)
	if s.Size() != 3 {
		t.Errorf("Size: got %d, want 3", s.Size())
	}
	v, ok := s.Pop()
	if !ok || v != 30 {
		t.Errorf("Pop: got %d, %v; want 30, true", v, ok)
	}
	v, ok = s.Pop()
	if !ok || v != 20 {
		t.Errorf("Pop: got %d, %v; want 20, true", v, ok)
	}
	v, ok = s.Pop()
	if !ok || v != 10 {
		t.Errorf("Pop: got %d, %v; want 10, true", v, ok)
	}
	v, ok = s.Pop()
	if ok {
		t.Errorf("Pop on empty stack: got ok=true, want false")
	}
}

func TestStackEmpty(t *testing.T) {
	s := New()
	if !s.IsEmpty() {
		t.Error("new stack should be empty")
	}
	_, ok := s.Pop()
	if ok {
		t.Error("Pop on empty stack should return ok=false")
	}
}

func TestStackLIFO(t *testing.T) {
	s := New()
	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	for i := 4; i >= 0; i-- {
		v, ok := s.Pop()
		if !ok || v != i {
			t.Errorf("Pop: got %d, %v; want %d, true", v, ok, i)
		}
	}
}
