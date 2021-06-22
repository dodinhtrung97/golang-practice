package test

import (
	"testing"

	v "variable-and-basic-data-struct/logicop"

	"github.com/stretchr/testify/assert"
)

func TestEnumImplicit(t *testing.T) {
	direction := v.North

	assert.Equal(t, direction, v.EnumEvalImplicit(direction))
}

func TestEnumExplicit(t *testing.T) {
	direction := v.North

	assert.Equal(t, direction, v.EnumEvalExplicit(direction))
}
