package linkedlist

import (
	"fmt"
)

// Node is linkedlist node
type Node struct {
	value int;
	next *Node;
}

// PrintList prints the linkedlist on std console
func PrintList(listHead Node) {
	for &listHead != nil {
		fmt.Println("Current list value", listHead.value)
		if listHead.next == nil {
			break;
		}
		listHead = *listHead.next;
	}
} 

// Push adds items to linkedlist
func Push(head *Node, newData int) {
	temp := *head;
	*head = Node{value: newData, next: &temp};
}

// PushAfter is similar to push, adds item after a node
func PushAfter(someNode *Node, newData int) {
	// Get addr of node after
	// temp := someNode;
	var newNode Node = Node{value: newData};
	var nextNode Node = *someNode.next;

	// If this is tail
	if someNode.next == nil {
		someNode.next = &newNode;
		*someNode.next = newNode;
	} else {
		newNode.next = &nextNode;
		*someNode.next = newNode;
	}
}

// DeleteKey deletes a key from linkedlist Node
func DeleteKey(listHead *Node, key int) bool {
	var tmp *Node = listHead;
	var prev *Node;

	// head holds the key
	if tmp != nil && tmp.value == key {
		*listHead = *tmp.next;
		return true;
	}

	// search until we have a match
	for &tmp != nil && tmp.value != key {
		prev = tmp;
		tmp = tmp.next;
	}

	// We have reach the end
	if tmp == nil {
		return false;
	}

	// Unlink the node from linked list 
	*prev.next = *tmp.next;

	return true;
}

// DeleteList deletes the entire linkedlist
func DeleteList(head *Node) {
	*head = Node{};
}

// GetCount gets count of Nodes in a linkedlist
func GetCount(head *Node) int {
	if head == nil {
		return 0;
	}

	return 1 + GetCount(head.next);
}

// Find checks if key exists in a linkedlist
func Find(head Node, key int) bool {
	var found bool = false;
	for &head != nil {
		if head.value == key {
			found = true;
			break;
		}
		if head.next == nil {
			break;
		}
		head = *head.next;
	}
	return found;
}

// FindRecursive is similar to Find, but uses recursion
func FindRecursive(head Node, key int) bool {
	if head.value  == key {
		return true;
	}

	if head.next == nil {
		return false;
	}

	return FindRecursive(*head.next, key);
}
