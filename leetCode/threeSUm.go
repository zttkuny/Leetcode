package leetCode

import "sort"

var res [][]int
var path []int

// 回溯法
func ThreeSum(nums []int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	sort.Ints(nums)
	backwarding(nums, 0, 0)
	return res
}

func backwarding(nums []int, sum int, StartIndex int) {
	if len(path) == 3 {
		if sum == 0 {
			tmp := make([]int, 3)
			copy(tmp, path)
			res = append(res, tmp)
		}
		return
	}

	for i := StartIndex; i < len(nums); i++ {
		if i > StartIndex && nums[i] == nums[i-1] { //树层去重
			continue
		}
		path = append(path, nums[i])
		backwarding(nums, sum-nums[i], i+1)
		path = path[:len(path)-1]
	}
}

// 哈希法
func ThreeSum_hash(nums []int) [][]int {
	res := make([][]int, 0) //结果切片
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		set := make(map[int]struct{}, 0)

		for j := i + 1; j < len(nums); j++ {

			target := -nums[i] - nums[j]
			if _, ok := set[target]; ok {
				res = append(res, []int{nums[i], target, nums[j]})
				delete(set, target)
			} else {
				set[nums[j]] = struct{}{}
			}
		}
	}
	return res
}
