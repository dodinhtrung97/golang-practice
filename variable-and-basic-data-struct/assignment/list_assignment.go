package assignment

// StaticArray returns a static array
func StaticArray() [3]int {
	newArray := [...]int{1, 2, 3}
	return newArray
}

// SliceCanAddMoreElem should prove that a slice can be appended
func SliceCanAddMoreElem() []int {
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	return slice
}

// SliceUpToIndex4 should slice an array of size 6 to size 4
func SliceUpToIndex4() []int {
	slice := []int{1, 2, 3, 4, 5, 6}
	return slice[:4]
}

// RemoveElementAtIndex takes int x and a slice and remove element at index x
func RemoveElementAtIndex(x int, slice []int) []int {
	slice = append(slice[:x], slice[x+1:]...)

	return slice
}
