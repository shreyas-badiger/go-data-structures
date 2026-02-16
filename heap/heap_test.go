package heap

import (
	"testing"
)

func TestHeapPushPop(t *testing.T) {
	h := NewHeap()
	h.Push(5)
	h.Push(2)
	h.Push(8)
	h.Push(1)
	if h.Size() != 4 {
		t.Errorf("Size: got %d, want 4", h.Size())
	}
	v := h.Pop()
	if v != 1 {
		t.Errorf("Pop: got %d, want 1", v)
	}
	v = h.Pop()
	if v != 2 {
		t.Errorf("Pop: got %d, want 2", v)
	}
	v = h.Pop()
	if v != 5 {
		t.Errorf("Pop: got %d, want 5", v)
	}
	v = h.Pop()
	if v != 8 {
		t.Errorf("Pop: got %d, want 8", v)
	}
	_ = h.Pop() // empty heap returns 0
}

func TestHeapEmpty(t *testing.T) {
	h := NewHeap()
	v := h.Pop()
	if v != 0 {
		t.Errorf("Pop on empty heap: got %d, want 0", v)
	}
	if h.Size() != 0 {
		t.Errorf("empty heap Size: got %d", h.Size())
	}
}

func TestHeapOrder(t *testing.T) {
	h := NewHeap()
	for _, v := range []int{3, 1, 4, 1, 5, 9, 2, 6} {
		h.Push(v)
	}
	var prev int
	first := true
	for h.Size() > 0 {
		v := h.Pop()
		if !first && v < prev {
			t.Errorf("min-heap order violated: %d after %d", v, prev)
		}
		first = false
		prev = v
	}
}
