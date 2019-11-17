package leetcode

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	a := robSin(nums[1:len(nums)])
	b := robSin(nums[:len(nums)-1])
	if a < b {
		return b
	}
	return a
}

func robSin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i][0] = dp[i-1][1]
		if dp[i][0] < dp[i-1][0] {
			dp[i][0] = dp[i-1][0]
		}
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	if dp[len(dp)-1][0] > dp[len(dp)-1][1] {
		return dp[len(dp)-1][0]
	}
	return dp[len(dp)-1][1]
}
