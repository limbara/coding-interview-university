package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeWithCap(t *testing.T) {
	assert := assert.New(t)
	initialCap := 8
	intArray := NewIntDArrayWithCap(initialCap)

	assert.Equal(initialCap, intArray.Size(), "Initial size is not %v", initialCap)
	assert.Equal(initialCap, intArray.Capacity(), "Initial capacity is not %v", initialCap)
	assert.Equal([]int{0, 0, 0, 0, 0, 0, 0, 0}, intArray.array, "Initial array items are not equal")
	assert.Equal(true, intArray.IsEmpty(), "Initial array is not empty")
}

func TestInsert(t *testing.T) {
	assert := assert.New(t)

}
