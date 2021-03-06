package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewHeap(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7}
	heap := NewHeap(&sl)
	require.NotNil(t, heap)

	heapSlice := []int{7, 4, 6, 1, 3, 2, 5}
	require.Equal(t, &heapSlice, heap.slice)
}

func TestTemoveTop(t *testing.T) {

	sl := []int{1, 2, 3, 4, 5, 6, 7}
	heap := NewHeap(&sl)
	require.NotNil(t, heap)

	require.Equal(t, 7, heap.RemoveTopItem())
	require.Equal(t, 6, heap.RemoveTopItem())
	require.Equal(t, 5, heap.Len())
}
