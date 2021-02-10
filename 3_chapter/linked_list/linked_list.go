package main

import (
	"errors"
	"fmt"
)

type node struct {
	data     int
	nextNode *node
}

type LinkedList struct {
	head *node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: &node{},
	}
}

func (ll *LinkedList) Iterate() {
	current := ll.head
	for current.nextNode != nil {
		current = current.nextNode
		fmt.Print(current, " ==> ")
	}
}

func (ll *LinkedList) FindNode(data int) (*node, error) {
	current := ll.head
	for current.nextNode != nil {
		if current.data == data {
			return current, nil
		}
		current = current.nextNode
	}
	return nil, errors.New("no data in list")
}

func (ll *LinkedList) FindNodeBefore(data int) (*node, error) {
	current := ll.head
	for current.nextNode != nil {
		if current.nextNode.data == data {
			return current, nil
		}
		current = current.nextNode
	}
	return nil, errors.New("no data in list")
}

func (ll *LinkedList) AddAtBegginning(data int) {
	nodeAfterHead := ll.head.nextNode
	ll.head.nextNode = &node{data, nodeAfterHead}
}

func (ll *LinkedList) AddAtEnd(data int) {
	current := ll.head
	for current.nextNode != nil {
		current = current.nextNode
	}
	current.nextNode = &node{data, nil}
}

func (ll *LinkedList) InsertNode(data int, nodeBeforeInsertion *node) {
	newNode := &node{data, nodeBeforeInsertion.nextNode}
	nodeBeforeInsertion.nextNode = newNode
}

func (ll *LinkedList) DeleteAfter(nodeBeforeDelete *node) {
	nodeBeforeDelete.nextNode = nodeBeforeDelete.nextNode.nextNode
}

func main() {
	fmt.Println()
	ll := NewLinkedList()

	fmt.Println("=== AddAtBegginning ===")
	ll.AddAtBegginning(1)
	ll.AddAtBegginning(2)
	ll.AddAtBegginning(3)
	ll.Iterate()
	fmt.Println()
	fmt.Println()

	fmt.Println("=== AddAtEnd ===")
	ll.AddAtEnd(10)
	ll.Iterate()
	fmt.Println()
	fmt.Println()

	fmt.Println("=== FindNode ===")
	node, _ := ll.FindNode(1)
	fmt.Println(node)
	fmt.Println()
	fmt.Println()

	fmt.Println("=== FindNodeBefore ===")
	node, _ = ll.FindNodeBefore(1)
	fmt.Println(node)
	fmt.Println()
	fmt.Println()

	fmt.Println("=== InsertNode ===")
	nodeBeforeInsertion, _ := ll.FindNode(2)
	ll.InsertNode(100, nodeBeforeInsertion)
	ll.Iterate()
	fmt.Println()
	fmt.Println()

	fmt.Println("=== Delete ===")
	nodeBeforeInsertion, _ = ll.FindNode(2)
	ll.DeleteAfter(nodeBeforeInsertion)
	ll.Iterate()
	fmt.Println()
	fmt.Println()

	fmt.Println()
}
