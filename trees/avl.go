package trees

import "math"

// AVLNode is AVL tree node
type AVLNode struct {
	value int;
	height int;
	leftChild *AVLNode;
	rightChild *AVLNode;
}

// AVLTree tree
type AVLTree struct {
	root *AVLNode;
}

func isAVLLeafNode(node AVLNode) bool {
	return (node.leftChild == nil && node.rightChild == nil);
}

func (node *AVLNode) Height() int {
	if node == nil {
		return 0;
	}
	if isAVLLeafNode(*node) {
		return 0;
	}
	leftHeight := node.leftChild.Height();
	rightHeight := node.rightChild.Height();
	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1;
}

// Insert a value to a tree node
func (node *AVLNode) Insert(value int) *AVLNode {
	if node == nil {
		return &AVLNode{value: value};
	}

	if value > node.value {
		node.rightChild = node.rightChild.Insert(value);
	} else {
		node.leftChild = node.leftChild.Insert(value);
	}
	node.height = node.Height();
	return node;
}

// Insert a value in tree
func (tree *AVLTree) Insert(value int) {
	tree.root = tree.root.Insert(value);
}
