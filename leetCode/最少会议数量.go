package leetCode

import (
	"container/heap"
	"sort"
)

func MinArragedRooms(job [][]int) int {
	sort.Slice(job, func(i, j int) bool {
		return job[i][0] < job[j][0]
	})

	minHeap := new(MinEndHeap)
	heap.Init(minHeap)
	heap.Push(minHeap, [2]int{job[0][0], job[0][1]})
	for i := 1; i < len(job); i++ {
		if (*minHeap)[0][1] <= job[i][0] {
			heap.Pop(minHeap)
		}
		heap.Push(minHeap, [2]int{job[i][0], job[i][1]})
	}
	return minHeap.Len()
}

type MinEndHeap [][2]int

func (h MinEndHeap) Len() int {
	return len(h)
}
func (h MinEndHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h MinEndHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinEndHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *MinEndHeap) Pop() any {
	tmp := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return tmp
}
