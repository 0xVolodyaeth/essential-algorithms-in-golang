package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewQueue(t *testing.T) {
	require.NotNil(t, NewQueue())
}

func TestEqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)

	afterHead, err := q.ll.Head().NextNode()
	require.NoError(t, err)
	require.Equal(t, 1, afterHead.Data())
	require.Equal(t, 1, q.Len())
}

func TestDequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)

	afterHead, err := q.ll.Head().NextNode()
	require.NoError(t, err)
	require.Equal(t, 1, afterHead.Data())
	require.Equal(t, 1, q.Len())

	val, err := q.Dequeue()
	require.NoError(t, err)
	require.Equal(t, 1, val)
	require.Equal(t, 0, q.Len())
}
