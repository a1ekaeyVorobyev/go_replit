package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFind(t *testing.T) {

	arr := [10]int{1, 2, 2, 2, 3, 3, 5, 6, 7, 8}
	index := find(2, arr)
	require.Equal(t, 1, index)

	index = find(12, arr)
	require.Equal(t, -1, index)

	index = find(4, arr)
	require.Equal(t, -1, index)

	index = find(0, arr)
	require.Equal(t, -1, index)
}
