
type Node struct{
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
    Capacity int
	Length   int
	LRU      *Node
	MRU      *Node
	Values   map[int]*Node
}

func (this *LRUCache) remove(node *Node){
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		this.LRU = this.LRU.Next
	}
	
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		this.MRU = this.MRU.Prev
	}
	node.Prev = nil
	node.Next = nil
}

func (this *LRUCache) append(node *Node){
	if this.MRU == nil {
		this.MRU = node;
		this.LRU = node;
		return
	}
	node.Prev = this.MRU
	this.MRU.Next = node;
	this.MRU = node;
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		Capacity: capacity,
		Length: 0,
		LRU: nil,
		MRU: nil,
		Values: make(map[int]*Node, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
    if node, exists := this.Values[key]; exists{
		this.remove(node)
		this.append(node)
		return node.Val
	} 
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, exists := this.Values[key]; exists{
		this.remove(node)
		this.append(node)
		node.Val = value
	} else {
		if this.Length == this.Capacity{
			delete(this.Values, this.LRU.Key)
			this.remove(this.LRU)
			this.Length--
		}		
		newNode := &Node{Key: key, Val: value}
		this.append(newNode)
		this.Values[key] = newNode;
		this.Length++
	}
}
