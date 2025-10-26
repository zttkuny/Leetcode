package leetCode

func MaxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))

	dp[1] = max(nums[1], nums[1]*nums[0])
	dp[0] = nums[0]

	res := max(dp[1], dp[0])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i], max(nums[i]*dp[i-1], nums[i]*nums[i-1]*dp[i-2]))
		res = max(res, dp[i])
	}
	return res
}
