package maxheap

import (
	"reflect"
	"testing"
)

func TestMaxHeapify1(t *testing.T) {

	r1 := maxHeap{
		[]int{1, 4, 3},
	}
	r1.heapify(1)
	if e := []int{4, 1, 3}; !reflect.DeepEqual(r1.arr, e) {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

	r2 := maxHeap{
		[]int{4, 1, 3},
	}
	r2.heapify(1)
	if e := []int{4, 1, 3}; !reflect.DeepEqual(r2.arr, e) {
		t.Fatalf("Expected %v but got %v", e, r2.arr)
	}

	r3 := maxHeap{
		[]int{3, 1, 4},
	}
	r3.heapify(1)
	if e := []int{4, 1, 3}; !reflect.DeepEqual(r3.arr, e) {
		t.Fatalf("Expected %v but got %v", e, r3.arr)
	}

}

func TestMaxHeapify2(t *testing.T) {

	r1 := maxHeap{
		[]int{5, 6, 3, 1, 7, 8, 2},
	}
	r1.heapify(1)
	if e := []int{6, 7, 3, 1, 5, 8, 2}; !reflect.DeepEqual(r1.arr, e) {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

	r2 := maxHeap{
		[]int{5, 6, 3, 1, 7, 8, 2},
	}
	r2.heapify(2)
	if e := []int{5, 7, 3, 1, 6, 8, 2}; !reflect.DeepEqual(r2.arr, e) {
		t.Fatalf("Expected %v but got %v", e, r2.arr)
	}

	r3 := maxHeap{
		[]int{5, 6, 3, 1, 7, 8, 2},
	}
	r3.heapify(3)
	if e := []int{5, 6, 8, 1, 7, 3, 2}; !reflect.DeepEqual(r3.arr, e) {
		t.Fatalf("Expected %v but got %v", e, r3.arr)
	}

}

func TestBuildMaxHeapify(t *testing.T) {

	r1 := NewMaxHeap([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7})
	if e := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}; !reflect.DeepEqual(r1.arr, e) {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}
}

func TestInsert(t *testing.T) {

	r1 := NewMaxHeap([]int{4, 1, 3})
	r1.insert(5)
	if e := []int{5, 4, 3, 1}; !reflect.DeepEqual(r1.arr, e) {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

	r2 := NewMaxHeap([]int{4, 1, 3})
	r2.insert(2)
	if e := []int{5, 2, 3, 1}; !reflect.DeepEqual(r2.arr, e) {
		t.Fatalf("Expected %v but got %v", e, r2.arr)
	}

	r3 := NewMaxHeap([]int{4, 1, 3})
	r3.insert(0)
	if e := []int{4, 1, 3, 0}; !reflect.DeepEqual(r3.arr, e) {
		t.Fatalf("Expected %v but got %v", e, r3.arr)
	}
}

func TestInsert2(t *testing.T) {

	r1 := NewMaxHeap([]int{4})
	r1.insert(19)
	if e := []int{19, 4}; !reflect.DeepEqual(r1.arr, e) {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

	r1.insert(8)
	if e := []int{19, 4, 8}; !reflect.DeepEqual(r1.arr, e) {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

	r1.insert(11)
	if e := []int{19, 11, 8, 4}; !reflect.DeepEqual(r1.arr, e) {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

	r1.insert(20)
	if e := []int{20, 19, 8, 4, 11}; !reflect.DeepEqual(r1.arr, e) {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

	if a, e := r1.getMax(), 20; a != e {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

}

func TestExtractMax(t *testing.T) {

	r1 := NewMaxHeap([]int{21, 19, 20, 4, 11})
	if a, e := r1.extractMax(), 21; a != e {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

	if a, e := r1.arr, []int{20, 19, 11, 4}; !reflect.DeepEqual(a, e) {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

	if a, e := r1.extractMax(), 20; a != e {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

	if a, e := r1.arr, []int{19, 4, 11}; !reflect.DeepEqual(a, e) {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

	if a, e := r1.extractMax(), 19; a != e {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

	if a, e := r1.extractMax(), 11; a != e {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

	if a, e := r1.extractMax(), 4; a != e {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

	if a, e := r1.extractMax(), -1; a != e {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}

}

func TestSort(t *testing.T) {

	r1 := NewMaxHeap([]int{21, 19, 20, 4, 11})
	r1.sort()

	if a, e := r1.arr, []int{21, 20, 19, 11, 4}; !reflect.DeepEqual(a, e) {
		t.Fatalf("Expected %v but got %v", e, r1.arr)
	}
}
