package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStack(t *testing.T) {
	st := NewStack()
	require.NotNil(t, st)
}

func TestPush(t *testing.T) {
	st := NewStack()
	st.Push(1)
	st.Push(2)

	node, err := st.ll.Head().NextNode()
	require.NoError(t, err)
	require.Equal(t, 2, node.Data())
	require.Equal(t, 2, st.Len())
}

func TestPop(t *testing.T) {
	st := NewStack()
	st.Push(1)
	st.Push(2)

	popped, err := st.Pop()
	require.NoError(t, err)

	require.Equal(t, 2, popped)
}

func TestIsEmpty(t *testing.T) {
	st := NewStack()
	st.Push(1)
	st.Push(2)

	_, err := st.Pop()
	require.NoError(t, err)

	_, err = st.Pop()
	require.NoError(t, err)

	isEmpty, err := st.IsEmpty()
	require.NoError(t, err)
	require.True(t, isEmpty)
}
