package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeWithCap(t *testing.T) {
	assert := assert.New(t)
	initialCap := 8
	intArray := NewIntDArrayWithCap(initialCap)

	assert.Equal(0, intArray.Size(), "Initial size is not %v", 0)
	assert.Equal(initialCap, intArray.Capacity(), "Initial capacity is not %v", initialCap)
	assert.Equal([]int{0, 0, 0, 0, 0, 0, 0, 0}, intArray.array, "Initial array items are not equal")
	assert.Equal(true, intArray.IsEmpty(), "Initial array should be empty.")
}

func TestInsert(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should results in correct order of items when insert", func(t *testing.T) {
		intArray := NewIntDArrayWithCap(4)

		tests := []struct {
			index, item int
			expected    []int
		}{
			{
				index:    0,
				item:     1,
				expected: []int{1, 0, 0, 0},
			},
			{
				index:    1,
				item:     2,
				expected: []int{1, 2, 0, 0},
			},
			{
				index:    0,
				item:     0,
				expected: []int{0, 1, 2, 0},
			},
			{
				index:    3,
				item:     3,
				expected: []int{0, 1, 2, 3},
			},
		}

		for _, test := range tests {
			intArray.Insert(test.index, test.item)
			assert.Equal(test.expected, intArray.array, "Insert item %v to index %v results in array items not equals.", test.index, test.item)
		}
	})

	t.Run("Should be able to grow double in capacity when capacity is not enough", func(t *testing.T) {
		intArray := NewIntDArrayWithCap(4)

		for v := range 5 {
			intArray.Insert(0, v)
		}

		assert.Equal(8, intArray.Capacity(), "capacity is not equals")

		for v := range 5 {
			intArray.Insert(0, v)
		}

		assert.Equal(16, intArray.Capacity(), "capacity is not equals")
	})

	t.Run("Should output the right size even when capacity is growing", func(t *testing.T) {
		intArray := NewIntDArrayWithCap(4)

		for v := range 5 {
			intArray.Insert(0, v)
		}

		assert.Equal(5, intArray.Size(), "size is not equals")

		for v := range 5 {
			intArray.Insert(0, v)
		}

		assert.Equal(10, intArray.Size(), "size is not equals")
	})

	t.Run("Should panic if insert is out of capacity", func(t *testing.T) {
		intArray := NewIntDArrayWithCap(4)

		assert.PanicsWithError(ErrIndexOutOfBound.Error(), func() {
			intArray.Insert(8, 5)
		}, "insert on index over the capacity should panic %v", ErrIndexOutOfBound.Error())
	})
}

func TestPreprend(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should results in correct order of items when prepend", func(t *testing.T) {
		intArray := NewIntDArrayWithCap(4)

		tests := []struct {
			item     int
			expected []int
		}{
			{
				item:     1,
				expected: []int{1, 0, 0, 0},
			},
			{
				item:     2,
				expected: []int{2, 1, 0, 0},
			},
			{
				item:     0,
				expected: []int{0, 2, 1, 0},
			},
			{
				item:     3,
				expected: []int{3, 0, 2, 1},
			},
		}

		for _, test := range tests {
			intArray.Prepend(test.item)
			assert.Equal(test.expected, intArray.array, "Preprend item %v results in array items not equals.", test.item)
		}
	})

	t.Run("Should be able to grow double in capacity when capacity is not enough", func(t *testing.T) {
		intArray := NewIntDArrayWithCap(4)

		for v := range 5 {
			intArray.Prepend(v)
		}

		assert.Equal(8, intArray.Capacity(), "capacity is not equals")

		for v := range 5 {
			intArray.Prepend(v)
		}

		assert.Equal(16, intArray.Capacity(), "capacity is not equals")
	})

	t.Run("Should output the right size even when capacity is growing", func(t *testing.T) {
		intArray := NewIntDArrayWithCap(4)

		for v := range 5 {
			intArray.Prepend(v)
		}

		assert.Equal(5, intArray.Size(), "size is not equals")

		for v := range 5 {
			intArray.Prepend(v)
		}

		assert.Equal(10, intArray.Size(), "size is not equals")
	})
}

func TestPush(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should results in correct order of items when push", func(t *testing.T) {
		intArray := NewIntDArrayWithCap(4)

		tests := []struct {
			item     int
			expected []int
		}{
			{
				item:     1,
				expected: []int{1, 0, 0, 0},
			},
			{
				item:     2,
				expected: []int{1, 2, 0, 0},
			},
			{
				item:     0,
				expected: []int{1, 2, 0, 0},
			},
			{
				item:     3,
				expected: []int{1, 2, 0, 3},
			},
		}

		for _, test := range tests {
			intArray.Push(test.item)
			assert.Equal(test.expected, intArray.array, "Preprend item %v results in array items not equals.", test.item)
		}
	})

	t.Run("Should be able to grow double in capacity when capacity is not enough", func(t *testing.T) {
		intArray := NewIntDArrayWithCap(4)

		for v := range 5 {
			intArray.Push(v)
		}

		assert.Equal(8, intArray.Capacity(), "capacity is not equals")

		for v := range 5 {
			intArray.Push(v)
		}

		assert.Equal(16, intArray.Capacity(), "capacity is not equals")
	})

	t.Run("Should output the right size even when capacity is growing", func(t *testing.T) {
		intArray := NewIntDArrayWithCap(4)

		for v := range 5 {
			intArray.Push(v)
		}

		assert.Equal(5, intArray.Size(), "size is not equals")

		for v := range 5 {
			intArray.Push(v)
		}

		assert.Equal(10, intArray.Size(), "size is not equals")
	})
}

func TestAt(t *testing.T) {
	assert := assert.New(t)
	intArray := NewIntDArrayWithCap(4)

	for v := range 5 {
		intArray.Push(v)
	}

	for v := range 5 {
		assert.Equal(v, intArray.At(v), "Array item at index %v is not %v", v, v)
	}
}

func TestPop(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should remove items from the tail of the array", func(t *testing.T) {
		intArray := NewIntDArrayWithCap(4)

		for v := range 5 {
			intArray.Push(v)
		}

		for _, expected := range []int{4, 3, 2, 1, 0} {
			assert.Equal(expected, intArray.Pop(), "Item popped doesn't match")
		}
	})

	t.Run("Should be able to shrink if the size is a quarter of current capacity", func(t *testing.T) {
		intArray := NewIntDArrayWithCap(4)

		for v := range 8 {
			intArray.Push(v)
		}

		assert.Equal(8, intArray.Size(), "Size doesn't match before shrink")
		assert.Equal(8, intArray.Capacity(), "Capacity doesn't match before shrink")

		for range 4 {
			intArray.Pop()
		}

		assert.Equal(4, intArray.Size(), "Size shouldn't shrink after 4 items pop")
		assert.Equal(8, intArray.Capacity(), "Capacity shouldn't shrink match 4 items pop")

		for range 2 {
			intArray.Pop()
		}

		assert.Equal(2, intArray.Size(), "Size doesn't match after shrink")
		assert.Equal(4, intArray.Capacity(), "Capacity doesn't match after shrink")
	})
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should result in the correct order of items when deleting", func(t *testing.T) {
		intArray := NewIntDArrayWithCap(4)

		for v := range 8 {
			intArray.Push(v)
		}

		tests := []struct {
			index    int
			expected []int
		}{
			{
				index:    1,
				expected: []int{0, 2, 3, 4, 5, 6, 7, 0},
			},
			{
				index:    2,
				expected: []int{0, 2, 4, 5, 6, 7, 0, 0},
			},
			{
				index:    0,
				expected: []int{2, 4, 5, 6, 7, 0, 0, 0},
			},
			{
				index:    3,
				expected: []int{2, 4, 5, 7, 0, 0, 0, 0},
			},
		}

		for _, test := range tests {
			intArray.Delete(test.index)
			assert.Equal(test.expected, intArray.array, "Delete index %v results in array items not equals.", test.index)
		}
	})

	t.Run("Should be able to shrink if the size is a quarter of current capacity", func(t *testing.T) {
		intArray := NewIntDArrayWithCap(4)

		for v := range 8 {
			intArray.Push(v)
		}

		assert.Equal(8, intArray.Size(), "Size doesn't match before shrink")
		assert.Equal(8, intArray.Capacity(), "Capacity doesn't match before shrink")

		for v := range 4 {
			intArray.Delete(v)
		}

		assert.Equal(4, intArray.Size(), "Size shouldn't shrink after 4 items deleted")
		assert.Equal(8, intArray.Capacity(), "Capacity shouldn't shrink match 4 items deleted")

		for v := range 2 {
			intArray.Delete(v)
		}

		assert.Equal(2, intArray.Size(), "Size doesn't match after shrink")
		assert.Equal(4, intArray.Capacity(), "Capacity doesn't match after shrink")
	})
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should result in the correct order of items when deleting", func(t *testing.T) {

		tests := []struct {
			input    []int
			remove   int
			expected []int
		}{
			{
				input:    []int{1, 2, 3, 4, 5, 6, 7, 8},
				remove:   1,
				expected: []int{2, 3, 4, 5, 6, 7, 8, 0},
			},
			{
				input:    []int{1, 2, 3, 4, 5, 6, 7, 8},
				remove:   8,
				expected: []int{1, 2, 3, 4, 5, 6, 7, 0},
			},
			{
				input:    []int{1, 2, 2, 2, 2, 2, 2, 8},
				remove:   2,
				expected: []int{1, 8, 0, 0},
			},
			{
				input:    []int{1, 2, 3, 4, 5, 6, 7, 8},
				remove:   9,
				expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
		}

		for _, test := range tests {
			intArray := NewIntDArrayWithCap(4)

			for _, v := range test.input {
				intArray.Push(v)
			}

			intArray.Remove(test.remove)

			assert.Equal(test.expected, intArray.array, "Array items content is not equal after removing %v", test.remove)
		}
	})
}

func TestFind(t *testing.T) {

	t.Run("Should return the index of item if found", func(t *testing.T) {
		assert := assert.New(t)
		intArray := NewIntDArrayWithCap(4)

		for v := range 8 {
			intArray.Push(v)
		}

		for v := range 8 {
			assert.Equal(v, intArray.Find(v), "Item index %v not found in array %v", v, intArray)
		}
	})

	t.Run("Should return -1 if item is not found", func(t *testing.T) {
		assert := assert.New(t)
		intArray := NewIntDArrayWithCap(4)

		for v := range 8 {
			intArray.Push(v)
		}

		assert.Equal(-1, intArray.Find(9), "Item 10 shouldn't be in the array", intArray)
		assert.Equal(-1, intArray.Find(10), "Item 10 shouldn't be in the array", intArray)
	})

}
