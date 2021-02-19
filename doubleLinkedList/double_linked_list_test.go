package doubleLinkedList

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type LinkedListTestSuite struct {
	ll *LinkedList
	suite.Suite
}

func TestLinkedListSuite(t *testing.T) {
	ll := NewLinkedList()

	llts := &LinkedListTestSuite{
		ll: ll,
	}

	require.NotNil(t, ll)
	suite.Run(t, llts)
}

func (s *LinkedListTestSuite) TestAddAtBegginining() {
	s.ll.AddAtBegginning(2)
	s.Equal(s.ll.head.nextNode.data, 2)
}

func (s *LinkedListTestSuite) TestDeleteAfter() {
	s.ll.DeleteAfter(s.ll.head)
	s.Equal(s.ll.head.nextNode, s.ll.tail)
}

func (s *LinkedListTestSuite) TestAddAtEnd() {
	ll := NewLinkedList()
	ll.AddAtEnd(2)
	s.Equal(ll.head.nextNode.data, 2)
}

func (s *LinkedListTestSuite) TestInsertNode() {
	ll := NewLinkedList()
	ll.AddAtBegginning(2)
	s.Equal(ll.head.nextNode.data, 2)

	ll.InsertNode(7, ll.head.nextNode)
	s.Equal(ll.head.nextNode.nextNode.data, 7)
}

func (s *LinkedListTestSuite) TestFindNode() {
	ll := NewLinkedList()
	ll.AddAtBegginning(2)
	ll.AddAtBegginning(4)

	node, err := ll.FindNode(2)
	s.NoError(err)
	s.Equal(ll.tail.previous, node)

	node, err = ll.FindNode(10)
	s.Error(err)
	s.Nil(node)
}

func (s *LinkedListTestSuite) TestFindNodeBefore() {
	ll := NewLinkedList()
	ll.AddAtBegginning(2)
	ll.AddAtBegginning(4)

	node, err := ll.FindNodeBefore(2)
	s.NoError(err)
	s.Equal(node, ll.head.nextNode)

	node, err = ll.FindNodeBefore(10)
	s.Error(err)
	s.Nil(node)
}

func (s *LinkedListTestSuite) TestCopyList() {
	ll := NewLinkedList()
	ll.AddAtBegginning(2)
	ll.AddAtBegginning(4)
	ll.AddAtBegginning(6)
	ll.AddAtBegginning(8)

	llCopy := ll.CopyList()
	input := llCopy.head.nextNode
	for input != llCopy.tail {
		node, err := ll.FindNode(input.data)
		s.NoError(err)
		s.Equal(node.data, input.data)
		input = input.nextNode
	}
}

// next : InsertionsortWithCopy
