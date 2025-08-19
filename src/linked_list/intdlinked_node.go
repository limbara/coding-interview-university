package linked_list

type IntDLinkedNode struct {
	Item int
	next *IntDLinkedNode
	prev *IntDLinkedNode
}

func NewIntDLinkedNode(item int, next *IntDLinkedNode, prev *IntDLinkedNode) *IntDLinkedNode {
	return &IntDLinkedNode{
		Item: item,
		next: next,
		prev: prev,
	}
}

func (i IntDLinkedNode) Next() *IntDLinkedNode {
	return i.next
}

func (i IntDLinkedNode) Prev() *IntDLinkedNode {
	return i.prev
}

func (i *IntDLinkedNode) InsertAfter(item int) {
	currentNext := i.next
	newNode := NewIntDLinkedNode(item, currentNext, i)
	i.next = newNode

	if currentNext != nil {
		currentNext.prev = newNode
	}
}

func (i *IntDLinkedNode) InsertBefore(item int) {
	currentPrev := i.prev
	newNode := NewIntDLinkedNode(item, i, currentPrev)
	i.prev = newNode

	if currentPrev != nil {
		currentPrev.next = newNode
	}
}

func (i *IntDLinkedNode) RemoveSelf() {
	prev := i.prev
	next := i.next

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}
}
