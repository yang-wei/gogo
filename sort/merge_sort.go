package sort

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	middle := len(arr) / 2
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])
	return merge(left, right)
}

func merge(left, right []int) []int {

	res := make([]int, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {

		if left[i] < right[j] {
			res[i+j] = left[i]
			i++
		} else {
			res[i+j] = right[j]
			j++
		}
	}
	for i < len(left) {
		res[i+j] = left[i]
		i++
	}
	for j < len(right) {
		res[i+j] = right[j]
		j++
	}
	return res
}
