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

	a := head
	b := head.Next
	c := head.Next.Next

	b.Next = a
	
	for c != nil {
		a = b
		b = c
		c = c.Next
		
		b.Next = a
	}

	head.Next = nil
	return b
}
