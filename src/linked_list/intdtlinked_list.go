package linked_list

type IntDTLinkedList struct {
	size int
	head *IntDLinkedNode
	tail *IntDLinkedNode
}

func NewIntDTLinkedList() *IntDTLinkedList {
	return &IntDTLinkedList{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (l IntDTLinkedList) Size() int {
	return l.size
}

func (l IntDTLinkedList) Empty() bool {
	return l.size == 0
}

func (l IntDTLinkedList) At(index int) (prev, current *IntDLinkedNode) {
	if index > l.size || index < 0 {
		panic(ErrIndexOutOfBound)
	}

	// if the index falls in the first half then use head and vice versa
	if index*2 <= l.size {
		current = l.head
		for index > 0 && current != nil {
			prev = current
			current = current.next
			index -= 1
		}

		if index != 0 {
			panic(ErrUnexpectedPointer)
		}

		return prev, current
	} else {
		if index == l.size {
			return l.tail, nil
		}

		indexFromBack := l.size - index - 1
		current = l.tail
		prev = current.prev

		for indexFromBack > 0 && current != nil {
			current = current.prev
			prev = current.prev
			indexFromBack -= 1
		}

		if indexFromBack != 0 {
			panic(ErrUnexpectedPointer)
		}

		return prev, current
	}

}

func (l IntDTLinkedList) ValueAt(index int) int {
	_, current := l.At(index)

	if current == nil {
		panic(ErrIndexOutOfBound)
	}

	return current.Item
}

func (l *IntDTLinkedList) PushFront(value int) {
	if l.head == nil {
		newNode := NewIntDLinkedNode(value, nil, nil)
		l.head = newNode
		l.tail = newNode
	} else {
		l.head.InsertBefore(value)
		l.head = l.head.prev
	}

	l.size += 1
}

func (l *IntDTLinkedList) PopFront() int {
	if l.Empty() {
		panic(ErrEmpty)
	}

	node := l.head
	l.head = l.head.next

	if l.head != nil {
		l.head.prev = nil
	}

	if node == l.tail {
		l.tail = nil
	}

	l.size -= 1

	return node.Item
}

func (l *IntDTLinkedList) PushBack(value int) {
	if l.head == nil {
		newNode := NewIntDLinkedNode(value, nil, nil)
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.InsertAfter(value)
		l.tail = l.tail.next
	}

	l.size += 1
}

func (l *IntDTLinkedList) PopBack() int {
	if l.Empty() {
		panic(ErrEmpty)
	}

	node := l.tail
	l.tail = l.tail.prev

	if l.tail != nil {
		l.tail.next = nil
	}

	if node == l.head {
		l.head = nil
	}

	l.size -= 1

	return node.Item
}

func (l IntDTLinkedList) Front() int {
	if l.Empty() {
		panic(ErrEmpty)
	}

	return l.head.Item
}

func (l IntDTLinkedList) Back() int {
	if l.Empty() {
		panic(ErrEmpty)
	}

	return l.tail.Item
}

func (l *IntDTLinkedList) Insert(index, value int) {
	prev, current := l.At(index)

	if prev == nil && current == l.head {
		l.PushFront(value)
	} else if prev == l.tail && current == nil {
		l.PushBack(value)
	} else {
		current.InsertBefore(value)
		l.size += 1
	}
}

func (l *IntDTLinkedList) Erase(index int) {
	if l.Empty() {
		panic(ErrEmpty)
	}

	_, current := l.At(index)

	if current == nil {
		panic(ErrIndexOutOfBound)
	}

	switch current {
	case l.head:
		l.PopFront()
	case l.tail:
		l.PopBack()
	default:
		current.RemoveSelf()
		l.size -= 1
	}
}

func (l IntDTLinkedList) NthValueFromEnd(n int) int {
	indexFromFront := l.Size() - n
	_, current := l.At(indexFromFront)

	if current == nil {
		panic(ErrIndexOutOfBound)
	}

	return current.Item
}

func (l *IntDTLinkedList) Reverse() {
	pointer := l.head
	for pointer != nil {
		next := pointer.next
		pointer.next = pointer.prev
		pointer.prev = next
		pointer = next
	}

	temp := l.head
	l.head = l.tail
	l.tail = temp
}

func (l *IntDTLinkedList) RemoveValue(value int) bool {
	pointer := l.head

	for pointer != nil {
		if pointer.Item == value {
			switch pointer {
			case l.head:
				l.PopFront()
			case l.tail:
				l.PopBack()
			default:
				pointer.RemoveSelf()
				l.size -= 1
			}
			return true
		}
		pointer = pointer.next
	}

	return false
}

func MapDTLinkedList(s *IntDTLinkedList, mapFunc func(node *IntDLinkedNode)) {
	pointer := s.head
	for pointer != nil {
		mapFunc(pointer)
		pointer = pointer.next
	}
}
