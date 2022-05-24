/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}

	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head

}

var s *ListNode

func reverseN(head *ListNode, before int) *ListNode {
	if before == 1 {
		s = head.Next
		return head
	}

	var p = reverseN(head.Next, before-1)
	head.Next.Next = head
	head.Next = s
	return p
}

// @lc code=end

