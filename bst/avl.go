package bst

import (
	"fmt"
	"strings"
)

type AVLNode struct {
	Val    int
	Height int
	Left   *AVLNode
	Right  *AVLNode
}

func NewAVLNode(val int) *AVLNode {
	return &AVLNode{
		Val:    val,
		Height: 0,
		Left:   nil,
		Right:  nil,
	}
}

func insert(root *AVLNode, val int) *AVLNode {
	if root == nil {
		return NewAVLNode(val)
	}
	if val > root.Val {
		root.Right = insert(root.Right, val)
	} else if root.Val > val {
		root.Left = insert(root.Left, val)
	} else {
		return root
	}
	updateHeight(root)
	return balance(root)
}

func remove(root *AVLNode, val int) *AVLNode {
	if root == nil {
		return root
	}
	if val > root.Val {
		// on the right
		root.Right = remove(root.Right, val)
	} else if val < root.Val {
		// on the left
		root.Left = remove(root.Left, val)
	} else {
		// on the root
		// #1 leave
		if root.Left == nil && root.Right == nil {
			return nil
			// #2 has left child only
		} else if root.Right == nil {
			return root.Left
			// #3 has right child only
		} else if root.Left == nil {
			return root.Right
		} else {
			// #4 root
			if getHeight(root.Right) > getHeight(root.Left) {
				// let smallest node on right side be the root
				min := findMin(root.Right)
				root.Val = min
				root.Right = remove(root.Right, min)
			} else {
				// let largest node on left side be the root
				max := findMax(root.Left)
				root.Val = max
				root.Left = remove(root.Left, max)
			}
		}
	}
	updateHeight(root)
	return balance(root)
}

func findMin(root *AVLNode) int {
	if root.Left == nil {
		return root.Val
	}
	return findMin(root.Left)
}

func findMax(root *AVLNode) int {
	if root.Right == nil {
		return root.Val
	}
	return findMax(root.Right)
}

func balance(root *AVLNode) *AVLNode {
	left, right := getHeight(root.Left), getHeight(root.Right)
	if left-right > 1 {
		// left right
		if getHeight(root.Left.Right) > getHeight(root.Left.Left) {
			root.Left = leftRotate(root.Left)
		}
		// left left
		return rightRotate(root)
	} else if right-left > 1 {
		// right left
		if getHeight(root.Right.Left) > getHeight(root.Right.Right) {
			root.Right = rightRotate(root.Right)
		}
		// right right
		return leftRotate(root)
	}
	return root
}

func rightRotate(root *AVLNode) *AVLNode {
	next := root.Left
	root.Left = next.Right
	next.Right = root
	updateHeight(next.Right)
	updateHeight(next)
	return next
}

func leftRotate(root *AVLNode) *AVLNode {
	next := root.Right
	root.Right = next.Left
	next.Left = root
	updateHeight(next.Left)
	updateHeight(next)
	return next
}

func getHeight(root *AVLNode) int {
	if root == nil {
		return -1
	}
	return root.Height
}

func updateHeight(root *AVLNode) {
	l, r := -1, -1
	if root.Left != nil {
		l = root.Left.Height
	}
	if root.Right != nil {
		r = root.Right.Height
	}
	if l > r {
		root.Height = l + 1
	} else {
		root.Height = r + 1
	}
}

func ascendingSort(root *AVLNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	result = append(result, ascendingSort(root.Left)...)
	result = append(result, root.Val)
	result = append(result, ascendingSort(root.Right)...)
	return result
}

func height(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func (node *AVLNode) print() string {

	queue := make([]*AVLNode, 0)
	queue = append(queue, node)

	rows := make([][]string, 0)

	for i := 1; len(queue) > 0; i = 2 * i {
		row := make([]string, i)
		ele := queue
		queue = nil
		// dequeue
		for i, e := range ele {
			if e == nil {
				row[i] = "x"
				continue
			} else {
				row[i] = fmt.Sprintf("%d(%d)", e.Val, e.Height)
			}
			queue = append(queue, e.Left)
			queue = append(queue, e.Right)
		}
		rows = append(rows, row)
	}

	ss := make([]string, 0)
	ss = append(ss, "")
	for _, row := range rows {
		ss = append(ss, strings.Join(row, " "))
	}
	ss = ss[:len(ss)-1]

	// indent row
	for i := len(ss) - 1; i >= 0; i-- {
		spaceToIndent := strings.Repeat("  ", len(ss)-i)
		ss[i] = fmt.Sprintf("%s%s", spaceToIndent, ss[i])
	}

	return strings.Join(ss, "\n")
}
