package leetCode

import (
	"fmt"
	"slices"
)

//import "container/heap"
//
//func HeapSort(nums []int) {
//	heapA := new(HeapArr)
//	for i := 0; i < len(nums); i++ {
//		heap.Push(heapA, nums[i])
//	}
//	heap.Init(heapA)
//	for i := len(nums) - 1; i >= 0; i-- {
//		nums[i] = heap.Pop(heapA).(int)
//	}
//
//}
//
//type HeapArr []int
//
//func (h HeapArr) Len() int           { return len(h) }
//func (h HeapArr) Less(i, j int) bool { return h[i] > h[j] }
//func (h HeapArr) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
//func (h *HeapArr) Push(x interface{}) {
//	*h = append(*h, x.(int))
//}
//
//func (h *HeapArr) Pop() interface{} {
//	old := (*h)[len(*h)-1]
//	*h = (*h)[:len(*h)-1]
//	return old
//}

func HeapSort(nums []int) {
	HeapAdjust(nums)
	fmt.Printf("%#v\n", nums)
	for i := 0; i < len(nums)-1; i++ {
		nums[0], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[0]
		Down(nums, 0, len(nums)-i-1)
	}
	slices.Reverse(nums)
	fmt.Printf("%#v", nums)
}

func HeapAdjust(nums []int) {
	for i := len(nums) - 2; i >= 0; i-- {
		Down(nums, i, len(nums)-1)
	}
}

func Down(nums []int, k int, len int) {
	i := k
	for {
		if i >= len-1 {
			break
		}
		j := 2*i + 1
		if j+1 < len && nums[j+1] < nums[j] {
			j = j + 1
		}
		if nums[i] > nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		}
		i = j
	}
}

func Up(nums []int, k int, len int) {
	i := k
	for {
		if i <= 0 {
			break
		}

		j := (i - 1) / 2
		if nums[j] > nums[i] {
			nums[i], nums[j] = nums[j], nums[i]
		}
		i = j
	}
}
