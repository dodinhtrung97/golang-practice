package assignment

type Student struct {
	Name    string
	Age     int
	Friends []string
}

func NewStudent() *Student {
	student := Student{
		Name:    "Sky",
		Age:     25,
		Friends: []string{"a", "b", "c"},
	}
	return &student
}

func ChangeStudentName(student *Student) *Student {
	student.Name = "Keiji"
	return student
}
