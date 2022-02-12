package sorting_algorithms

import (
	"testing"
)

func arraysAndSorted() [][][]int {
	return [][][]int{
		{{3, 5, 2}, {2, 3, 5}},
		{{}, {}},
		{{1, 2, 3, 4}, {1, 2, 3, 4}},
		{{-1, -2, -3, -4}, {-4, -3, -2, -1}},
		{{99, -11, 5, 2, 19, 37, -1000}, {-1000, -11, 2, 5, 19, 37, 99}},
	}
}

func TestQuicksort(t *testing.T) {
	for _, arrayAndSorted := range arraysAndSorted() {
		array := arrayAndSorted[0]
		arrayOriginal := append([]int{}, array...)
		sorted := arrayAndSorted[1]
		Quicksort(array)

		for idx := range array {
			if array[idx] != sorted[idx] {
				t.Errorf(
					"Quicksort failed to sort array %v.\nExpected %v\nRecieved %v",
					arrayOriginal,
					sorted,
					array,
				)
				break
			}
		}
	}
}

func TestHeapsort(t *testing.T) {
	for _, arrayAndSorted := range arraysAndSorted() {
		array := arrayAndSorted[0]
		arrayOriginal := append([]int{}, array...)
		sorted := arrayAndSorted[1]
		Heapsort(array)

		for idx := range array {
			if array[idx] != sorted[idx] {
				t.Errorf(
					"Heapsort failed to sort array %v.\nExpected %v\nRecieved %v",
					arrayOriginal,
					sorted,
					array,
				)
				break
			}
		}
	}
}

func TestMergesort(t *testing.T) {
	for _, arrayAndSorted := range arraysAndSorted() {
		array := arrayAndSorted[0]
		arrayOriginal := append([]int{}, array...)
		sorted := arrayAndSorted[1]
		Mergesort(array)

		for idx := range array {
			if array[idx] != sorted[idx] {
				t.Errorf(
					"Mergesort failed to sort array %v.\nExpected %v\nRecieved %v",
					arrayOriginal,
					sorted,
					array,
				)
				break
			}
		}
	}
}
