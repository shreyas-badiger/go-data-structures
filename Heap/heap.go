package heap

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
