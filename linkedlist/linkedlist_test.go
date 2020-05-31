package linkedlist

import (
	"testing"
)

func TestAddAtHead(t *testing.T) {
	obj := new(MyLinkedList)

	obj.AddAtHead(1)
	obj.AddAtHead(2)

	if a, e := obj.toString(), "2->1"; a != e {
		t.Fatalf("Expected %v but got %v", e, a)
	}
}

func TestAddAtTail(t *testing.T) {
	obj := new(MyLinkedList)

	obj.AddAtTail(1)
	obj.AddAtTail(2)

	if a, e := obj.toString(), "1->2"; a != e {
		t.Fatalf("Expected %v but got %v", e, a)
	}
}

func TestAddAtIndex(t *testing.T) {
	obj := MyLinkedList{
    Node: &Node{1, &Node{3, nil}},
  }
  obj.AddAtIndex(1, 2)

	if a, e := obj.toString(), "1->2->3"; a != e {
		t.Fatalf("Expected %v but got %v", e, a)
	}
}

func TestDeleteAtIndex(t *testing.T) {
	obj := MyLinkedList{
    Node: &Node{1, &Node{2, &Node{3, nil}}},
  }
  obj.DeleteAtIndex(1)

	if a, e := obj.toString(), "1->3"; a != e {
		t.Fatalf("Expected %v but got %v", e, a)
	}
}

func TestGet(t *testing.T) {
	obj := MyLinkedList{
    Node: &Node{1, &Node{2, &Node{3, nil}}},
  }

	if a, e := obj.Get(1), 2; a != e {
		t.Fatalf("Expected %v but got %v", e, a)
	}
}
