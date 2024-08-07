package main

import (
	"algos/linkedlist"
	"algos/queues"
	"algos/sorting"
	"algos/trees"
	"fmt"
)

func demoQueue() {
	var q queues.Queue
	q.Init(4)
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println("Dequeued", q.Dequeue())
	fmt.Println("first item in queue", q.Peek())
}

func demoLinkedList() {
	var head linkedlist.Node
	head = linkedlist.Node{}
	linkedlist.Push(&head, 3)
	linkedlist.Push(&head, 5)
	linkedlist.Push(&head, 6)
	linkedlist.PrintList(head)
	deleted := linkedlist.DeleteKey(&head, 5)
	fmt.Println("Is key deleted: ", deleted)
	totalItems := linkedlist.GetCount(&head)
	fmt.Println("total linkedlist nodes: ", totalItems)

	// check if a key exists in linkedlist
	keyExists := linkedlist.Find(head, 6)
	fmt.Println("does 6 exists in LL: ", keyExists)

	// Find using recursive method
	keyExistsRecur := linkedlist.FindRecursive(head, 6)
	fmt.Println("does 6 exists in LL (using recursion): ", keyExistsRecur)
}

func demoBST() {
	var tree trees.Bst
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(4)
	tree.Insert(10)
	fmt.Println(tree)
	found := tree.Find(-15)
	fmt.Println(found)

	var otherTree trees.Bst
	otherTree.Insert(5)
	otherTree.Insert(15)
	otherTree.Insert(4)
	otherTree.Insert(10)

	areEqual := tree.Equals(otherTree)
	fmt.Println("Trees are equal?", areEqual)

	tree.Inorder()
}

func demoAVLTree() {
	var tree trees.AVLTree
	tree.Insert(10)
	tree.Insert(30)
	tree.Insert(20)
	fmt.Println(tree)
}

func demoHeapTree() {
	var heap trees.Heap
	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(17)
	heap.Insert(4)
	heap.Insert(22)

	heap.Remove()
}

func demoHeapify(input []int) {
	output := trees.Heapify(input)
	fmt.Println("heapified", output)
}

func main() {
	// to run Linkedlist algo
	// demoLinkedList()

	// to run BST Algo
	// demoBST();

	// to check queue
	// demoQueue();

	// Run AVLTree
	// demoAVLTree();

	// Run Heap tree
	// demoHeapTree();
	// demoHeapify([]int{5, 3, 8, 4});
	sorting.RunBubbleSort()
}
