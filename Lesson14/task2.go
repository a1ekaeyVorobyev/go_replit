package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Задание 2. Функция возвращающая несколько значений


Что нужно сделать

Напишите программу, которая с помощью функции генерирует 3 случайные точки в двумерном пространстве (две координаты), а затем с помощью другой функции преобразует эти координаты по формулам: x = 2 * x + 10, y = -3 * y - 5.
*/

type Point struct {
	x int
	y int
}

const constSize = 10
const constSizeArray = 3

func main() {
	fmt.Println(" Функция возвращающая несколько значений")
	list := make([]Point, constSizeArray)
	fmt.Println("==Создание===")
	for i := range list {
		list[i].Generate()
		fmt.Println(list[i])
	}
	fmt.Println("==Трансформация===")
	for i := range list {
		//list[i].Transformation()
		list[i].x, list[i].y = Transformation1(list[i])
		fmt.Println(list[i])
	}
}

func (p *Point) Generate() {
	rand.Seed(time.Now().UnixNano())
	p.x = rand.Intn(constSize)
	p.y = rand.Intn(constSize)
	return
}

func (p Point) String() string {
	return fmt.Sprintf("[x = %d, y= %d]", p.x, p.y)
}

func (p *Point) Transformation() {
	p.x = 2 * p.x
	p.y = -3*p.y - 5
}

func Transformation1(p Point) (int, int) {
	return 2 * p.x, -3*p.y - 5
}
