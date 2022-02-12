package sorting_algorithms

func Mergesort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	var mid int = len(arr) / 2
	left := append([]int{}, arr[:mid]...)
	right := append([]int{}, arr[mid:]...)

	Mergesort(left)
	Mergesort(right)

	idxLeft, idxRight, idx := 0, 0, 0

	for idxLeft < len(left) && idxRight < len(right) {
		if left[idxLeft] < right[idxRight] {
			arr[idx] = left[idxLeft]
			idxLeft++
		} else {
			arr[idx] = right[idxRight]
			idxRight++
		}
		idx++
	}

	for idxLeft < len(left) {
		arr[idx] = left[idxLeft]
		idxLeft++
		idx++
	}

	for idxRight < len(right) {
		arr[idx] = right[idxRight]
		idxRight++
		idx++
	}
}
