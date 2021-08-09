// go test -bench=. -benchmem
// go test -bench=. -benchmem

package main

import (
	"math/rand"
	"testing"
)

const MAX_RAND_LIMIT = 255
const MAX_LEN_ARRAY = 1024

var arr []int = generateRandomArray(1024)

func generateRandomArray(arrayLen int) []int {
	var a []int
	for i := 0; i < arrayLen; i++ {
		a = append(a, rand.Intn(MAX_RAND_LIMIT))
	}
	return a
}

func BenchmarkCocktailSort(b *testing.B) {
	a := arr
	cocktailSort(a)
}

func BenchmarkInsertionSort(b *testing.B) {
	a := arr
	insertionSort(a)
}

func BenchmarkGnomeSort(b *testing.B) {
	a := arr
	gnomeSort(a)
}

func BenchmarkMergeSort(b *testing.B) {
	a := arr
	mergeSort(a)
}

func BenchmarkBubbleSort(b *testing.B) {
	a := arr
	bubbleSort(a)
}
