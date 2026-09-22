/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	nodeMap := make(map[*Node]*Node)
	
	ll := &Node{}
    newHead := ll
 
	headPtr := head

	for headPtr != nil {
		newNode := &Node{Val: headPtr.Val}
		newHead.Next = newNode

		nodeMap[headPtr] = newNode
		
		newHead = newHead.Next
		headPtr = headPtr.Next
	}

	newHead = ll.Next
	headPtr = head

	for headPtr != nil {
		targetNode := nodeMap[headPtr.Random]
		newHead.Random = targetNode

		newHead = newHead.Next
		headPtr = headPtr.Next
	}


	return ll.Next
}
