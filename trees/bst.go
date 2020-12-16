package trees

import (
	"errors"
	"fmt"
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

// Equals compares two nodes recursively and checks if they are equal
func (node *Node) Equals(other *Node) bool {
	if node == nil && other == nil {
		return true;
	}

	if node.value != other.value {
		return false;
	}

	if isLeafNode(*node) == true {
		return isLeafNode(*other) == true;
	}

	if node.leftChild == nil {
		return other.leftChild == nil;
	}

	if node.rightChild == nil {
		return other.rightChild == nil;
	}

	return (
		node.leftChild.Equals(other.leftChild) &&
		node.rightChild.Equals(other.rightChild));
}

// Equals compares equality of two trees
func (tree Bst) Equals(other Bst) bool {
	if tree.root == nil && other.root == nil {
		return true;
	}

	return tree.root.Equals(other.root);
}

func inorder(node *Node) {
	if node == nil {
		return;
	}

	inorder(node.leftChild); // Visit Left first
	fmt.Printf("%d ", node.value); // Vist root
	inorder(node.rightChild); // Visit right child
}

// Inorder BST traversal
func (tree Bst) Inorder() {
	fmt.Print("Inorder:")
	inorder(tree.root);
	fmt.Print("\n")
}