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
	obj.AddAtTail(3)

	if a, e := obj.toString(), "1->2->3"; a != e {
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

	obj.AddAtIndex(0, 0)
	if a, e := obj.toString(), "0->1->2->3"; a != e {
		t.Fatalf("Expected %v but got %v", e, a)
	}

	obj.AddAtIndex(999, 999)
	if a, e := obj.toString(), "0->1->2->3"; a != e {
		t.Fatalf("Expected %v but got %v", e, a)
	}
}

func TestAddAtIndex2(t *testing.T) {

	obj := new(MyLinkedList)
	obj.AddAtIndex(0, 10)
	obj.AddAtIndex(0, 20)
	obj.AddAtIndex(1, 30)

	if a, e := obj.Get(0), 20; a != e {
		t.Fatalf("Expected %v but got %v", e, a)
	}

	if a, e := obj.toString(), "20->30->10"; a != e {
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

func TestDeleteAtIndex2(t *testing.T) {
	obj := MyLinkedList{
		Node: &Node{1, &Node{2, &Node{3, nil}}},
	}
	obj.DeleteAtIndex(5)
	if a, e := obj.toString(), "1->2->3"; a != e {
		t.Fatalf("Expected %v but got %v", e, a)
	}
	obj.DeleteAtIndex(0)
	if a, e := obj.toString(), "2->3"; a != e {
		t.Fatalf("Expected %v but got %v", e, a)
	}
	obj.DeleteAtIndex(999)
	if a, e := obj.toString(), "2->3"; a != e {
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

func TestAll(t *testing.T) {
	obj := new(MyLinkedList)

	obj.AddAtHead(1)
	obj.AddAtHead(3)
	obj.AddAtIndex(1, 2)
	t.Logf("Node: %v", obj.toString())
	if a, e := obj.Get(1), 2; a != e {
		t.Fatalf("Expected %v but got %v. (Node: %v)", e, a, obj.toString())
	}
	obj.DeleteAtIndex(0)
	t.Logf("Node: %v", obj.toString())
	if a, e := obj.Get(0), 2; a != e {
		t.Fatalf("Expected %v but got %v. (Node: %v)", e, a, obj.toString())
	}

}
