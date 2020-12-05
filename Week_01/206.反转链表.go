/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (71.03%)
 * Likes:    1377
 * Dislikes: 0
 * Total Accepted:    382.6K
 * Total Submissions: 538.5K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 *
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode
	curr := head
	for curr != nil {
		next = curr.Next // 将当前节点的原本下一个节点取出存放，否则覆盖无法找回
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

// @lc code=end
