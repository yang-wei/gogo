package binarytree

import (
	"reflect"
	"testing"
)

func TestSerialize(t *testing.T) {

	c := Constructor()
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	if a, e := c.serialize(node), "1,2,X,6,X,X,3,4,X,X,5,X,X"; a != e {
		t.Fatalf("Expected %v but got %v", e, a)
	}

	if a, e := c.deserialize(c.serialize(node)), node; !reflect.DeepEqual(a, e) {
		t.Fatalf("Expected %v but got %v", e, a)
	}
}
