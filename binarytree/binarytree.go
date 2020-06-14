package binarytree

// TreeNode does thing
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codes does thing
type Codec struct {
}

// Constructor does thing
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	return ""
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	return nil
}
