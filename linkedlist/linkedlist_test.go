package linkedlist

import (
  "testing"
)

func TestAddAtHead(t *testing.T) {
  obj := new(MyLinkedList)
  obj.AddAtHead(1)
  obj.AddAtHead(2)

  expected := &Node{2, &Node{1, nil}}

  if a, e := obj.Node, expected; a != e {
    t.Fatalf("Expected %v but got %v", e, a)
  }

}
