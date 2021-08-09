package main

import (
	"fmt"
)

/*
Задание 1. Слияние отсортированных массивов


Что нужно сделать

Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.
*/

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{10, 9, 8, 7, 6, 5}
	result := MergeArray(a[:], b[:])
	fmt.Println(result)
}

func MergeArray(a, b []int) []int {
	return append(a, b...)
}
