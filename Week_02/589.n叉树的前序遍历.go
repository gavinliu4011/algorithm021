/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    result := []int{}
	helper(root,&result)
	return result
}


func helper(root *Node, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	for _, node := range root.Children {
		helper(node, result)
	}
}
// @lc code=end

