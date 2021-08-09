package main

import (
	"fmt"
	"math"
)

/*
Задача
Задание 1: Разложение ex в ряд Тейлора
e^x = 1 +x+x^2/2!+x^3/3!
для x = 0 и n = 1 вывод должен быть 1
для x = 1 и n = 3 вывод должен быть 2.7182539682539684
для x = 1 и n = 5 вывод должен быть 2.7182815255731922
*/

func main() {
	var x, n int
	fmt.Println("Программа Разложение e^x в ряд Тейлора")
	fmt.Print("Введите x: ")
	fmt.Scan(&x)
	fmt.Print("Введите количество знаков после ,: ")
	fmt.Scan(&n)
	result := ePowX(x, n)
	fmt.Println("result=", result)
}

func ePowX(x, n int) float64 {
	epsolon := 1 / math.Pow(10, float64(n))
	result := 1.0 + float64(x)
	previewResult := 0.0
	fact := 1
	k := 2

	for {
		if k > 0 {
			fact *= k
		}
		result += math.Pow(float64(x), float64(k)) / float64(fact)
		if math.Abs(result-previewResult) < epsolon {
			break
		}
		k++
		previewResult = result
	}
	return result
}
