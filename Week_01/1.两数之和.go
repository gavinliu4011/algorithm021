/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 *
 * https://leetcode-cn.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (49.41%)
 * Likes:    9719
 * Dislikes: 0
 * Total Accepted:    1.6M
 * Total Submissions: 3.2M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 *
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
 *
 *
 *
 * 示例:
 *
 * 给定 nums = [2, 7, 11, 15], target = 9
 *
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 *
 *
 */

// @lc code=start
package main

func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := sumMap[target-num]; ok {
			return []int{j, i}
		}
		sumMap[num] = i
	}

	return nil
}

// @lc code=end
