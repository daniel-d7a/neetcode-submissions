/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var a *ListNode = nil
	b := head

	for b != nil {
		c := b.Next
		b.Next = a
		a = b
		b = c
	}

	return a
}
