/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
package main

func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil {
		return []int{}
	}
	window, res := []int{}, []int{}
	for i := range nums {
		// 当下标大于等于窗口中k个数字时，需要判断当前下标减去当前窗口中最大值得下标是否大于 k 个数字
		if i >= k && i-window[0] >= k {
			window = window[1:] // 向后移动窗口
		}
		for len(window) > 0 && nums[i] >= nums[window[len(window)-1]] {
			// 遍历过程,当前元素比队列中的元素大，就将原来队列中的元素祭天
			window = window[:len(window)-1]
		}
		window = append(window, i)
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}

// func main() {
// 	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
// 	fmt.Print(maxSlidingWindow(nums, 3))
// }

// @lc code=end
