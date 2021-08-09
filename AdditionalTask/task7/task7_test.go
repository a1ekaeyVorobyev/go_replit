package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
Проверить на примерах:
1 1 результат uint8
1 -1 результат int8
640 100 результат uint16
-640 100 результат int32
3000 3000 результат uint32
-3000 3000 результат int32
*/

func TestCocktailSort(t *testing.T) {

	arr := []int{1, 2, 5, 2, 2, 6, 3, 2, 9, 1, 7}
	arrTest := []int{1, 1, 2, 2, 2, 2, 3, 5, 6, 7, 9}
	cocktailSort(arr)
	require.Equal(t, arr, arrTest)
}

func TestInsertionSort(t *testing.T) {

	arr := []int{1, 2, 5, 2, 2, 6, 3, 2, 9, 1, 7}
	arrTest := []int{1, 1, 2, 2, 2, 2, 3, 5, 6, 7, 9}
	insertionSort(arr)
	require.Equal(t, arr, arrTest)
}

func TestGnomeSort(t *testing.T) {

	arr := []int{1, 2, 5, 2, 2, 6, 3, 2, 9, 1, 7}
	arrTest := []int{1, 1, 2, 2, 2, 2, 3, 5, 6, 7, 9}
	gnomeSort(arr)
	require.Equal(t, arr, arrTest)
}

func TestMergeSort(t *testing.T) {

	arr := []int{1, 2, 5, 2, 2, 6, 3, 2, 9, 1, 7}
	arrTest := []int{1, 1, 2, 2, 2, 2, 3, 5, 6, 7, 9}
	arr = mergeSort(arr)
	require.Equal(t, arr, arrTest)
}

func TestBubbleSort(t *testing.T) {

	arr := []int{1, 2, 5, 2, 2, 6, 3, 2, 9, 1, 7}
	arrTest := []int{1, 1, 2, 2, 2, 2, 3, 5, 6, 7, 9}
	bubbleSort(arr)
	require.Equal(t, arr, arrTest)
}
