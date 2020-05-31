package linkedlist

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
	Node, n := this.Node, 0
	for Node != nil {
		if n == index {
			return Node.Val
		}
		Node = Node.Next
	}
	return -1
}

/** Add a Node of value val before the first element of the linked list. After the insertion, the new Node will be the first Node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.Node = &Node{val, this.Node}
}

/** Append a Node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.Node == nil {
		this.AddAtHead(val)
	} else {
		this.Node.Next = &Node{val, nil}
	}
}

/** Add a Node of value val before the index-th Node in the linked list. If index equals to the length of linked list, the Node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	dh, n := &Node{0, this.Node}, 0
	for dh != nil {
		if n == index {
			nextTmp := dh.Next
			dh.Next = &Node{val, nextTmp}
			return
		}
		dh = dh.Next
	}
}

/** Delete the index-th Node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	dh, n := &Node{0, this.Node}, 0
	for dh != nil {
		if n == index {
			if dh.Next != nil {
				dh.Next = dh.Next.Next
			} else {
				dh.Next = nil
			}
		}
		dh = dh.Next
	}

}
