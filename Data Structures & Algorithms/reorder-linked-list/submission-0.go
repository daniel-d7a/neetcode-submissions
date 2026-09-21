/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    part1 := &ListNode{} // from 0 to floor(n/2)
    part2 := &ListNode{} // from floor(n/2) to n


	count := 0
	temp := head
	for temp != nil {
		count++
		temp = temp.Next
	}	

	part1 = head
	temp = head 
	for i := 0; i < count/2 - 1; i++  {
		temp = temp.Next
	}
	part2 = reverseList(temp.Next)


	ll := &ListNode{}
	node := ll

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

	head = ll
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
