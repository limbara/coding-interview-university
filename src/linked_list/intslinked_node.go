package linked_list

type IntSLinkedNode struct {
	Item int
	next *IntSLinkedNode
}

func NewIntSLinkedNode(item int, next *IntSLinkedNode) *IntSLinkedNode {
	return &IntSLinkedNode{
		Item: item,
		next: next,
	}
}

func (i IntSLinkedNode) Next() *IntSLinkedNode {
	return i.next
}

func (i *IntSLinkedNode) InsertAfter(item int) {
	node := NewIntSLinkedNode(item, i.next)
	i.next = node
}
