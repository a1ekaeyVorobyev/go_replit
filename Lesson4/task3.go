package main

import (
	"fmt"
)

/*
Задача 3. Банкомат

Пользователи банкомата хотят снимать деньги. Но банкомат умеет выдавать только купюры по 100 рублей, а максимальная сумма снятия — 100 000 рублей.

Напишите программу, которая проверяет допустимость суммы средств, которую ввёл пользователь. Сумма не должна быть меньше 100 и больше 100 000 руб.
*/
func main() {
	var money int
	maxSumma := 100000
	minSumma := 100

	fmt.Println("Программа Банкомат")
	fmt.Printf("Введите сумму для снятия:\n")
	fmt.Scan(&money)

	if money >= minSumma && money <= maxSumma {
		fmt.Printf("Вы можете снять сумму %v\n", money)
	} else if money < minSumma {
		fmt.Printf("Вы можете снять сумму %v , так как она меньше минимальной %v\n", money, minSumma)
	} else {
		fmt.Printf("Вы можете снять сумму %v , так как она больше максимальной %v\n", money, maxSumma)
	}

}