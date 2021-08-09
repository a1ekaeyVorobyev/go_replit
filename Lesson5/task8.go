package main

import (
	"fmt"
)

/*
8. Игра “Угадывание числа” (дополнительно)

Что нужно сделать

Ну, и какой же компьютер без игр? Давайте научим его играть в “угадывание чисел”. Пользователь загадывает число от 1 до 10 (включительно). Программа пытается это число угадать, для этого она выводит число, а пользователь должен ответить угадала программа, загаданное число больше или меньше.

Советы и рекомендации

Программа не должна делать больше 4 попыток в процессе угадывания.
*/

func main() {
	var val string
	numberOfAttempts := 4
	number := 0
	isGuess := false
	numberLimitMax := 10
	numberLimitMin := 1
	cntOfAttempts := 0

	for i := 1; i <= numberOfAttempts; i++ {
		cntOfAttempts += 1
		fmt.Println("Программа Угадывание числа")
		fmt.Println("Загадывает число от 1 до 10 (включительно)")
		number = (numberLimitMax-numberLimitMin)/2 + numberLimitMin
		fmt.Printf("Это число %v \n", number)
		fmt.Println("нажмите y, если угадали и n если нет")
		fmt.Scan(&val)
		if val == "y" {
			isGuess = true
			break
		} else {
			fmt.Println("нажмите 1 если число больше и 2 если меньше")
			fmt.Scan(&val)
			if val == "1" {
				numberLimitMin = number
			} else {
				numberLimitMax = number
			}
			fmt.Println(numberLimitMin, numberLimitMax)
		}
	}
	if isGuess {
		fmt.Printf("Мы угадали чило c %v попытки. Оно = %v\n", cntOfAttempts, number)
	}
}
