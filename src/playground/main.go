package main

import (
	"data/linked_list"
	"fmt"
)

func main() {
	// array := array.NewIntDArray()
	linkedList := linked_list.NewIntSLinkedList()
	fmt.Printf("Is Linked list empty %v\n", linkedList.Empty())
	fmt.Printf("Linked list size %v\n", linkedList.Size())

	linkedList.Insert(0, 2)
	linkedList.PushFront(1)
	linkedList.PushBack(3)
	linkedList.Insert(3, 4)

	fmt.Printf("Is Linked list empty %v\n", linkedList.Empty())

	printLinkedList(linkedList)

	for i := range linkedList.Size() {
		fmt.Printf("Value at index %v is %v\n", i, linkedList.ValueAt(i))
	}

	fmt.Printf("Front of the linked list is %v\n", linkedList.Front())
	fmt.Printf("Back of the linked list is %v\n", linkedList.Back())

	fmt.Printf("Pop Front of the linked list returns %v\n", linkedList.PopFront())
	printLinkedList(linkedList)

	fmt.Printf("Pop Back of the linked list returns %v\n", linkedList.PopBack())
	printLinkedList(linkedList)

	for _, v := range []int{1, 2, 3, 4} {
		fmt.Printf("PushBack %v to linked list\n", v)
		linkedList.PushBack(v)
	}

	printLinkedList(linkedList)

	fmt.Printf("2 nth value from the back of the linked list is %v\n\n", linkedList.NthValueFromEnd(2))
	fmt.Printf("1 nth value from the back of the linked list is %v\n\n", linkedList.NthValueFromEnd(1))

	for i := range 3 {
		fmt.Printf("Erase index %v node\n", i)
		linkedList.Erase(i)
		printLinkedList(linkedList)
	}

	for _, v := range []int{1, 2, 3, 4} {
		fmt.Printf("PushBack %v to linked list\n", v)
		linkedList.PushBack(v)
	}
	printLinkedList(linkedList)

	fmt.Println("Reversing the linked list")
	linkedList.Reverse()
	printLinkedList(linkedList)

	fmt.Println("RemoveValue 4 from linked list")
	linkedList.RemoveValue(4)
	printLinkedList(linkedList)

	fmt.Println("RemoveValue 4 from linked list")
	linkedList.RemoveValue(4)
	printLinkedList(linkedList)

	fmt.Println("RemoveValue 4 from linked list")
	linkedList.RemoveValue(4)
	printLinkedList(linkedList)
}

func printLinkedList(l *linked_list.IntSLinkedList) {
	fmt.Print("list now holds:\t")
	linked_list.MapSLinkedList(l, func(node *linked_list.IntSLinkedNode) {
		fmt.Printf("%v\t", node.Item)
	})
	fmt.Printf("\nsize :\t%v\n\n", l.Size())
}
