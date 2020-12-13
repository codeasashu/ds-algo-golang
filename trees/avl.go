package trees

// AVLNode is AVL tree node
type AVLNode struct {
	value int;
	leftChild *AVLNode;
	rightChild *AVLNode;
}

// AVLTree tree
type AVLTree struct {
	root *AVLNode;
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
	return node;
}

// Insert a value in tree
func (tree *AVLTree) Insert(value int) {
	tree.root = tree.root.Insert(value);
}
