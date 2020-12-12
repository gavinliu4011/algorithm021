/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	courseMap := make(map[int]int, 0)
	result := []int{}
	for _, val := range nums1 {
		courseMap[val]++
	}
	for _, val := range nums2 {
		if courseMap[val] > 0 {
			result = append(result, val)
			courseMap[val]--
		}
	}
	return result
}

func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	index1, index2 := 0, 0
	result := []int{}
	for index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			result = append(result, nums1[index1])
			index1++
			index2++
		}
	}
	return result
}

// @lc code=end
