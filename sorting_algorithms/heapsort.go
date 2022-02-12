package sorting_algorithms

import "fmt"

func maxHeapify(arr []int, start int, end int) {
	largestIdx := start
	leftChildIdx := 2*start + 1
	rightChildIdx := 2*start + 2

	for _, idx := range []int{leftChildIdx, rightChildIdx} {
		if idx < end && arr[largestIdx] < arr[idx] {
			largestIdx = idx
		}
	}

	if largestIdx != start {
		arr[start], arr[largestIdx] = arr[largestIdx], arr[start]
		maxHeapify(arr, largestIdx, end)
	}
}

func Heapsort(arr []int) {
	arrLength := len(arr)

	for i := int(arrLength / 2); i >= 0; i-- {
		maxHeapify(arr, i, arrLength)
		fmt.Println("arr", i, arr)
	}

	for i := arrLength - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		maxHeapify(arr, 0, i)
	}
}
