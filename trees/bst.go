package trees

import (
	"errors"
)

// Node is BST tree node
type Node struct {
	value int;
	leftChild *Node;
	rightChild *Node;
}

// Bst tree
type Bst struct {
	root *Node;
}

func isLeafNode(node Node) bool {
	return (node.leftChild == nil && node.rightChild == nil);
}

// Insert a value to a tree node
func (node *Node) Insert(value int) error {
	if node == nil {
		return nil;
	}

	if value > node.value {
		if node.rightChild == nil {
			node.rightChild = &Node{value: value};
		} else {
			node.rightChild.Insert(value);
		}
	} else if value < node.value {
		if node.leftChild == nil {
			node.leftChild = &Node{value: value};
		} else {
			node.leftChild.Insert(value);
		}
	} else {
		return errors.New("Duplicate value");
	}
	return nil;
}

// Insert a value in tree
func (tree *Bst) Insert(value int) {
	if tree.root == nil {
		var node Node = Node{value: value}
		tree.root = &node;
	} else {
		tree.root.Insert(value);
	}
}

// Find a value in a node, recursively
func (node *Node) Find(value int) bool {
	if node == nil {
		return false;
	}
	if node.value == value {
		return true;
	}
	return node.leftChild.Find(value) || node.rightChild.Find(value);
}

// Find a value in tree
func (tree *Bst) Find(value int) bool {
	if tree.root.value == value {
		return true;
	}
	return tree.root.Find(value);
}