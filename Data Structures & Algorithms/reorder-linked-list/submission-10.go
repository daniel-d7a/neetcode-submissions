/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	count := 0
	temp := head
	for temp != nil {
		count++
		temp = temp.Next
	}	

	part1 := head
	temp = head 
	for i := 0; i < (count-1)/2; i++  {
		temp = temp.Next
	}
	second := temp.Next
	temp.Next = nil

    var prev *ListNode
    curr := second

	for curr != nil {
        temp := curr.Next
        curr.Next = prev
        prev = curr
        curr = temp
    }

	part2 := prev

	node := &ListNode{}

	for i := 0; i < count; i++ {
		if i % 2 == 0 {
			node.Next = part1
			part1 = part1.Next
		} else {
			node.Next = part2
			part2 = part2.Next
		}
		node = node.Next
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
