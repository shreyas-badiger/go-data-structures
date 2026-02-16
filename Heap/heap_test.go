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
	v, ok := h.Pop()
	if !ok || v != 1 {
		t.Errorf("Pop: got %d, %v; want 1, true", v, ok)
	}
	v, ok = h.Pop()
	if !ok || v != 2 {
		t.Errorf("Pop: got %d, %v; want 2, true", v, ok)
	}
	v, ok = h.Pop()
	if !ok || v != 5 {
		t.Errorf("Pop: got %d, %v; want 5, true", v, ok)
	}
	v, ok = h.Pop()
	if !ok || v != 8 {
		t.Errorf("Pop: got %d, %v; want 8, true", v, ok)
	}
	v, ok = h.Pop()
	if ok {
		t.Errorf("Pop on empty heap: got ok=true, want false")
	}
}

func TestHeapEmpty(t *testing.T) {
	h := NewHeap()
	_, ok := h.Pop()
	if ok {
		t.Error("Pop on empty heap should return ok=false")
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
		v, ok := h.Pop()
		if !ok {
			t.Fatal("unexpected ok=false")
		}
		if !first && v < prev {
			t.Errorf("min-heap order violated: %d after %d", v, prev)
		}
		first = false
		prev = v
	}
}
