package main

import (
	"fmt"
)

/*
Задание 4. Область видимости переменных


Что нужно сделать

Напишите программу, в которой будет 3 функции, попарно использующие разные глобальные переменные. Функции, должны прибавлять к поданном на вход числу глобальную переменную и возвращать результат. Затем вызвать по очереди три функции, передавая результат из одной в другую.
*/
const constNumber = 10

var (
	gNumber1 = 20
	gNumber2 = 2
	gNumber3 = 10
)

func main() {
	var x int

	fmt.Println(" Именованные возвращаемые значения")

	fmt.Printf("Введите число:\n")
	fmt.Scan(&x)
	result := Multiplication(x)
	result = Add(result)
	fmt.Println("(x*10)+10=", result)
	result = f1(x)
	fmt.Println("(x+gNumber1)*gNumber2/gNumber3 = ", result)
}

func Add(x int) (result int) {
	result = x + constNumber
	return
}

func Multiplication(x int) (result int) {
	result = x * constNumber
	return
}

func f1(x int) int {
	x += gNumber1
	return f2(x)
}

func f2(x int) int {
	x *= gNumber2
	return f3(x)
}

func f3(x int) int {
	return x / gNumber3
}
