package array

import "errors"

type IntDArray struct {
	array []int
	len   int
	cap   int
}

var ErrIndexOutOfBound = errors.New("index out of bounds")
var ErrArrayEmpty = errors.New("array is empty")

func NewIntDArray() *IntDArray {
	return &IntDArray{
		array: make([]int, 16),
		len:   0,
		cap:   16,
	}
}

func NewIntDArrayWithCap(cap int) *IntDArray {
	return &IntDArray{
		array: make([]int, cap),
		len:   0,
		cap:   cap,
	}
}

func (a IntDArray) Size() int {
	return a.len
}

func (a IntDArray) Capacity() int {
	return a.cap
}

func (a IntDArray) IsEmpty() bool {
	return a.len == 0
}

func (a IntDArray) At(index int) int {
	if index >= a.len {
		panic(ErrIndexOutOfBound)
	}

	return a.array[index]
}

func (a *IntDArray) Push(item int) {
	if a.len == a.cap {
		newCap := max(2, a.cap*2)
		a.resize(newCap)
	}

	a.array[a.len] = item
	a.len++
}

func (a *IntDArray) Insert(index, item int) {
	if index >= a.cap {
		panic(ErrIndexOutOfBound)
	}

	a.Push(item)

	// shift items to the right
	i := a.len - 1
	for i > index {
		a.array[i] = a.array[i-1]
		i--
	}
	a.array[i] = item
}

func (a *IntDArray) Prepend(item int) {
	a.Insert(0, item)
}

func (a *IntDArray) Pop() int {
	if a.IsEmpty() {
		panic(ErrArrayEmpty)
	}

	value := a.array[a.len-1]
	a.array[a.len-1] = 0
	a.len--

	if isQuarter(a.len, a.cap) {
		a.resize(a.cap / 2)
	}

	return value
}

func (a *IntDArray) Delete(index int) int {
	if index >= a.len {
		panic(ErrIndexOutOfBound)
	}

	value := a.array[index]

	// shift items to the left
	for i := index; i < a.len-1; i++ {
		a.array[i] = a.array[i+1]
	}
	a.Pop()

	return value
}

func (a *IntDArray) Remove(item int) {
	j := 0

	for i := range a.len {
		if a.array[i] != item {
			a.array[j] = a.array[i]
			j++
		}
	}

	for i := j; i < a.len; i++ {
		a.array[i] = 0
	}

	a.len = j

	if isQuarter(a.len, a.cap) {
		a.resize(a.cap / 2)
	}
}

func (a IntDArray) Find(item int) int {
	for i := 0; i < a.len; i++ {
		if a.array[i] == item {
			return i
		}
	}

	return -1
}

func (a *IntDArray) resize(newCap int) {
	newArray := make([]int, newCap)

	for i := range a.len {
		newArray[i] = a.array[i]
	}
	a.array = newArray
	a.cap = newCap
}

func isQuarter(a, b int) bool {
	return a*4 == b
}
