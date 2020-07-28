package sort

import (
	"fmt"
	"reflect"
	"testing"
)

type tdt struct {
	in  []int
	out []int
}

func sortTests() []tdt {
	return []tdt{
		{
			[]int{5, 1, 4, 2, 6, 3},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 2, 1, 2},
			[]int{1, 1, 2, 2},
		},
		{
			[]int{},
			[]int{},
		},
	}
}

func TestInsertSort(t *testing.T) {

	for _, tt := range sortTests() {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			if got, want := insertSort(tt.in), tt.out; !reflect.DeepEqual(want, got) {
				t.Errorf("got %v ,want %v", got, want)
			}
		})
	}

}


func TestMergeSort(t *testing.T) {

	for _, tt := range sortTests() {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			if got, want := mergeSort(tt.in), tt.out; !reflect.DeepEqual(want, got) {
				t.Errorf("got %v ,want %v", got, want)
			}
		})
	}
}
