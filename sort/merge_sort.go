package sort

func mergeSort(l []int) []int {
  if len(l) <= 1 {
    return l
  }
  half := len(l) / 2
  left, right := mergeSort(l[:half]), mergeSort(l[half:])

  res := make([]int, 0, len(l))
  i, j := 0, 0
  for i < len(left) || j < len(right) {
      if i == len(left) {
          res = append(res, right[j:]...)
          break
      }

      if j == len(right) {
          res = append(res, left[i:]...)
          break
      }

      if left[i] < right[j] {
         res = append(res, left[i])
         i++
      } else {
         res = append(res, right[j])
         j++
      }
  }
  return res
}

