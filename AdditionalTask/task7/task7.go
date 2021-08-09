package main

import (
	"fmt"
)

/*
1) Напишите сортировку массива пузырьком (bubble sort) и сортировку вставками.
*/

func main() {
	array := []int{1, 2, 5, 2, 2, 6, 3, 2, 9, 1, 7}
	arr := array

	fmt.Println("--bubbleSort--")
	fmt.Println(arr)
	bubbleSort(arr)
	fmt.Println(arr)

	fmt.Println("--cocktailSort--")
	arr = []int{1, 2, 5, 2, 2, 6, 3, 2, 9, 1, 7}
	fmt.Println(arr)
	cocktailSort(arr)
	fmt.Println(arr)

	fmt.Println("--insertionSort--")
	arr = []int{1, 2, 5, 2, 2, 6, 3, 2, 9, 1, 7}
	fmt.Println(arr)
	insertionSort(arr)
	fmt.Println(arr)

	fmt.Println("--gnomeSort--")
	arr = []int{1, 2, 5, 2, 2, 6, 3, 2, 9, 1, 7}
	fmt.Println(arr)
	gnomeSort(arr)
	fmt.Println(arr)

	fmt.Println("--mergeSort--")
	arr = []int{1, 2, 5, 2, 2, 6, 3, 2, 9, 1, 7}
	fmt.Println(arr)
	arr = mergeSort(arr)
	fmt.Println(arr)
}

//Сортировка пузырьком (англ. Bubble sort) — для каждой пары индексов производится обмен, если элементы расположены не по порядку. Сложность алгоритма: {\displaystyle O(n^{2})}O(n^{2}).
func bubbleSort(arr []int) {
	flag := true
	for flag {
		flag = false
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				flag = true
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
	}
}

//Сортировка перемешиванием (англ. Cocktail sort). Сложность алгоритма: {\displaystyle O(n^{2})}O(n^{2}).
func cocktailSort(arr []int) {
	left := 0
	right := len(arr) - 1
	for left < right {
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		right--
		for i := right; i > left; i-- {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
		left++
	}
}

//Сортировка вставками (англ. Insertion sort) — определяем, где текущий элемент должен находиться в упорядоченном списке, и вставляем его туда. Сложность алгоритма: {\displaystyle O(n^{2})}O(n^{2}).
func insertionSort(arr []int) {
	var n = len(arr)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j = j - 1
		}
	}
}

//Гномья сортировка (англ. Gnome sort; первоначально опубликована под названием «глупая сортировка» [stupid sort] за простоту реализации) — сходна с сортировкой вставками. Сложность алгоритма — {\displaystyle O(n^{2})}O(n^{2}); рекурсивная версия требует дополнительно {\displaystyle O(n^{2})}O(n^{2}) памяти.

func gnomeSort(arr []int) {
	for i, j := 1, 2; i < len(arr); {
		if arr[i-1] > arr[i] {
			arr[i-1], arr[i] = arr[i], arr[i-1]
			i--
			if i > 0 {
				continue
			}
		}
		i = j
		j++
	}
}

//Сортировка слиянием (англ. Merge sort) — выстраиваем первую и вторую половину списка отдельно, а затем объединяем упорядоченные списки. Сложность алгоритма: {\displaystyle O(n\log n)}O(n\log n). Требуется {\displaystyle O(n)}O(n) дополнительной памяти.

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	n := len(arr) / 2
	l := mergeSort(arr[:n])
	r := mergeSort(arr[n:])

	return Merge(l, r)
}

func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}
