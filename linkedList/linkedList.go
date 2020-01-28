package linkedList

import (
	"errors"
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	Size int
	Head *Node
}

// NewList - To create new linked list
func NewList() *LinkedList {
	return &LinkedList{}
}

// CreateNode - To create new node
func (list *LinkedList) CreateNode(nodeValue interface{}) *Node {
	return &Node{Value: nodeValue}
}

// Insert - To insert new node in the linked list
func (list *LinkedList) Insert(newNode *Node) {

	if list.Size == 0 {
		list.Head = newNode
	} else {
		currentNode := list.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = newNode
	}

	list.Size++

}

/*

InsertAfter - Insert new node after provided value in second parameter as targetNode
param 1 - newNode     - node which you want insert
param 2 - targetNode - after this node, newNode will be inserted

steps:
	NewNode.next −> RightNode;
	LeftNode.next −> NewNode;

*/

func (list *LinkedList) InsertAfter(newNode, targetNode *Node) error {

	if list.Size == 0 {
		return errors.New("list is empty")
	}

	if targetNode == nil {
		return errors.New("targetNode cannot be nil")
	}

	//iterate link list from head
	//break loop when target node is found

	currentNode := list.Head
	isTargetNodeFound := false

	for currentNode.Next != nil {
		if currentNode.Value == targetNode.Value {
			//targetNode found in the link list
			currentNode = currentNode
			isTargetNodeFound = true
			break
		}
		currentNode = currentNode.Next
	}

	if isTargetNodeFound {
		if currentNode.Next == nil {
			// this is last node
			currentNode.Next = newNode
			list.Size++
		} else {
			//this is right node
			rightNode := currentNode.Next
			//newNode in the middle
			newNode.Next = rightNode
			//this is left node
			currentNode.Next = newNode
			list.Size++
		}
	} else {
		//target node value not found in the link list
		return errors.New("target node not found")
	}

	return nil

}

/*

Remove -  Remove target node from link list

steps:
	LeftNode.next −> TargetNode.next
	TargetNode.next −> nil;

*/
func (list *LinkedList) Remove(targetNode *Node) error {

	if list.Size == 0 {
		return errors.New("list is empty")
	}

	if list.Head.Value == targetNode.Value {
		list.Head = list.Head.Next
		return nil
	}

	currentNode := list.Head

	for currentNode.Next != nil {
		if currentNode.Next.Value == targetNode.Value {
			currentNode.Next = currentNode.Next.Next
			break
		}

		currentNode = currentNode.Next
	}

	return nil
}

/* Display - To print the values of linked list */
func (list *LinkedList) Display() {
	node := list.Head
	for node != nil {
		fmt.Println("VALUE:", node.Value)
		node = node.Next
	}
}
