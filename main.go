package main

import (
	"LeetCode/leetCode"
	"fmt"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param str1 string字符串
 * @param str2 string字符串
 * @return int整型
 */

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func editDistance(str1 string, str2 string) int {

	dp := make([][]int, len(str1)+1)
	for i := range dp {
		dp[i] = make([]int, len(str2)+1)
		dp[i][0] = i
	}

	for i := 0; i < len(str2); i++ {
		dp[0][i] = i
	}

	for i := 1; i < len(str1)+1; i++ {
		for j := 1; j < len(str2)+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1]+1, min(dp[i-1][j]+1, dp[i][j-1]+1))
			}
		}
	}

	return dp[len(str1)][len(str2)]
	// write code here
}

func main() {
	intervals := [][]int{
		{1, 5},
		{2, 6},
		{4, 10},
		{10, 12},
		{3, 4},
		{7, 8},
		{8, 11},
		{9, 12},
		{13, 15},
		{14, 16},
	}

	fmt.Println(leetCode.MinArragedRooms(intervals))
}
