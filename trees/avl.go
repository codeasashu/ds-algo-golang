package trees

import (
	"math"
)

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

// Height returns height of a node
func (node *AVLNode) Height() int {
	if node == nil {
		return -1;
	}
	return node.height;
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

	node.height = nodeHeight(*node);

	node = node.balance();

	return node;
}

func (node *AVLNode) balance() *AVLNode {
	if isLeftHeavy(*node) {
		if balanceFactor(*node.leftChild) < 0 {
			node.leftChild = node.leftChild.rotateLeft();
			// node.leftChild.height = nodeHeight(*node.leftChild);
		}
		node = node.rotateRight();
		node.height = nodeHeight(*node);
	} else if isRightHeavy(*node) {
		if balanceFactor(*node.rightChild) > 0 {
			node.rightChild = node.rightChild.rotateRight();
			// node.rightChild.height = nodeHeight(*node.rightChild);
		}
		node = node.rotateLeft();
		// node.height = nodeHeight(*node);
	}
	return node;
}

func (node *AVLNode) rotateLeft() *AVLNode {
	newRoot := node.rightChild;
	if newRoot.leftChild != nil {
		node.rightChild = newRoot.leftChild;
	} else {
		node.rightChild = nil;
	}
	newRoot.leftChild = node;

	node.height = nodeHeight(*node);
	newRoot.height = nodeHeight(*newRoot);

	return newRoot;
}

func (node *AVLNode) rotateRight() *AVLNode {
	newRoot := node.leftChild;
	if newRoot.rightChild != nil {
		node.leftChild = newRoot.rightChild;
	} else {
		node.leftChild = nil;
	}
	newRoot.rightChild = node;

	node.height = nodeHeight(*node);
	newRoot.height = nodeHeight(*newRoot);

	return newRoot;
}

func nodeHeight(node AVLNode) int {
	nodeHeight := math.Max(
		float64(node.leftChild.Height()), float64(node.rightChild.Height()));
	return int(nodeHeight) + 1;
}

func isLeftHeavy(node AVLNode) bool {
	return balanceFactor(node) > 1;
}

func isRightHeavy(node AVLNode) bool {
	return balanceFactor(node) < -1;
}

func balanceFactor(node AVLNode) int {
	return node.leftChild.Height() - node.rightChild.Height();
}

// Insert a value in tree
func (tree *AVLTree) Insert(value int) {
	tree.root = tree.root.Insert(value);
}
