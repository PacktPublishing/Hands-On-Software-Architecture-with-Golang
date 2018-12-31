package main

import "math/rand"


func quickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	left, right := 0, len(array) - 1

	// Pick a pivot  randomly and move it to the end
	pivot:= rand.Int() % len(array)
	array[pivot], array[right] = array[right], array[pivot]

	// Partition
	for i := range array {
		if array[i] < array[right] {
			array[i], array[left] = array[left], array[i]
			left++
		}
	}

	// Put the pivot in place
	array[left], array[right] = array[right], array[left]

	// Recurse
	quickSort(array[:left])
	quickSort(array[left + 1:])

	return array
}
