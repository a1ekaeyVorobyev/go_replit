package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeterminant(t *testing.T) {

	var a = [][]float64{
		{2, 1, 1},
		{3, 0, 1},
	}

	var b = [][]float64{
		{3, 1},
		{2, 1},
		{1, 0},
	}

	var result = [][]float64{
		{9, 3},
		{10, 3},
	}

	d, _ := multiplication(a, b)
	require.Equal(t, d, result)

	_, ok := multiplication(a, a)
	require.Equal(t, ok, false)

	a = [][]float64{
		{3, 2, 1},
		{0, 1, 2},
	}

	b = [][]float64{
		{1},
		{2},
		{3},
	}

	result = [][]float64{
		{10},
		{8},
	}

	d, _ = multiplication(a, b)
	require.Equal(t, d, result)

}
