package doubleLinkedList

import (
	"errors"
	"fmt"
)

type node struct {
	data     int
	nextNode *node
	previous *node
}

type LinkedList struct {
	head *node
	tail *node
}

func NewLinkedList() *LinkedList {
	head := &node{}
	tail := &node{}

	head.nextNode = tail
	tail.previous = head

	return &LinkedList{
		head: head,
		tail: tail,
	}
}

func (ll *LinkedList) IterateFromTail() {
	current := ll.tail
	for current.previous != nil {
		current = current.previous
		if current == ll.head {
			break
		}
		if current.previous == ll.head {
			fmt.Println(current)
			break
		}
		fmt.Print(current, " ==> ")
	}
}

func (ll *LinkedList) IterateFromHead() {
	current := ll.head
	for current.nextNode != nil {
		current = current.nextNode
		if current == ll.tail {
			break
		}
		if current.nextNode == ll.tail {
			fmt.Println(current)
			break
		}
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
	newNode := &node{
		data:     data,
		previous: ll.head,
		nextNode: nodeAfterHead,
	}

	ll.head.nextNode = newNode
	nodeAfterHead.previous = newNode
}

func (ll *LinkedList) AddAtEnd(data int) {
	nodeBeforeTail := ll.tail.previous
	newNode := &node{
		data: data,
	}

	newNode.nextNode = ll.tail
	newNode.previous = nodeBeforeTail
	ll.tail.previous = newNode
	nodeBeforeTail.nextNode = newNode
}

func (ll *LinkedList) InsertNode(data int, nodeBeforeInsertion *node) {
	newNode := &node{
		data: data,
	}

	// update references to next node
	newNode.nextNode = nodeBeforeInsertion.nextNode
	nodeBeforeInsertion.nextNode = newNode

	// updated references to previous node
	newNode.nextNode.previous = newNode
	newNode.previous = nodeBeforeInsertion
}

func (ll *LinkedList) DeleteAfter(nodeBeforeDelete *node) {
	nodeToBeDeleted := nodeBeforeDelete.nextNode
	nodeAfterDeleted := nodeToBeDeleted.nextNode

	nodeBeforeDelete.nextNode = nodeAfterDeleted
	nodeAfterDeleted.previous = nodeBeforeDelete
	nodeToBeDeleted = nil
}

func (ll *LinkedList) CopyList() *LinkedList {
	llCopy := new(LinkedList)
	head := new(node)
	tail := new(node)

	head.data = 0
	head.nextNode = tail
	head.previous = nil

	tail.data = 0
	tail.previous = head
	tail.nextNode = nil

	llCopy.head = head
	llCopy.tail = tail

	current := ll.head.nextNode
	for current.nextNode != nil {
		llCopy.AddAtEnd(current.data)
		current = current.nextNode
	}

	return llCopy
}

func (ll *LinkedList) InsertSort() {
	input := ll.head.nextNode

	for input != ll.tail {
		nextNode := input
		input = input.nextNode

		afterMe := ll.head
		for afterMe != ll.tail && afterMe.nextNode.data < nextNode.data {
			afterMe = afterMe.nextNode
		}

		ll.InsertNode(nextNode.data, afterMe)
		ll.DeleteAfter(nextNode.previous)
	}
}

func (ll *LinkedList) SelectionSort() *LinkedList {

	sortedList := NewLinkedList()
	input := ll.head
	for input.nextNode != ll.tail {
		bestAfterMe := input
		bestValue := bestAfterMe.nextNode.data

		afterMe := input.nextNode
		for afterMe.nextNode != ll.tail {
			if afterMe.nextNode.data > bestValue {
				bestAfterMe = afterMe
				bestValue = afterMe.nextNode.data
			}
			afterMe = afterMe.nextNode
		}

		bestNode := bestAfterMe.nextNode
		ll.DeleteAfter(bestAfterMe)
		sortedList.AddAtBegginning(bestNode.data)

	}

	return sortedList
}

func (ll *LinkedList) FindMax() *node {

	input := ll.head
	bestNode := input.nextNode
	bestValue := input.nextNode.data
	for input != ll.tail {

		if input.data > bestValue {
			bestValue = input.data
			bestNode = input
		}
		input = input.nextNode
	}

	return bestNode
}

func (ll *LinkedList) FindNodeBeforeMax() *node {
	maxNodePrevious := ll.FindMax().previous
	if maxNodePrevious == ll.head {
		return nil
	}

	return maxNodePrevious
}

// Slice returns linked list from n to end
func (ll *LinkedList) Slice(n int) *LinkedList {
	slicedList := NewLinkedList()
	slicedList.tail = ll.tail

	counter := 0
	input := ll.head.nextNode
	for input != ll.tail {
		if counter == n+1 {
			slicedList.head.nextNode = input
			input.previous = slicedList.head
			return slicedList
		}
		input = input.nextNode
		counter++
	}

	return nil
}

func (ll *LinkedList) Head() *node {
	return ll.head
}
func (ll *LinkedList) Tail() *node {
	return ll.tail
}

func (n *node) NextNode() (*node, error) {
	if n.nextNode == nil {
		return nil, errors.New("there is no next node")
	}
	return n.nextNode, nil
}

func (n *node) Data() int {
	return n.data
}
