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

func TestHeapSort(t *testing.T) {
	arr := []int{2, 3, 1, 4, 7, 5, 6, 8, 10, 9}
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(t, exp, HeapSort(arr))
}

func TestQuickSort(t *testing.T) {
	arr := []int{2, 3, 1, 4, 7, 5, 6, 8, 10, 9}
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	QuickSort(arr, 0, len(arr)-1)
	require.Equal(t, exp, arr)
}
