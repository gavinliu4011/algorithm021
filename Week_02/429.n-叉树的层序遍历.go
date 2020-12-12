/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := [][]int{}
	helper2(root, 0, &result)
	return result
}

func helper2(root *Node, level int, result *[][]int) {
	if root == nil{
        return
    }
    if len(*result) == level{
        *result = append(*result, []int{})
	}
	(*result)[level] = append((*result)[level], root.Val)
	for _, children := range root.Children {
		helper2(children, level+1, result)
	}
}

// @lc code=end
