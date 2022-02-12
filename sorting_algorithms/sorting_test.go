package sorting_algorithms

import (
	"testing"
)

func arraysAndSorted() [][][]int {
	return [][][]int{
		{{3, 5, 2}, {2, 3, 5}},
	}
}

func TestQuicksort(t *testing.T) {
	for _, arrayAndSorted := range arraysAndSorted() {
		array := arrayAndSorted[0]
		sorted := arrayAndSorted[1]
		Quicksort(array)

		for idx := range array {
			if array[idx] != sorted[idx] {
				t.Errorf("Quicksort failed to sort array %v", array)
				break
			}
		}
	}
}

func TestHeapsort(t *testing.T) {
	for _, arrayAndSorted := range arraysAndSorted() {
		array := arrayAndSorted[0]
		sorted := arrayAndSorted[1]
		Heapsort(array)

		for idx := range array {
			if array[idx] != sorted[idx] {
				t.Errorf("Heapsort failed to sort array %v", array)
				break
			}
		}
	}
}

func TestMergesort(t *testing.T) {
	for _, arrayAndSorted := range arraysAndSorted() {
		array := arrayAndSorted[0]
		sorted := arrayAndSorted[1]
		Mergesort(array)

		for idx := range array {
			if array[idx] != sorted[idx] {
				t.Errorf("Mergesort failed to sort array %v", array)
				break
			}
		}
	}
}
