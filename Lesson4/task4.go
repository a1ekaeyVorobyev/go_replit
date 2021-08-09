package main

import (
	"fmt"
)

/*
Задача 4. Три числа: еще попытка

Напишите программу, которая запрашивает у пользователя три числа и выводит количество чисел, которые больше, либо равны 5.
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
	cnt := 0
	for i, n := range arr {
    fmt.Printf("i is %d berfore\n", i)
		if n > contolNumber {
			i += 1
			fmt.Printf("Число %v, которое ввели %v, больше 5:\n", n, i)
			cnt++
		}
    fmt.Printf("i is %d after\n", i)
	}
	if cnt == 1 {
		fmt.Printf("Вы ввели %v число больше 5:\n", cnt)
	} else if cnt > 1 && cnt < 5 {
		fmt.Printf("Вы ввели %v чисела больше 5:\n", cnt)
	} else if cnt > 4 {
		fmt.Printf("Вы ввели %v чисел больше 5:\n", cnt)
	} else {
		fmt.Printf("Среде веденных чисел %v нет ниодного больше 5:\n", arr)
	}

}
