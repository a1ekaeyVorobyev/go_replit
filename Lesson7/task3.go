package main

import (
	"fmt"
)

/*
3. Вывод елочки
Что нужно сделать

Усложним задачу рисования, и попробуем вывести елочку.

В первой строке необходимо вывести одну звездочку, во второй - на две больше, в третьей - еще на две больше и так до количества строк указанных пользователем.
*/
func main() {
	sizeTree := 0
	fmt.Println("Программа  Вывод елочки")

	fmt.Printf("Введите размер елки:\n")
	fmt.Scan(&sizeTree)

	if sizeTree < 1 {
		fmt.Println("размер елки должен быть больше 0. Вы ввели", sizeTree)
		return
	}
  cntPoint := (sizeTree-1)*2+1
	fmt.Printf("--елочка--\n")
	for i := 0; i < sizeTree; i++ {
		for j := 0; j < cntPoint; j++ {
			if j>(sizeTree-2-i) && j<(sizeTree+i){
				fmt.Print("*")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Print("\n")
	}
}
