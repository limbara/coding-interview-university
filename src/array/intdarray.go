package array

type IntDArray struct {
	array []int
	len   int
	cap   int
}

func NewIntDArray() *IntDArray {
	return &IntDArray{
		array: make([]int, 16),
		len:   16,
		cap:   16,
	}
}

func NewIntDArrayWithCap(cap int) *IntDArray {
	return &IntDArray{
		array: make([]int, cap),
		len:   cap,
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
	if index >= a.cap {
		panic("index out of bounds")
	}

	return a.array[index]
}

func (a *IntDArray) Push(item int) {
	if a.len == a.cap {
		newCap := a.cap * 2
		if newCap == 0 {
			newCap = 2
		}
		a.resize(newCap)
	}

	a.array[a.len] = item
	a.len++
}

func (a *IntDArray) Insert(index, item int) {
	if index >= a.cap {
		panic("index out of bounds")
	}

	a.Push(item)

	// shift items to the right
	pushIndex := a.len - 1
	i := pushIndex
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
	if a.len == 0 {
		panic("array is empty")
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
		panic("index out of bounds")
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

	newLen := a.len

	for i := range newLen {
		newArray[i] = a.array[i]
	}

	a.array = newArray
	a.cap = newCap
	a.len = newLen
}

func isQuarter(a, b int) bool {
	return a*4 == b
}
