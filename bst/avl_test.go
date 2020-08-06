package bst

import (
	"reflect"
	"testing"
)

func tree() *AVLNode {
	return &AVLNode{
		Val:    41,
		Height: 3,
		Left: &AVLNode{
			Val: 20, Height: 2,
			Left: &AVLNode{
				Val:    11,
				Height: 0,
			},
			Right: &AVLNode{
				Val:    29,
				Height: 1,
				Left: &AVLNode{
					Val:    26,
					Height: 0,
				},
			},
		},
		Right: &AVLNode{
			Val:    65,
			Height: 1,
			Left: &AVLNode{
				Val:    50,
				Height: 0,
			},
		},
	}
}

func TestInsert1(t *testing.T) {
	node := &AVLNode{
		Val:    26,
		Height: 1,
		Left: &AVLNode{
			Val:    23,
			Height: 0,
		},
	}
	actual := insert(node, 29)
	expected := &AVLNode{
		Val:    26,
		Height: 1,
		Left: &AVLNode{
			Val:    23,
			Height: 0,
		},
		Right: &AVLNode{
			Val:    29,
			Height: 0,
		},
	}

	assert(t, expected, actual)
}

func TestInsert2(t *testing.T) {
	node := &AVLNode{
		Val:    26,
		Height: 1,
		Left: &AVLNode{
			Val:    23,
			Height: 0,
		},
	}
	actual := insert(node, 23)
	expected := &AVLNode{
		Val:    26,
		Height: 1,
		Left: &AVLNode{
			Val:    23,
			Height: 0,
		},
	}
	assert(t, expected, actual)
}

// left left case
func TestInsert3(t *testing.T) {
	node := &AVLNode{
		Val:    26,
		Height: 1,
		Left: &AVLNode{
			Val:    23,
			Height: 0,
		},
	}
	actual := insert(node, 20)
	expected := &AVLNode{
		Val:    23,
		Height: 1,
		Left: &AVLNode{
			Val:    20,
			Height: 0,
		},
		Right: &AVLNode{
			Val:    26,
			Height: 0,
		},
	}
	assert(t, expected, actual)
}

// left right case
func TestInsert4(t *testing.T) {
	node := &AVLNode{
		Val:    26,
		Height: 1,
		Left: &AVLNode{
			Val:    23,
			Height: 0,
		},
	}
	actual := insert(node, 24)
	expected := &AVLNode{
		Val:    24,
		Height: 1,
		Left: &AVLNode{
			Val:    23,
			Height: 0,
		},
		Right: &AVLNode{
			Val:    26,
			Height: 0,
		},
	}
	assert(t, expected, actual)
}

// right right
func TestInsert5(t *testing.T) {
	node := &AVLNode{
		Val:    26,
		Height: 1,
		Right: &AVLNode{
			Val:    29,
			Height: 0,
		},
	}
	actual := insert(node, 32)
	expected := &AVLNode{
		Val:    29,
		Height: 1,
		Left: &AVLNode{
			Val:    26,
			Height: 0,
		},
		Right: &AVLNode{
			Val:    32,
			Height: 0,
		},
	}
	assert(t, expected, actual)
}

// right left
func TestInsert6(t *testing.T) {
	node := &AVLNode{
		Val:    26,
		Height: 1,
		Right: &AVLNode{
			Val:    29,
			Height: 0,
		},
	}
	actual := insert(node, 27)
	expected := &AVLNode{
		Val:    27,
		Height: 1,
		Left: &AVLNode{
			Val:    26,
			Height: 0,
		},
		Right: &AVLNode{
			Val:    29,
			Height: 0,
		},
	}
	assert(t, expected, actual)
}

func TestInsert7(t *testing.T) {
	node := NewAVLNode(5)
	node = insert(node, 6)
	node = insert(node, 7)

	expected := &AVLNode{
		Val:    6,
		Height: 1,
		Left: &AVLNode{
			Val:    5,
			Height: 0,
		},
		Right: &AVLNode{
			Val:    7,
			Height: 0,
		},
	}
	assert(t, expected, node)

	node = insert(node, 4)
	node = insert(node, 3)

	expected = &AVLNode{
		Val:    6,
		Height: 2,
		Left: &AVLNode{
			Val:    4,
			Height: 1,
			Left: &AVLNode{
				Val:    3,
				Height: 0,
			},
			Right: &AVLNode{
				Val:    5,
				Height: 0,
			},
		},
		Right: &AVLNode{
			Val:    7,
			Height: 0,
		},
	}
	assert(t, expected, node)

	node = insert(node, 2)

	expected = &AVLNode{
		Val:    4,
		Height: 2,
		Left: &AVLNode{
			Val:    3,
			Height: 1,
			Left: &AVLNode{
				Val:    2,
				Height: 0,
			},
		},
		Right: &AVLNode{
			Val:    6,
			Height: 1,
			Left: &AVLNode{
				Val:    5,
				Height: 0,
			},
			Right: &AVLNode{
				Val:    7,
				Height: 0,
			},
		},
	}
	assert(t, expected, node)

}

func TestRemove1(t *testing.T) {
	node := &AVLNode{
		Val:    26,
		Height: 1,
		Left: &AVLNode{
			Val:    23,
			Height: 0,
		},
		Right: &AVLNode{
			Val:    29,
			Height: 0,
		},
	}
	actual := remove(node, 29)
	expected := &AVLNode{
		Val:    26,
		Height: 1,
		Left: &AVLNode{
			Val:    23,
			Height: 0,
		},
	}
	assert(t, expected, actual)

	actual = remove(node, 23)
	expected = &AVLNode{
		Val:    26,
		Height: 0,
	}
	assert(t, expected, actual)

	actual = remove(node, 26)
	expected = nil
	assert(t, expected, actual)
}

func TestRemove2(t *testing.T) {
	node := &AVLNode{
		Val:    26,
		Height: 1,
		Left: &AVLNode{
			Val:    23,
			Height: 0,
		},
		Right: &AVLNode{
			Val:    29,
			Height: 1,
			Left: &AVLNode{
				Val:    27,
				Height: 0,
			},
		},
	}
	actual := remove(node, 29)
	expected := &AVLNode{
		Val:    26,
		Height: 1,
		Left: &AVLNode{
			Val:    23,
			Height: 0,
		},
		Right: &AVLNode{
			Val:    27,
			Height: 0,
		},
	}
	assert(t, expected, actual)
}

func TestRemove3(t *testing.T) {
	node := &AVLNode{
		Val:    26,
		Height: 1,
		Left: &AVLNode{
			Val:    23,
			Height: 0,
		},
		Right: &AVLNode{
			Val:    29,
			Height: 1,
			Right: &AVLNode{
				Val:    32,
				Height: 0,
			},
		},
	}
	actual := remove(node, 29)
	expected := &AVLNode{
		Val:    26,
		Height: 1,
		Left: &AVLNode{
			Val:    23,
			Height: 0,
		},
		Right: &AVLNode{
			Val:    32,
			Height: 0,
		},
	}
	assert(t, expected, actual)
}

func TestRemove4(t *testing.T) {
	node := &AVLNode{
		Val:    26,
		Height: 2,
		Left: &AVLNode{
			Val:    23,
			Height: 0,
		},
		Right: &AVLNode{
			Val:    29,
			Height: 1,
			Left: &AVLNode{
				Val:    27,
				Height: 0,
			},
			Right: &AVLNode{
				Val:    32,
				Height: 0,
			},
		},
	}
	actual := remove(node, 26)
	expected := &AVLNode{
		Val:    27,
		Height: 2,
		Left: &AVLNode{
			Val:    23,
			Height: 0,
		},
		Right: &AVLNode{
			Val:    29,
			Height: 1,
			Right: &AVLNode{
				Val:    32,
				Height: 0,
			},
		},
	}
	assert(t, expected, actual)
}

func TestRemove5(t *testing.T) {
	node := &AVLNode{
		Val:    26,
		Height: 2,
		Left: &AVLNode{
			Val:    23,
			Height: 1,
			Left: &AVLNode{
				Val:    20,
				Height: 0,
			},
			Right: &AVLNode{
				Val:    24,
				Height: 0,
			},
		},
		Right: &AVLNode{
			Val:    29,
			Height: 0,
		},
	}
	actual := remove(node, 26)
	expected := &AVLNode{
		Val:    24,
		Height: 2,
		Left: &AVLNode{
			Val:    23,
			Height: 1,
			Left: &AVLNode{
				Val:    20,
				Height: 0,
			},
		},
		Right: &AVLNode{
			Val:    29,
			Height: 0,
		},
	}
	assert(t, expected, actual)
}

func TestAscendingSort(t *testing.T) {
	node := &AVLNode{
		Val:    26,
		Height: 2,
		Left: &AVLNode{
			Val:    23,
			Height: 1,
			Left: &AVLNode{
				Val:    20,
				Height: 0,
			},
			Right: &AVLNode{
				Val:    24,
				Height: 0,
			},
		},
		Right: &AVLNode{
			Val:    29,
			Height: 0,
		},
	}
	if e, a := []int{20, 23, 24, 26, 29}, ascendingSort(node); !reflect.DeepEqual(e, a) {
		t.Fatalf("Expected %v but got %v", e, a)
	}
}

func assert(t *testing.T, expected, actual *AVLNode) {
	if e, a := expected, actual; !reflect.DeepEqual(e, a) {
		t.Logf("Exepected %s", e.print())
		t.Logf("Actual %s", a.print())
		t.Fatalf("Expected %v but got %v", e, a)
	}
}
