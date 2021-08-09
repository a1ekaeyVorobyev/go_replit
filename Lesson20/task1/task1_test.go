package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeterminant(t *testing.T) {

	var result float64

	var b = [][]float64{
		{1, -2, 3},
		{4, 0, 6},
		{-7, 8, 9},
	}
	d := determinant(b)
	result = 204
	require.Equal(t, d, result)

	b = [][]float64{
		{1, 0, -2},
		{3, 2, 1},
		{1, 2, -2},
	}
	d = determinant(b)
	result = -14
	require.Equal(t, d, result)

}

func TestSortArithmeticMean(t *testing.T) {
	var b = [][]float64{
		{10, -2, 3},
		{4, 1, 6},
		{-7, 8, 9},
	}
	var ideal = [][]float64{
		{-7, 8, 9},
		{10, -2, 3},
		{4, 1, 6},
	}
	result := sortArithmeticMean(b)
	require.Equal(t, ideal, result)
}

func TestIsValidation(t *testing.T) {
	var b = [][]float64{
		{10, -2, 3},
		{4, 1, 6},
		{-7, 8, 9},
	}
	result := isValidation(b)
	require.Equal(t, false, result)
	b = [][]float64{
		{-2, -1, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	result = isValidation(b)
	require.Equal(t, true, result)
}
