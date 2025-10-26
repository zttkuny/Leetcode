package leetCode

var Des []int

func MergeSort(arr []int, low, high int) {
	if low == high {
		return
	}
	mid := low + (high-low)/2
	MergeSort(arr, low, mid)
	MergeSort(arr, mid+1, high)
	Merge(arr, Des, low, mid, high)
}

func Merge(arr []int, des []int, low, mid, high int) {
	for i := low; i <= high; i++ {
		des[i] = arr[i]
	}
	i, j, k := low, mid+1, low
	for i <= mid && j <= high {
		if des[i] <= des[j] {
			arr[k] = des[i]
			i++
			k++
		} else {
			arr[k] = des[j]
			j++
			k++
		}

	}
	for ; i <= mid; i, k = i+1, k+1 {
		arr[k] = des[i]
	}
	for ; j <= high; j, k = j+1, k+1 {
		arr[k] = des[j]
	}
}
