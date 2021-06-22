package test

import (
	"testing"

	v "variable-and-basic-data-struct/assignment"

	"github.com/stretchr/testify/assert"
)

func TestStaticArray(t *testing.T) {
	array := v.StaticArray()
	assert.NotNil(t, array)
}

func TestSliceCanAddMoreElem(t *testing.T) {
	slice := v.SliceCanAddMoreElem()
	assert.Equal(t, len(slice), 4)
}

func TestSliceUpToIndex4(t *testing.T) {
	slice := v.SliceUpToIndex4()
	assert.Equal(t, len(slice), 4)
}

func TestRemoveElementAtIndex(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	slice = v.RemoveElementAtIndex(2, slice)
	assert.Equal(t, len(slice), 5)
}
