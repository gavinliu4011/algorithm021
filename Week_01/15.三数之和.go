/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (30.13%)
 * Likes:    2785
 * Dislikes: 0
 * Total Accepted:    371.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？请你找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例：
 *
 * 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */

// @lc code=start
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for k := 0; k < len(nums)-2; k++ {
		if nums[k] > 0 {
			continue
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		l, r := k+1, len(nums)-1
		for l < r {
			sum := nums[k] + nums[l] + nums[r]
			if sum > 0 {
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < 0 {
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			} else {
				result = append(result, []int{nums[k], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
				for l < r && nums[l] == nums[l-1] {
					l++
				}

			}
		}
	}
	return result
}

func main() {
	nums := []int{-2, 0, 0, 2, 2}
	fmt.Print(threeSum(nums))
}

// @lc code=end
