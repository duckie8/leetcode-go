/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {

	var l, f = head, head

	for f != nil && f.Next != nil {
		l = l.Next
		f = f.Next.Next
	}

	if f != nil {
		l = l.Next
	}

	var left = head
	var right = reverse(l)

	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true

}

func reverse(a *ListNode) *ListNode {
	var pre, cur, nxt *ListNode
	pre = nil
	cur = a

	for cur != nil {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

// @lc code=end

