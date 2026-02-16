package queues

import (
	"testing"
)

func TestQueueAddRemove(t *testing.T) {
	q := NewQueue()
	if !q.IsEmpty() {
		t.Error("new queue should be empty")
	}
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	q.Add(n1)
	q.Add(n2)
	if q.Size() != 2 {
		t.Errorf("Size: got %d, want 2", q.Size())
	}
	if q.Front() != n1 {
		t.Error("Front should be n1")
	}
	got := q.Remove()
	if got != n1 || got.Val != 1 {
		t.Errorf("Remove: got %v, want n1 (Val=1)", got)
	}
	got = q.Remove()
	if got != n2 || got.Val != 2 {
		t.Errorf("Remove: got %v, want n2 (Val=2)", got)
	}
	if !q.IsEmpty() {
		t.Error("queue should be empty after two removes")
	}
}

func TestQueueRemoveEmpty(t *testing.T) {
	q := NewQueue()
	if q.Remove() != nil {
		t.Error("Remove on empty queue should return nil")
	}
	if q.Front() != nil {
		t.Error("Front on empty queue should return nil")
	}
}
