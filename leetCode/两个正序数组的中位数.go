package leetCode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
	return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
	l1, l2 := 0, 0
	for k > 1 && l1 < len(nums1) && l2 < len(nums2) {
		mid1 := 0
		if (l1 + k/2 - 1) >= len(nums1) {
			mid1 = len(nums1) - 1
		} else if (l1 + k/2 - 1) < l1 {
			mid1 = l1
		} else {
			mid1 = l1 + k/2 - 1
		}
		mid2 := 0
		if (l2 + k/2 - 1) >= len(nums2) {
			mid2 = len(nums2) - 1
		} else if (l2 + k/2 - 1) < l2 {
			mid2 = l2
		} else {
			mid2 = l2 + k/2 - 1
		}

		if nums1[mid1] <= nums2[mid2] {
			k -= mid1 - l1 + 1
			l1 = mid1 + 1
		} else {
			k -= mid2 - l2 + 1
			l2 = mid2 + 1
		}
	}

	if l1 >= len(nums1) {
		return nums2[l2]
	}
	if l2 >= len(nums2) {
		return nums1[l1]
	}

	return min(nums1[l1], nums2[l2])
}
