package main

import (
	"fmt"
)

/*
Задание 2. Функция, реверсирующая массив

Что нужно сделать

Напишите функцию, принимающую на вход массив и возвращающую массив, в котором элементы идут в обратном порядке по сравнению с исходным.
Напишите программу, демонстрирующую работу этого метода.


Что оценивается

При вводе 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 программа должна выводить при помощи дополнительной функции, реверсировав массив: 10, 9, 8, 7, 6, 5, 4, 3, 2, 1.
*/

func main() {
	arr := make([]int, 10)

	fmt.Println("Функция, реверсирующая массив")
	for i := range arr {
		fmt.Printf("Введите %d число: ", i)
		fmt.Scan(&arr[i])
	}
	fmt.Println("Входной массив   ", arr)
	arr = Reverse(arr)
	fmt.Println("Реверсный массив ", arr)
}

func Reverse(arr []int) []int {
	j := len(arr)
	cnt := j / 2
	for i := 0; i < cnt; i++ {
		j--
		fmt.Printf("меняем %d элемент со значанием %d на %d элемент со значанием %d. \n", i, arr[i], j, arr[j])
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}