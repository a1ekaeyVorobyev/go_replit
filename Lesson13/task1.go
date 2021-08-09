package main

import (
	"fmt"
)

/*
Задача
Цель задания: написать функцию, которая принимает две функции на вход и вызывает их в противоположном порядке.

Что входит в домашнее задание: две маленькие функции, одна большая, принимающая две маленькие на вход.

Что нужно сделать: две маленькие функции, одна большая, принимающая две маленькие на вход.

Советы: подсказка - функции в качестве аргумента через func f(VARIABLE_NAME func(INPUT_TYPES)) - пример, func f(g func(int))

Что оценивается: наличие двух маленьких функций, правильная сигнатура у второй и успешная сборка и выполнение программы с вызовом большой функции.
*/

func main() {
  d:= change(2,3,Add,Multiplication)
  fmt.Println("(x*y)+2",d)

  d = change(2,3,Multiplication,Add)
  fmt.Println("(x+y)*2",d)
}

func change(x,y int,f1 func(int,int) int, f2 func(int,int) int) int{
  res1:= f2(x,y)
  res2:= f1(res1,2)
  return res2
}

func Add(x,y int)int{
  return x+y
}

func Multiplication(x,y int)int{
  return x*y
}