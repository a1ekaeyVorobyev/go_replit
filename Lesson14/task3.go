package main

import (
	"fmt"
)

/*
Задание 3. Именованные возвращаемые значения


Что нужно сделать

Напишите программу, которая на вход получает число, затем с помощью двух функции преобразует его. Первая умножает, а вторая прибавляет число, используя именованные возвращаемые значения
*/

func main() {
	var x int

	fmt.Println(" Именованные возвращаемые значения")

	fmt.Printf("Введите число:\n")
	fmt.Scan(&x)
	result := Multiplication(x)
	result = Add(result)
	fmt.Println("(x*x)+(x*x)=", result)

}

func Add(x int) (result int) {
	result = x + x
	return
}

func Multiplication(x int) (result int) {
	result = x * x
	return
}
