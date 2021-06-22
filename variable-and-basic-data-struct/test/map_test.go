package test

import (
	"testing"

	v "variable-and-basic-data-struct/assignment"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	myMap := v.NewMap()
	assert.Equal(t, myMap["a"], 5)
}

func TestMapRetrieveKey(t *testing.T) {
	myMap := v.NewMap()

	_, ok := myMap["a"]
	assert.True(t, ok)

	_, ok = myMap["asdas"]
	assert.False(t, ok)
}
