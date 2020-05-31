type Node struct {
    Val int
    Next *Node
}

type MyLinkedList struct {
    node *Node
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    return MyLinkedList{
        node: nil,
    }
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    node, n := this.node, 0
    for node != nil {
        if n == index {
            return node.Val
        }
        node = node.Next
    }
    return -1
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
    this.node = &Node{val, this.node}
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    if this.node == nil {
        this.AddAtHead(val)
    } else {
        this.node.Next = &Node{val, nil}
    }
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    dh, n := &Node{0, this.node}, 0
    for dh != nil {
        if n == index {
            nextTmp := dh.Next
            dh.Next = &Node{val, nextTmp}
            return
        }
        dh = dh.Next
    }
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    dh, n := &Node{0, this.node}, 0
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
