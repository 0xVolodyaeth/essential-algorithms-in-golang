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

func (ll *LinkedList) InsertSortWithoutCopy() {
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

func (ll *LinkedList) InsertionSort() *LinkedList {

	sorted := NewLinkedList()
	input := ll.head.nextNode

	for input != ll.tail {
		nextNode := input
		input = input.nextNode

		afterMe := sorted.head
		for afterMe.nextNode != sorted.tail && afterMe.nextNode.data < nextNode.data {
			afterMe = afterMe.nextNode
		}
		sorted.InsertNode(nextNode.data, afterMe)
	}

	return sorted
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
		if counter == n {
			slicedList.head.nextNode = input
			input.previous = slicedList.head
			return slicedList
		}
		input = input.nextNode
		counter++
	}

	return nil
}

// func main() {
// 	fmt.Println()
// 	ll := NewLinkedList()

// 	fmt.Println("=== AddAtBegginning ===")
// 	ll.AddAtBegginning(1)
// 	ll.AddAtBegginning(2)
// 	ll.AddAtBegginning(3)
// 	ll.IterateFromHead()
// 	fmt.Println()
// 	fmt.Println()

// 	ll.AddAtEnd(2)
// 	fmt.Println("=== AddAtEnd ===")
// 	ll.AddAtEnd(10)
// 	ll.IterateFromHead()
// 	fmt.Println()
// 	fmt.Println()

// 	fmt.Println("=== FindNode ===")
// 	node, _ := ll.FindNode(1)
// 	fmt.Println(node)
// 	fmt.Println()
// 	fmt.Println()

// 	fmt.Println("=== FindNodeBefore ===")
// 	node, _ = ll.FindNodeBefore(1)
// 	fmt.Println(node)
// 	fmt.Println()
// 	fmt.Println()

// 	fmt.Println("=== InsertNode ===")
// 	nodeBeforeInsertion, _ := ll.FindNode(2)
// 	ll.InsertNode(100, nodeBeforeInsertion)
// 	ll.IterateFromHead()
// 	fmt.Println()
// 	fmt.Println()

// 	fmt.Println("=== Delete ===")
// 	nodeBeforeInsertion, _ = ll.FindNode(2)
// 	ll.DeleteAfter(nodeBeforeInsertion)
// 	ll.IterateFromHead()
// 	fmt.Println()
// 	fmt.Println()

// 	fmt.Println("=== IterateFromHead ===")
// 	ll.IterateFromHead()
// 	fmt.Println()
// 	fmt.Println()

// 	fmt.Println("=== IterateFromTail ===")
// 	ll.IterateFromTail()
// 	fmt.Println()

// 	fmt.Println("=== CopyList ===")

// 	ll.IterateFromHead()
// 	fmt.Println()
// 	llCopy := ll.CopyList()
// 	llCopy.IterateFromHead()
// 	fmt.Println()

// 	fmt.Println()
// 	fmt.Println("=== InsertionSort with copying===")
// 	sorted := llCopy.InsertionSort()
// 	sorted.IterateFromHead()
// 	sorted.IterateFromTail()

// 	fmt.Println()
// 	fmt.Println("=== InsertionSort without copying ===")
// 	llCopy.IterateFromHead()
// 	llCopy.InsertSortWithoutCopy()
// 	llCopy.IterateFromHead()

// 	fmt.Println()
// 	fmt.Println("=== SelectionSort ===")
// 	ll.IterateFromHead()
// 	sortedLl := ll.SelectionSort()
// 	sortedLl.IterateFromHead()

// 	fmt.Println()
// 	fmt.Println("=== Slice Linked List ===")
// 	ll = NewLinkedList()
// 	ll.AddAtBegginning(1)
// 	ll.AddAtBegginning(2)
// 	ll.AddAtBegginning(3)
// 	ll.AddAtBegginning(4)
// 	ll.IterateFromHead()
// 	llSliced := ll.Slice(1)
// 	llSliced.IterateFromHead()
// }
