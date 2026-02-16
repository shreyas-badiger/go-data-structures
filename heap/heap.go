package heap

import (
	"fmt"
	"strings"
)

type Heap struct {
	items []int
}

func NewHeap() *Heap {
	return &Heap{}
}

func (h *Heap) Push(value int) {
	h.items = append(h.items, value)
	h.bubbleUp(len(h.items) - 1)
}

// Pop removes and returns the minimum element.
func (h *Heap) Pop() int {
	if len(h.items) == 0 {
		return 0
	}
	rootIndex, lastIndex := 0, len(h.items)-1
	rootNode, lastNode := h.items[rootIndex], h.items[lastIndex]

	h.items[rootIndex] = lastNode
	h.items = h.items[:lastIndex]

	if len(h.items) > 0 {
		h.bubbleDown(rootIndex)
	}
	return rootNode
}

func (h *Heap) Size() int {
	return len(h.items)
}

func (h *Heap) Print() {
	fmt.Println(h.items)
}

func (h *Heap) PrettyPrint() {
	if len(h.items) == 0 {
		fmt.Println("(empty heap)")
		return
	}
	hght := heapHeight(len(h.items))
	cellW := 5
	width := cellW * (1 << hght)
	for level := 0; level < hght; level++ {
		start := (1 << level) - 1
		count := 1 << level
		if start >= len(h.items) {
			break
		}
		if start+count > len(h.items) {
			count = len(h.items) - start
		}
		slotW := width >> level
		var line string
		for c := 0; c < count && start+c < len(h.items); c++ {
			s := fmt.Sprintf("[%d]", h.items[start+c])
			pad := slotW - len(s)
			line += strings.Repeat(" ", pad/2) + s + strings.Repeat(" ", (pad+1)/2)
		}
		fmt.Println(strings.TrimRight(line, " "))
	}
}

func heapHeight(n int) int {
	if n <= 0 {
		return 0
	}
	h := 0
	for (1 << h) <= n {
		h++
	}
	return h
}

// Internal: helper to move up the node to its right place
func (h *Heap) bubbleUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.items[index] < h.items[parent] {
			h.swap(index, parent)
			index = parent
		} else {
			break
		}
	}
}

// Internal: helper to move down the node to its right place
func (h *Heap) bubbleDown(index int) {
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < len(h.items) && h.items[left] < h.items[smallest] {
			smallest = left
		}
		if right < len(h.items) && h.items[right] < h.items[smallest] {
			smallest = right
		}

		if smallest != index {
			h.swap(index, smallest)
			index = smallest
		} else {
			break
		}
	}

}

// Internal: helper to swap elements
func (h *Heap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

/*
Understanding the relationship between array elements and parents child relationships.
					2
			3				4
		6		7		9		11

  0,		1,		2,		3,   	4,  	5,  	6,  	7
  2,		3,		4,		6,		7,		9,		11

parent - i
leftChild - 2i + 1
rightChild - 2i + 2

for a given i, parent - (i - 1) / 2

*/