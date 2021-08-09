package main

import (
	"fmt"
)

func main() {
	var startHeight, growth, ate, finishHeight int

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
	fmt.Println("Задайте  нужный рост бамбука в м.")
	fmt.Scan(&finishHeight)
	startHeight *= 100
	finishHeight *= 100
	temp := ate
	cntDay := 0
	for temp < finishHeight {
		temp += (growth - ate)
		cntDay++
	}
	fmt.Printf("Бамбук вырастет за %v дней с %v до %v см.\n", cntDay, startHeight, finishHeight)
}
