package linked_list

import "errors"

var ErrIndexOutOfBound = errors.New("index of out bounds")
var ErrUnexpectedPointer = errors.New("error unexpected pointer")
var ErrEmpty = errors.New("empty")

type IntSLinkedList struct {
	size int
	head *IntSLinkedNode
}

func NewIntSLinkedList() *IntSLinkedList {
	return &IntSLinkedList{
		size: 0,
		head: nil,
	}
}

func (s IntSLinkedList) Size() int {
	return s.size
}

func (s IntSLinkedList) Empty() bool {
	return s.size == 0
}

func (s IntSLinkedList) at(index int) (prev, current *IntSLinkedNode) {
	if index > s.size || index < 0 {
		panic(ErrIndexOutOfBound)
	}

	current = s.head
	for index > 0 && current != nil {
		prev = current
		current = current.next
		index -= 1
	}

	if index != 0 {
		panic(ErrUnexpectedPointer)
	}

	return prev, current
}

func (s IntSLinkedList) ValueAt(index int) int {
	_, current := s.at(index)

	if current == nil {
		panic(ErrIndexOutOfBound)
	}

	return current.Item
}

func (s *IntSLinkedList) PushFront(value int) {
	// if empty then there's no head
	if s.Empty() {
		s.head = NewIntSLinkedNode(value, nil)
	} else {
		s.head = NewIntSLinkedNode(value, s.head)
	}
	s.size += 1
}

func (s *IntSLinkedList) PopFront() int {
	if s.Empty() {
		panic(ErrEmpty)
	}

	node := s.head
	s.head = s.head.next
	s.size -= 1

	return node.Item
}

func (s *IntSLinkedList) PushBack(value int) {
	// if empty then there's no head
	if s.Empty() {
		s.head = NewIntSLinkedNode(value, nil)
	} else {
		pointer := s.head

		for pointer.next != nil {
			pointer = pointer.next
		}

		// pointer now holds the last item
		pointer.InsertAfter(value)
	}
	s.size += 1
}

func (s *IntSLinkedList) PopBack() int {
	if s.Empty() {
		panic(ErrEmpty)
	}

	prev := s.head

	for prev.next != nil && prev.next.next != nil {
		prev = prev.next
	}

	// if prev is head, then remove head
	if prev == s.head {
		s.head = nil
		s.size -= 1

		return prev.Item
	} else {
		// prev now holds the item where prev.next is the last item
		tail := prev.next
		prev.next = nil
		s.size -= 1

		return tail.Item
	}
}

func (s *IntSLinkedList) Front() int {
	if s.Empty() {
		panic(ErrEmpty)
	}

	return s.head.Item
}

func (s *IntSLinkedList) Back() int {
	if s.Empty() {
		panic(ErrEmpty)
	} else {
		pointer := s.head

		for pointer.next != nil {
			pointer = pointer.next
		}

		return pointer.Item
	}
}

func (s *IntSLinkedList) Insert(index, value int) {
	prev, current := s.at(index)

	if current == s.head {
		s.head = NewIntSLinkedNode(value, s.head)
	} else {
		prev.InsertAfter(value)
	}

	s.size += 1
}

func (s *IntSLinkedList) Erase(index int) {
	prev, current := s.at(index)

	if current == s.head {
		s.head = current.next
	} else {
		prev.next = current.next
	}
	s.size -= 1
}

func (s *IntSLinkedList) NthValueFromEnd(n int) int {
	indexFromFront := s.Size() - n

	if indexFromFront < 0 {
		panic(ErrIndexOutOfBound)
	} else {
		return s.ValueAt(indexFromFront)
	}
}

func (s *IntSLinkedList) Reverse() {
	var prev *IntSLinkedNode = nil
	pointer := s.head

	for pointer != nil {
		next := pointer.next
		pointer.next = prev
		prev = pointer
		pointer = next
	}

	s.head = prev
}

func (s *IntSLinkedList) RemoveValue(value int) bool {
	var prev *IntSLinkedNode = nil
	pointer := s.head

	for pointer != nil {
		if pointer.Item == value {
			if pointer == s.head {
				s.head = s.head.next
			} else {
				prev.next = pointer.next
			}
			s.size -= 1
			return true
		}
		prev = pointer
		pointer = pointer.next
	}

	return false
}

func MapSLinkedList(s *IntSLinkedList, mapFunc func(node *IntSLinkedNode)) {
	pointer := s.head
	for pointer != nil {
		mapFunc(pointer)
		pointer = pointer.next
	}
}
