package main

import (
	"algos/linkedlist"
	"algos/queues"
	"algos/trees"
	"fmt"
)

func demoQueue() {
	var q queues.Queue;
	q.Init(4);
	q.Enqueue(10);
	q.Enqueue(20);
	q.Enqueue(30);

	fmt.Println("Dequeued", q.Dequeue());
	fmt.Println("first item in queue", q.Peek());
}

func demoLinkedList() {
	var head linkedlist.Node;
	head = linkedlist.Node{};
	linkedlist.Push(&head, 3);
	linkedlist.Push(&head, 5);
	linkedlist.Push(&head, 6);
	linkedlist.PrintList(head);
	deleted := linkedlist.DeleteKey(&head, 5);
	fmt.Println("Is key deleted: ", deleted);
	totalItems := linkedlist.GetCount(&head);
	fmt.Println("total linkedlist nodes: ", totalItems);

	// check if a key exists in linkedlist
	keyExists := linkedlist.Find(head, 6);
	fmt.Println("does 6 exists in LL: ", keyExists);

	// Find using recursive method
	keyExistsRecur := linkedlist.FindRecursive(head, 6);
	fmt.Println("does 6 exists in LL (using recursion): ", keyExistsRecur);
}

func demoBST() {
	var tree trees.Bst;
	tree.Insert(5);
	tree.Insert(15);
	tree.Insert(4);
	tree.Insert(10);
	fmt.Println(tree);
	found := tree.Find(-15);
	fmt.Println(found);
}

func main() {
	// to run Linkedlist algo
	// demoLinkedList();
	
	// to run BST Algo
	// demoBST();

	// to check queue
	// demoQueue();
}