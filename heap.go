package main

import (
	"fmt"
	"math"
)

type Heap struct {
	data []int
}

func (h *Heap) root_node() int {
	return h.data[0]
}

func (h *Heap) last_node() int {
	return h.data[len(h.data)-1:][0]
}

func (h *Heap) left_child_index(index int) int {
	return (index * 2) + 1
}

func (h *Heap) right_child_index(index int) int {
	return (index * 2) + 2
}

func (h *Heap) parent_index(index int) int {
	return int(math.Trunc((float64(index) - 1) / 2))
}

func (h *Heap) insert(value int) {
	h.data = append(h.data, value)
	newNodeIndex := len(h.data) - 1

	for newNodeIndex > 0 &&
		(h.data[newNodeIndex] > h.data[(h.parent_index(newNodeIndex))]) {
		h.data[h.parent_index(newNodeIndex)], h.data[newNodeIndex] =
			h.data[newNodeIndex], h.data[h.parent_index(newNodeIndex)]
	}

	newNodeIndex = h.parent_index(newNodeIndex)
}

func (h *Heap) hasGreaterChild(trickleNodeIdx int) bool {
	if h.left_child_index(trickleNodeIdx) < len(h.data) &&
		h.data[h.left_child_index(trickleNodeIdx)] > h.data[trickleNodeIdx] {
		return true
	}
	if h.right_child_index(trickleNodeIdx) < len(h.data) &&
		h.data[h.right_child_index(trickleNodeIdx)] > h.data[trickleNodeIdx] {
		return true
	}
	return false
}

func (h *Heap) calculateLargerChildIdx(trickleNodeIdx int) (result int) {
	if !(h.right_child_index(trickleNodeIdx) < len(h.data)) {
		result = h.left_child_index(trickleNodeIdx)
	} else {
		if h.data[h.right_child_index(trickleNodeIdx)] > h.data[h.left_child_index(trickleNodeIdx)] {
			result = h.right_child_index(trickleNodeIdx)
		} else {
			result = h.left_child_index(trickleNodeIdx)
		}
	}
	return result
}

func (h *Heap) delete() {
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	trickleNodeIdx := 0

	for h.hasGreaterChild(trickleNodeIdx) {
		largerChildIdx := h.calculateLargerChildIdx(trickleNodeIdx)

		h.data[trickleNodeIdx], h.data[largerChildIdx] =
			h.data[largerChildIdx], h.data[trickleNodeIdx]

		trickleNodeIdx = largerChildIdx
	}
}

func main() {
	var heap1 Heap = Heap{}
	heap1.data = []int{100, 88, 25, 87, 16, 8, 12, 86, 50, 2, 15, 3}
	fmt.Println(heap1.data)
	// fmt.Println(heap1.root_node())
	// fmt.Println(heap1.last_node())
	// fmt.Println(heap1.right_child_index(2)) // 6
	// fmt.Println(heap1.left_child_index(2))  // 5
	// fmt.Println(heap1.parent_index(5))      // 2
	// heap1.insert(10)
	// fmt.Println(heap1.data)
	heap1.delete()
	fmt.Println(heap1.data)
}
