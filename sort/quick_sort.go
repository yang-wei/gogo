package sort 

// Hoare
func quickSort(nums []int) []int {
  _quickSort(nums, 0, len(nums) - 1)
  return nums
}

func _quickSort(nums []int, left, right int) {
  if left < right {
    pivotIdx := partition(nums, left, right)
    _quickSort(nums, left, pivotIdx)
    _quickSort(nums, pivotIdx + 1, right)
  }
}

func partition(nums []int, left, right int) int {
  pivot := nums[(left+right)/2]
  for {
    for nums[left] < pivot {
      left++
    }
    for nums[right] > pivot {
      right--
    }
    if left >= right {
      break
    }
    nums[left], nums[right] = nums[right], nums[left]
    left++
    right--
  }
  return right
}
