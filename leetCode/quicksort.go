package leetCode

func QuickSort(arr []int, left, right int) {
	if right <= left {
		return
	}

	pivot := parition(arr, left, right)
	QuickSort(arr, left, pivot-1)
	QuickSort(arr, pivot+1, right)
}

func parition(arr []int, l, r int) int {
	l1, r1 := l, r

	for l1 < r1 {
		//
		for ; l1 < r1 && arr[r1] >= arr[l]; r1-- {
		}

		for ; l1 < r1 && arr[l1] <= arr[l]; l1++ {
		}
		arr[l1], arr[r1] = arr[r1], arr[l1]
	}
	arr[l], arr[l1] = arr[l1], arr[l]
	return l1
}
