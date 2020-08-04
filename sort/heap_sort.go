package sort

func heapSort(arr []int) []int {
  buildMaxHeap(arr)
  for i := len(arr) - 1; i > 0; i-- {
    arr[0], arr[i] = arr[i], arr[0]
    heapify(arr, 0, i)
  }
  return arr
}

func buildMaxHeap(arr []int) {
  for i := len(arr) / 2; i >= 0; i-- {
    heapify(arr, i, len(arr))
  }
}

func heapify(arr []int, idx int, length int) {
  max := idx
  left := 2*idx + 1
  right := 2*idx + 2

  if left < length && arr[left] > arr[max] {
    max = left
  }
  if right < length && arr[right] > arr[max] {
    max = right
  }

  if max != idx {
    arr[max], arr[idx] = arr[idx], arr[max]
    heapify(arr, max, length)
  }
}
