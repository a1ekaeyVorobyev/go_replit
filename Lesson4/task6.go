package main

import (
	"fmt"
)

/*
Задача 6. Студенты* (сложная задача)

Перед старостой группы стоит задача разделить весь курс, состоящий из N студентов, на K групп. Напишите программу, которая поможет старосте сделать это: он вводит N, K и порядковый номер студента, а программа определяет, в какую группу он попадает.

Подсказка: в одну группу могут попадать студенты из разных частей списка.
*/
func main() {
	var cntStudent, cntGroup, numberStudent, groupStudent int
	fmt.Println("Программа Студенты")
	fmt.Printf("Введите количество студентов:\n")
	fmt.Scan(&cntStudent)
	fmt.Printf("Введите количество групп:\n")
	fmt.Scan(&cntGroup)
	fmt.Printf("Введите порядковый номер студента:\n")
	fmt.Scan(&numberStudent)
	cntStudetnOfGroup := cntStudent / cntGroup
	if cntStudent%cntGroup > 0 {
		cntStudetnOfGroup++
	}
	if numberStudent > cntStudent {
		fmt.Printf("Номер студента %v не может быть больше общего числа студентов %v:\n", numberStudent, cntStudent)
		return
	}
	if numberStudent <= cntGroup {
		groupStudent = numberStudent
	} else {
		groupStudent = numberStudent % cntGroup
		if groupStudent == 0 {
			groupStudent = cntGroup
		}
	}
	fmt.Printf("Студент под номером %v попал в группу № %v.\n", numberStudent, groupStudent)
}