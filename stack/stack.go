package stack

import "github.com/vpatsenko/essential-algorithms-in-golang/doubleLinkedList"

type Stack struct {
	ll  *doubleLinkedList.LinkedList
	len int
}

func NewStack() *Stack {
	return &Stack{
		ll: doubleLinkedList.NewLinkedList(),
	}
}

func (s *Stack) Push(value int) {
	s.ll.AddAtBegginning(value)
	s.len++
}

func (s *Stack) Pop() (int, error) {
	poppedNode, err := s.ll.Head().NextNode()
	if err != nil {
		return 0, err
	}
	s.ll.DeleteAfter(s.ll.Head())
	s.len--
	return poppedNode.Data(), nil
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) IsEmpty() (bool, error) {
	nodeAfterHead, err := s.ll.Head().NextNode()
	if err != nil {
		return false, err
	}

	if nodeAfterHead == s.ll.Tail() {
		return true, nil
	}
	return false, nil
}
