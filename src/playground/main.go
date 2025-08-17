package main

import (
	"data/array"
	"fmt"
)

func main() {
	a := array.NewIntDArrayWithCap(4)
	a.Push(5)
	fmt.Println(a)
	a.Insert(0, 1)
	fmt.Println(a)
	a.Prepend(2)
	fmt.Println(a)

	fmt.Println(a.Pop(), a.Pop())

	fmt.Println(a)

	a.Prepend(2)
	fmt.Println(a)
	a.Insert(0, 1)
	a.Push(4)
	fmt.Println(a)
	a.Push(3)
	a.Push(4)
	a.Push(4)
	a.Push(5)
	fmt.Println(a)

	a.Remove(4)
	fmt.Println(a)
	a.Remove(3)
	fmt.Println(a)
	a.Remove(2)
	fmt.Println(a)
	// found := a.Find(3)

	// fmt.Println(found)

	// intArray := array.NewIntDArrayWithCap(4)

	// for v := range 8 {
	// 	intArray.Push(v)
	// }

	// fmt.Println(intArray)
}
