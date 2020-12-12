/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
	n1:= m-1
	n2 :=n-1
	length := m+n-1
	for n2>=0{
		if n1>=0&&nums1[n1]>nums2[n2]{
			nums1[length]= nums1[n1]
			n1--
		}else {
			nums1[length]= nums2[n2]
			n2--
		}
		length--
	}
}
// @lc code=end


// func main() {
// 	nums1 := []int{1,2,3,0,0,0,0}
// 	nums2 := []int{2,5,6,1}
// 	merge(nums1,3,nums2,4)
// }

