package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseArray(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	reversed, err := ReverseArray(slice)
	require.NoError(t, err)

	for idx, val := range reversed {
		require.Equal(t, val, slice[idx])
	}
}
