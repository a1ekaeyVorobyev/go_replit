package main

import (
	"fmt"
)

/*
Задание 2. Сумма двух чисел по единице
Напишите программу, которая запрашивает у пользователя два числа и складывает их в цикле следующим образом: берёт первое число и прибавляет к нему по единице, пока не достигнет суммы двух чисел.
*/

func main() {
	var number1, number2 int
	fmt.Println("Программа   Сумма двух чисел по единице")

	fmt.Printf("Введите 1 число:\n")
	fmt.Scan(&number1)
	fmt.Printf("Введите 2 число:\n")
	fmt.Scan(&number2)
	sum := number1 + number2
	for number1 < sum {
		number1 += 1
	}
	fmt.Println("Сумма чисел =", number1)
}
