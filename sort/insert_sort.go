package sort

func insertSort(l []int) []int {
	for i := 1; i < len(l); i++ {
		for j := i; j > 0; j-- {
			if l[j] >= l[j-1] {
				break
			} else {
				l[j], l[j-1] = l[j-1], l[j]
			}
		}
	}
	return l
}
