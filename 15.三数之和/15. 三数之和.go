package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			tmp := nums[i] + nums[j] + nums[k]
			switch {
			case tmp > 0:
				k -= 1
			case tmp < 0:
				j += 1
			case tmp == 0:
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for {
					k -= 1
					if k == 0 || nums[k+1] != nums[k] {
						break
					}
				}
				for {
					j += 1
					if j == len(nums)-1 || nums[j-1] != nums[j] {
						break
					}
				}
			}
		}
	}
	return res
}
