package binarytree

import (
	"testing"
)

func TestSerialize(t *testing.T) {

	c := Constructor()
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	if a, e := c.serialize(node), "1,2,3,4,5,nil,7"; a != e {
		t.Fatalf("Expected %v but got %v", e, a)
	}
}
