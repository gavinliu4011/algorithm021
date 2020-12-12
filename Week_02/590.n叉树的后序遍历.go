/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	result := []int{}
	helper(root,&result)
	return result
}

func helper(root *Node, result *[]int) {
	if root == nil {
		return
	}
	for _, node := range root.Children {
		helper(node, result)
	}
	*result = append(*result, root.Val)
}

// @lc code=end

