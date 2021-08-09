package main

import (
	"fmt"
)

func main() {
	var startHeight, growth, ate, cntDay int

	fmt.Println("Программа расчета роста бамбука")
	fmt.Println("Задайте начальный рост бамбука в м.:")
	fmt.Scan(&startHeight)
	fmt.Println("Задайте рост бамбука в день в см.:")
	fmt.Scan(&growth)
	fmt.Println("Задайте скорость поедания бамбука гусеницами в см.")
	fmt.Scan(&ate)
	if growth < ate {
		fmt.Printf("Бамбук не вырастет так скорость поедания (%v) больше скорости роста (%v). \n", ate, growth)
		return
	}
	fmt.Println("Задайте количество дней роста бамбука")
	fmt.Scan(&cntDay)
	startHeight *= 100
	height := (growth-ate)*(cntDay-1) + growth
	fmt.Printf("Бамбук за %v дней вырос на %v. И его рост состовляет %v.", cntDay, height, startHeight+height)
}
