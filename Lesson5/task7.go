package main

import (
	"fmt"
)

/*
7. Счастливый билет

Что нужно сделать
Все мы в детстве, да и не только в детстве, получив бумажный билет, проверяли, а не является ли он “счастливым”. Давайте напишем программу, в которую пользователь будет вводить четырехзначный номер билета, а программа будет выводить, является ли он счастливым (сумма старших двух разрядов равна сумме двух младших разрядов) или даже зеркальным.


Например:
1234 -> обычный билет
3425 -> счастливый билет
1221 -> зеркальный билет
*/

func main() {
	var numberTicket, number1, number2, number3, number4 int

	fmt.Println("Программа Решение квадратного уравнения")

	fmt.Printf("Введите номер билетв a:\n")
	fmt.Scan(&numberTicket)
	number1 = numberTicket / 1000
	number2 = numberTicket/100 - number1*10
	number3 = numberTicket/10 - number1*100 - number2*10
	number4 = numberTicket - number1*1000 - number2*100 - number3*10
	if number1 == number4 && number2 == number3 {
		fmt.Printf("%v -> зеркальный билет\n", numberTicket)
	} else if (number1 + number2) == (number3 + number4) {
		fmt.Printf("%v -> счастливый билет\n", numberTicket)
	} else {
		fmt.Printf("%v -> обычный билет\n", numberTicket)
	}
}
