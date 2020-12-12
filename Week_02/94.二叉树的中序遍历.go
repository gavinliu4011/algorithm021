/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	helper(root,&result)
	return result
}

func helper(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	// 左
	helper(root.Left, result)
	// 根
	*result = append(*result, root.Val)
	// 右
	helper(root.Right,result)
}

// @lc code=end
