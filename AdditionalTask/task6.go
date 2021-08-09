package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Доп. задание:
1) Попробуйте разобраться в массивах и слайсах, понять в чем разница. Попробуйте реализовать матрицу на слайсах, которую можно ввести с клавиатуры произвольного размера. Реализуйте умножение матрицы на число.
*/
func main() {
	number,sizeSlice := 0
	fmt.Println("Программа  ")
  fmt.Print("Введите размер матрицы: ")
	fmt.Scan(&sizeSlice)

  fmt.Print("Введите число на которое надо умножить: ")
	fmt.Scan(&number)

  arr := make([][]int,sizeSlice,sizeSlice)
}