package main

import (
	"fmt"
)

/*
Задача 2. Три числа

Напишите программу, которая запрашивает у пользователя три числа и сообщает, есть ли среди них число, большее, чем 5.
*/
func main() {
	var total, examScore int
	cntNumber := 3
	contolNumber := 5

	arr := make([]int, cntNumber)

	fmt.Println("Программа Три числа")
	for i := 1; i <= cntNumber; i++ {
		fmt.Printf("Введите %v число:\n", i)
		fmt.Scan(&arr[i-1])
		total += examScore
	}
	flag := true
	for i, n := range arr {
		if n > contolNumber {
			i += 1
			fmt.Printf("Число %v, которое ввели %v, больше 5:\n", n, i)
			flag = false
		}
	}
	if flag {
		fmt.Printf("Среде веденных чисел %v нет ниодного больше 5:\n", arr)
	}

}
