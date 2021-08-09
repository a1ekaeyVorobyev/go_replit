package main

import (
	"fmt"
)

/*
Задача 1. Баллы ЕГЭ

Сумма проходных баллов в институт составляет 275 баллов.

Напишите программу, которая запрашивает у пользователя результаты ЕГЭ по трём экзаменам и проверяет, поступил ли он в институт или нет.
*/
func main() {
	var total, examScore int
	cntOfExams := 3
	passingGrade := 275

	fmt.Println("Программа Баллы ЕГЭ")
	for i := 1; i <= cntOfExams; i++ {
		fmt.Printf("Введите оценку за %v экзамен:\n", i)
		fmt.Scan(&examScore)
		total += examScore
	}
	if total > passingGrade {
		fmt.Printf("Вы прошли. Набрали %v из необходимых %v:\n", total, passingGrade)
	} else if total == passingGrade {
		fmt.Printf("Вы eле прошли. Набрали %v из необходимых %v:\n", total, passingGrade)
	} else {
		fmt.Printf("Вы не прошли. Набрали %v из необходимых %v:\n", total, passingGrade)
	}
}