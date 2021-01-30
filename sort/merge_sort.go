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


// in place
func _mergeSort(x []int) []int {
  mergeSortHelper(x, 0, len(x)-1)
  return x
}

func _mergeSortHelper(x []int, left, right int) {
  if left < right {
    mid := (left + right) / 2
    _mergeSortHelper(x, left, mid)
    _mergeSortHelper(x, mid+1, right)
    _merge(x, left, right)
  }
}

func _merge(x []int, left, right int) {
  
  size := right - left + 1
  tmp := make([]int, size) 
  
  // [4 8 9 1 2]
  // left: 0 right: 4
  mid := (left + right) / 2
  // mid = 2
  i, j := left, mid+1  // 0, 3
  
  for curr := 0; curr < size; curr++ {
      if i > mid {
        tmp[curr] = x[j]
        j++   
      } else if j > right {
        tmp[curr] = x[i]
        i++   
      } else {
        if x[i] < x[j] {
          tmp[curr] = x[i]
          i++
        } else {
          tmp[curr] = x[j] 
          j++
        }
      }
  }
  for i, j:= 0, left; j <= right; i, j = i+1, j+1 {
    x[j] = tmp[i] 
  }
}
