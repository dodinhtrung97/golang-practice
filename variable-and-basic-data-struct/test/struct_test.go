package test

import (
	"testing"

	v "variable-and-basic-data-struct/assignment"

	"github.com/stretchr/testify/assert"
)

func TestCreateStudent(t *testing.T) {
	student := v.NewStudent()
	assert.Equal(t, (*student).Name, "Sky")

	(*student).Name = "Keiji"
	assert.Equal(t, (*student).Name, "Keiji")
}

func TestCreateStudentByNew(t *testing.T) {
	student := new(v.Student)
	assert.Equal(t, (*student).Name, "")

	(*student).Name = "Keiji"
	assert.Equal(t, (*student).Name, "Keiji")
}

// student without * doesn't indicate a pointer address in this case
// this is golang's syntactic sugar
// only applies for structs, maps and slices
func TestCreateStudentWithSyntaxSugar(t *testing.T) {
	student := new(v.Student)
	assert.Equal(t, student.Name, "")

	(*student).Name = "Keiji"
	assert.Equal(t, student.Name, "Keiji")
}
