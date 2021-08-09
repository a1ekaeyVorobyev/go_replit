package main

import (
	"fmt"
)

/*
Задание 1. Подсчёт чётных и нечётных чисел в массиве

Что нужно сделать

Разработайте программу, позволяющую ввести 10 целых чисел, а затем вывести из них количество чётных и нечётных чисел. Для ввода и подсчёта используйте разные циклы.



Что оценивается

Для введённых чисел 1, 1, 1, 2, 2, 2, 3, 3, 3, 4 программа должна вывести: чётных — 4, нечётных — 6.
*/

func main() {
	arr :=make([]int,10)

	fmt.Println("Подсчёт чётных и нечётных чисел в массиве")
	for i := range arr {
		fmt.Printf("Введите %d число: ", i)
		fmt.Scan(&arr[i])
	}
	cntParity, cntNotParity := СountingParity(arr)
	fmt.Printf("чётных — %d, нечётных — %d \n", cntParity, cntNotParity)
}

func СountingParity(arr []int) (int, int) {
	cntParity := 0
	cntNotParity := 0
	for _,x := range arr {
		if x%2 == 0 {
			cntParity++
		} else {
			cntNotParity++
		}
	}
	return cntParity, cntNotParity
}
