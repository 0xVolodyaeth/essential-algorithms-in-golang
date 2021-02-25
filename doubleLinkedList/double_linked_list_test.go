package doubleLinkedList

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddAtBegginining(t *testing.T) {
	ll := NewLinkedList()
	ll.AddAtBegginning(2)

	insertedNode := ll.head.nextNode

	// Check head is connected to inserted node
	require.Equal(t, ll.head.nextNode, insertedNode)
	require.Equal(t, ll.head, insertedNode.previous)

	// Check tail is connected to inserted node
	require.Equal(t, ll.tail.previous, insertedNode)
	require.Equal(t, ll.tail, insertedNode.nextNode)
}

func TestDeleteAfter(t *testing.T) {
	ll := NewLinkedList()
	ll.AddAtBegginning(2)

	head := ll.head

	ll.DeleteAfter(head)
	require.Equal(t, ll.head.nextNode, ll.tail)
}

func TestAddAtEnd(t *testing.T) {
	ll := NewLinkedList()
	ll.AddAtEnd(2)
	require.Equal(t, ll.head.nextNode.data, 2)
}

func TestInsertNode(t *testing.T) {
	ll := NewLinkedList()
	ll.AddAtBegginning(2)
	require.Equal(t, ll.head.nextNode.data, 2)

	ll.InsertNode(7, ll.head.nextNode)
	require.Equal(t, ll.head.nextNode.nextNode.data, 7)
}

func TestFindNode(t *testing.T) {
	ll := NewLinkedList()
	ll.AddAtBegginning(2)
	ll.AddAtBegginning(4)

	node, err := ll.FindNode(2)
	require.NoError(t, err)
	require.Equal(t, ll.tail.previous, node)

	node, err = ll.FindNode(10)
	require.Error(t, err)
	require.Nil(t, node)
}

func TestFindNodeBefore(t *testing.T) {
	ll := NewLinkedList()
	ll.AddAtBegginning(2)
	ll.AddAtBegginning(4)

	node, err := ll.FindNodeBefore(2)
	require.NoError(t, err)
	require.Equal(t, node, ll.head.nextNode)

	node, err = ll.FindNodeBefore(10)
	require.Error(t, err)
	require.Nil(t, node)
}

func TestCopyList(t *testing.T) {
	ll := NewLinkedList()
	ll.AddAtBegginning(2)
	ll.AddAtBegginning(4)
	ll.AddAtBegginning(6)
	ll.AddAtBegginning(8)

	llCopy := ll.CopyList()
	input := llCopy.head.nextNode
	for input != llCopy.tail {
		node, err := ll.FindNode(input.data)
		require.NoError(t, err)
		require.Equal(t, node.data, input.data)
		input = input.nextNode
	}
}

func TestInserionSort(t *testing.T) {
	ll := NewLinkedList()
	ll.AddAtBegginning(1)
	ll.AddAtBegginning(3)
	ll.AddAtBegginning(4)
	ll.AddAtBegginning(5)
	ll.AddAtBegginning(2)

	ll.InsertSort()
	input := ll.head.nextNode
	for i := 1; i <= 5; i++ {
		require.Equal(t, input.data, i)
		input = input.nextNode
	}
}

func TestFindMax(t *testing.T) {
	ll := NewLinkedList()
	ll.AddAtBegginning(2)
	ll.AddAtBegginning(4)
	ll.AddAtBegginning(1)

	max := ll.FindMax()
	require.Equal(t, max.data, 4)
}

func TestFindNodeBeforeMax(t *testing.T) {
	ll := NewLinkedList()
	ll.AddAtBegginning(2)
	ll.AddAtBegginning(4)
	ll.AddAtBegginning(1)

	beforeMax := ll.FindNodeBeforeMax()
	require.Equal(t, beforeMax.data, 1)
}

func TestSlice(t *testing.T) {
	ll := NewLinkedList()
	ll.AddAtBegginning(1)
	ll.AddAtBegginning(2)
	ll.AddAtBegginning(3)
	ll.AddAtBegginning(4)

	slicedLl := ll.Slice(1)
	require.Equal(t, slicedLl.head.nextNode.data, ll.head.nextNode.nextNode.nextNode.data)
	require.Equal(t, slicedLl.head.nextNode.nextNode.data, ll.head.nextNode.nextNode.nextNode.nextNode.data)
}

func TestSelectionSort(t *testing.T) {
	ll := NewLinkedList()
	ll.AddAtBegginning(1)
	ll.AddAtBegginning(3)
	ll.AddAtBegginning(4)
	ll.AddAtBegginning(5)
	ll.AddAtBegginning(2)

	sorted := ll.SelectionSort()
	input := sorted.head.nextNode
	for i := 1; i <= 5; i++ {
		require.Equal(t, input.data, i)
		input = input.nextNode
	}
}
