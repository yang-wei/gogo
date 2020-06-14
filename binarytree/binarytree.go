package binarytree

import (
	"strconv"
	"strings"
)

// TreeNode does thing
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	splitter = ","
	empty    = "X"
)

// Codec does thing
type Codec struct {
}

// Constructor does thing
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	r := _serialize(root)
	return strings.Join(r, splitter)
}

// build a DFS preorder
func _serialize(root *TreeNode) []string {
	r := []string{}

	if root == nil {
		return []string{empty}
	}

	r = append(r, strconv.Itoa(root.Val))
	r = append(r, _serialize(root.Left)...)
	r = append(r, _serialize(root.Right)...)
	return r
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	l := strings.Split(data, splitter)
	return _deserialize(&l)
}

func _deserialize(l *[]string) *TreeNode {

	copied := *l
	if len(copied) == 0 {
		return nil
	}

	val, rest := copied[0], copied[1:]
	*l = rest
	if val == empty {
		return nil
	}

	v, _ := strconv.Atoi(val)

	tree := &TreeNode{
		Val: v,
	}
	tree.Left = _deserialize(l)
	tree.Right = _deserialize(l)

	return tree
}
