package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Задание 2. Нахождение первого вхождения числа в упорядоченном массиве (числа могут повторяться)


Что нужно сделать

Заполните упорядоченный массив из 12 элементов и введите число. Необходимо реализовать поиск первого вхождения заданного числа в массив. Сложность алгоритма должна быть минимальная.



Что оценивается

Верность индекса.

При вводе массива 1 2 2 2 3 4 5 6 7 8 9 10 и вводе числа 2 программа должна вывести индекс 1.
*/
const n = 10

func main() {
	rand.Seed(time.Now().UnixNano())
	var arr [n]int
	for i := 0; i < n; i++ {
		arr[i] = 10*i + rand.Intn(10)
	}
	value := arr[rand.Intn(n)]
	fmt.Printf("Ищем %v в массиве %v. \n", value, arr)
	index := find(value, arr)

	fmt.Println(index)

	for i := 0; i < n; i++ {
		arr[i] = 1
	}
	value = arr[rand.Intn(n)]
	fmt.Printf("Ищем %v в массиве %v. \n", value, arr)
	index = find(value, arr)
	fmt.Println(index)
}

func find(number int, arr [n]int) int {
	min := 0
	max := len(arr) - 1
	if arr[0] > number || arr[max] < number {
		return -1
	}
	for max >= min {

		middle := (max + min) / 2
		if number == arr[middle] {
			index := isFirstElement(middle, arr)
			return index
		} else if number > arr[middle] {
			min = middle + 1
		} else {
			max = middle - 1
		}

	}
	return -1
}

func isFirstElement(index int, arr [n]int) int {
	if index > 0 && arr[index] == arr[index-1] {
		return isFirstElement(index-1, arr)
	}
	return index
}
