package sorting_algorithms

func Quicksort(arr []int) {
	quicksortHelper(arr, 0, len(arr)-1)
}

func quicksortHelper(arr []int, low int, high int) {
	if low < high {
		partitionIndex := partition(arr, low, high)
		quicksortHelper(arr, low, partitionIndex-1)
		quicksortHelper(arr, partitionIndex+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}
