package main

import (
	"data/linked_list"
	"fmt"
)

func main() {
	dlinkedList := linked_list.NewIntDTLinkedList()

	// printFront := func() {
	// 	defer func() {
	// 		if r := recover(); r != nil {
	// 			fmt.Printf("Front of the linked list is %v\n", r)
	// 		}
	// 	}()
	// 	fmt.Printf("Front of the linked list is %v\n", dlinkedList.Front())
	// }

	// printBack := func() {
	// 	defer func() {
	// 		if r := recover(); r != nil {
	// 			fmt.Printf("Back of the linked list is %v\n", r)
	// 		}
	// 	}()
	// 	fmt.Printf("Back of the linked list is %v\n", dlinkedList.Back())
	// }

	// printAt := func(index int) {
	// 	defer func() {
	// 		if r := recover(); r != nil {
	// 			fmt.Printf("At %v results in %v\n", index, r)
	// 		}
	// 	}()
	// 	prev, current := dlinkedList.At(index)
	// 	fmt.Printf("At index %v of the linked list prev is %v and current is %v\n", index, prev, current)
	// }

	// fmt.Printf("Linked list size is %v\n", dlinkedList.Size())
	// fmt.Printf("Is linked list empty: %v\n", dlinkedList.Empty())

	// for v := range 5 {
	// 	fmt.Printf("PushFront %v\n", v)
	// 	dlinkedList.PushFront(v)
	// }
	// printLinkedList(dlinkedList)

	// fmt.Printf("Front of the linked list is %v\n", dlinkedList.Front())
	// fmt.Printf("Back of the linked list is %v\n", dlinkedList.Back())

	// for v := range 5 {
	// 	fmt.Printf("ValueAt index %v is %v\n", v, dlinkedList.ValueAt(v))
	// }

	// printLinkedList(dlinkedList)

	// for range 5 {
	// 	fmt.Printf("PopFront from the linked list results in %v\n", dlinkedList.PopFront())
	// }
	// printLinkedList(dlinkedList)
	// fmt.Printf("Is linked list empty: %v\n", dlinkedList.Empty())

	// printFront()

	// printBack()

	for v := range 5 {
		fmt.Printf("PushBack %v to linked list\n", v)
		dlinkedList.PushBack(v)
	}
	printLinkedList(dlinkedList)
	// printAt(0)
	// printAt(5)

	// printFront()
	// printBack()

	for range 5 {
		fmt.Printf("PopBack from the linked list results in %v\n", dlinkedList.PopBack())
	}
	printLinkedList(dlinkedList)
	// printFront()
	// printBack()

	for v := range 5 {
		dlinkedList.Insert(0, v)
	}
	printLinkedList(dlinkedList)
	dlinkedList.Insert(0, 10)
	printLinkedList(dlinkedList)

	dlinkedList.Erase(4)
	printLinkedList(dlinkedList)

	fmt.Printf("2 Nth value from the end of the linked list is %v\n", dlinkedList.NthValueFromEnd(2))

	dlinkedList.Reverse()
	printLinkedList(dlinkedList)

	fmt.Printf("2 Nth value from the end of the linked list is %v\n", dlinkedList.NthValueFromEnd(2))

	fmt.Println("RemoveValue 0", dlinkedList.RemoveValue(0))
	printLinkedList(dlinkedList)
	fmt.Println("RemoveValue 2", dlinkedList.RemoveValue(2))
	printLinkedList(dlinkedList)
	fmt.Println("RemoveValue 3", dlinkedList.RemoveValue(3))
	printLinkedList(dlinkedList)
	fmt.Println("RemoveValue 4", dlinkedList.RemoveValue(4))
	printLinkedList(dlinkedList)
	fmt.Println("RemoveValue 10", dlinkedList.RemoveValue(10))
	printLinkedList(dlinkedList)

	for v := range 5 {
		dlinkedList.PushBack(v)
	}
	printLinkedList(dlinkedList)

	fmt.Println("Erase index 4")
	dlinkedList.Erase(4)
	printLinkedList(dlinkedList)
}

func printLinkedList(l *linked_list.IntDTLinkedList) {
	fmt.Print("list now holds:\t\n")
	linked_list.MapDTLinkedList(l, func(node *linked_list.IntDLinkedNode) {
		fmt.Printf("%v\t", node.Item)
	})
	fmt.Printf("\nsize :\t%v\n\n", l.Size())
}
