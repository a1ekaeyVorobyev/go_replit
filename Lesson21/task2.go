package main

import (
	"fmt"
)

/*
Что нужно сделать

Напишите функцию, которая на вход принимает функцию вида A func (int, int) int, а внутри оборачивает и вызывает её при выходе (через defer).

Вызовите эту функцию с тремя разными анонимными функциями A. Тела функций могут быть любыми, но главное, чтобы все три выполняли разное действие.
*/
func main() {
	f := func(x, y int) int { return x + y }
  m := func(x, y int) int { return x * y }
	action(f)
  action(m)            
}

func action(f func(x, y int) int) {
	var x, y int

	defer func() {
		fmt.Println("Результат:",f(x, y))
	}()
	fmt.Println("Введите значение x:")
	fmt.Scan(&x)
	fmt.Println("Введите значение y:")
	fmt.Scan(&y)
}
