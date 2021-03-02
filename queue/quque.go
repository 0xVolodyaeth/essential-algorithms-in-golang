package queue

import (
	"errors"

	"github.com/vpatsenko/essential-algorithms-in-golang/doubleLinkedList"
)

type Queue struct {
	ll  *doubleLinkedList.LinkedList
	len int
}

func NewQueue() *Queue {
	return &Queue{doubleLinkedList.NewLinkedList(), 0}
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) Enqueue(val int) {
	q.ll.AddAtEnd(val)
	q.len++
}

func (q *Queue) Dequeue() (int, error) {

	next, err := q.ll.Head().NextNode()
	if err != nil {
		return 0, err
	}
	if next == q.ll.Tail() {
		return 0, errors.New("no values in queue")
	}

	node, err := q.ll.Head().NextNode()
	if err != nil {
		return 0, err
	}

	q.ll.DeleteAfter(q.ll.Head())
	q.len--

	return node.Data(), nil
}
