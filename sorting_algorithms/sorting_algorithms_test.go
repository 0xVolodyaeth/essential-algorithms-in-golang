package sorting_algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{2, 3, 1, 4, 7, 5, 6, 8, 10, 9}
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(t, exp, InsertionSort(arr))
}

func TestSelectionSort(t *testing.T) {
	arr := []int{2, 3, 1, 4, 7, 5, 6, 8, 10, 9}
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(t, exp, SelectionSort(arr))
}
func TestBubbleSort(t *testing.T) {
	arr := []int{2, 3, 1, 4, 7, 5, 6, 8, 10, 9}
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(t, exp, BubbleSort(arr))
}

// func TestMakeHeap(t *testing.T) {
// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// heap := makeHeap(arr)
// expectedHeap := []int{10, 9, 6, 7, 8, 2, 5, 1, 4, 3}
// require.Equal(t, expectedHeap, heap)
// }

func TestHeapSort(t *testing.T) {
	arr := []int{2, 3, 1, 4, 7, 5, 6, 8, 10, 9}
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(t, exp, HeapSort(arr))
}
