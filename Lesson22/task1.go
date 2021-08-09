package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Задание 1. Подсчёт чётных и нечётных чисел в массиве


Что нужно сделать

Заполните массив неупорядоченными числами на основе генератора случайных чисел. Введите число. Программа должна найти это число в массиве и вывести, сколько чисел находится в массиве после введённого. При отсутствии введённого числа в массиве — вывести 0. Для удобства проверки реализуйте вывод массива на экран.
*/
const n = 10

func main() {
	rand.Seed(time.Now().UnixNano())
	var arr [n]int
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n * 10)
	}
	value := arr[rand.Intn(n)]
	fmt.Printf("Ищем %v в массиве %v. \n", value, arr)
	index := find(value, arr)
	if index > 0 {
		fmt.Printf("Число %v находится в массиве, порядковый номер %v. \n", value, index)
		fmt.Println(arr[index-1:])
	} else {
		fmt.Printf("Число %v не находится в массиве, порядковый номер  %v. \n", value, index)
	}

	value = 10000
	fmt.Printf("Ищем %v в массиве %v. \n", value, arr)
	index = find(value, arr)
	if index > 0 {
		fmt.Printf("Число %v находится в массиве, порядковый номер %v. \n", value, index-1)
		fmt.Println(arr[index-1:])
	} else {
		fmt.Printf("Число %v не находится в массиве. Значение порядкового номера = %v. \n", value, index)
	}

}

func find(number int, arr [n]int) int {
	for i, n := range arr {
		if n == number {
			return i + 1
		}
	}
	return 0
}
