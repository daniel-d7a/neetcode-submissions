/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
	ll := &ListNode{}
	solution := ll

	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0

		if l1 != nil{
			sum = sum + l1.Val
			l1 = l1.Next
		} 
		if l2 != nil{
			sum = sum + l2.Val
			l2 = l2.Next
		} 

		total := sum + carry
 		solution.Next = &ListNode{Val: total % 10}

		carry = total / 10

		solution = solution.Next		
	}

	if carry != 0 {
		solution.Next = &ListNode{Val: carry}
	}

	return ll.Next
}
