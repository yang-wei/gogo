package linkedlist

import (
	"strconv"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Node *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		Node: nil,
	}
}

/** Get the value of the index-th Node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	n := 0
	for node := this.Node; node != nil; node = node.Next {
		if n == index {
			return node.Val
		}
		n++
	}
	return -1

}

/** Add a Node of value val before the first element of the linked list. After the insertion, the new Node will be the first Node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.Node = &Node{val, this.Node}
}

/** Append a Node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	dh := &Node{0, this.Node}
	node := dh
	for node.Next != nil {
		node = node.Next
	}
	node.Next = &Node{val, nil}
	this.Node = dh.Next
}

/** Add a Node of value val before the index-th Node in the linked list. If index equals to the length of linked list, the Node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	dh, n := &Node{0, this.Node}, 0
	for node := dh; node != nil; node = node.Next {
		if n == index {
			node.Next = &Node{val, node.Next}
			break
		}
		n++
	}
	this.Node = dh.Next
}

/** Delete the index-th Node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	dh, n := &Node{0, this.Node}, 0
	for node := dh; node != nil; node = node.Next {
		if n == index {
			if node.Next != nil {
				node.Next = node.Next.Next
			} else {
				node.Next = nil
			}
			break
		}
		n++
	}

	this.Node = dh.Next

}

func (this *MyLinkedList) toString() string {
	l, node := make([]string, 0), this.Node
	for node != nil {
		l = append(l, strconv.Itoa(node.Val))
		node = node.Next
	}
	return strings.Join(l, "->")
}
