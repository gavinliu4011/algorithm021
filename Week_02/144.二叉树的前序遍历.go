/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	helper(root,&result)
	return result
}

func helper(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	// 根-左-右
	*result = append(*result,root.Val)
	helper(root.Left,result)
	helper(root.Right,result)
}

// @lc code=end

