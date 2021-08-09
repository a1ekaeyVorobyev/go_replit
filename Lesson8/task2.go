package main

import (
	"fmt"
	"strings"
)

/*
Задание 2:
Пользователь вводит будний день недели в сокращенной форме(пн, вт, ср, чт, пт) и получает развернутый список всех последующих рабочих дней, включая пятницу.
Пример:
пользователь вводит:
вт
программа дает ответ:
вторник
среда
четверг
пятница
*/
func main() {
	fmt.Println("Программа ")
	var dayOfWeek string
	fmt.Println("Введите день недели:")
	fmt.Scan(&dayOfWeek)

	dayOfWeek = strings.ToLower(dayOfWeek)

	switch dayOfWeek {
	case "пн":
		fmt.Println("понедельник")
		fallthrough
	case "вт":
		fmt.Println("вторник")
		fallthrough
	case "ср":
		fmt.Println("среда")
		fallthrough
	case "чт":
		fmt.Println("четверг")
		fallthrough
	case "пт":
		fmt.Println("пятница")
	case "сб":
		fmt.Println("суббота")
		fallthrough
	case "вс":
		fmt.Println("воскресенье")
	default:
		fmt.Println("Я не знаю текай день недели!")
	}
}
