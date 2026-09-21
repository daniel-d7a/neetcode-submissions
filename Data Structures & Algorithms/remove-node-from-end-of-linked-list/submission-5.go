/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head.Next == nil {
		return nil
	}

	head = reverseList(head)

	temp := head
	for i := 0; i < n - 2; i++ {
		temp = temp.Next
	}

	if n == 1 && head.Next != nil {
		return reverseList(head.Next)
	} else if temp.Next != nil {
		del := temp.Next
		temp.Next = temp.Next.Next
		del.Next = nil
		return reverseList(head)
	} else {
		return nil
	}

}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head

    for curr != nil {
        temp := curr.Next
        curr.Next = prev
        prev = curr
        curr = temp
    }
    return prev
}
