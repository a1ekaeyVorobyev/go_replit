package main

import (
	"fmt"
)

func main() {
	var b = [][]float64{
		{1, -2, 3},
		{4, 0, 6},
		{-7, 8, 9},
	}
	d := determinant(b)
	fmt.Printf("Для матрицы: %v \n", b)
	fmt.Printf("Опредилитель = %v \n", d)

}

func determinant(a [][]float64) (result float64) {
	if len(a) == 1 {
		return a[0][0]
	}
	sign := 1
	for i, x := range a[0] {
		result += float64(sign) * x * determinant(excludeColumn(a[1:], i))
		sign *= -1
	}
	return
}

func excludeColumn(a [][]float64, i int) (result [][]float64) {
	result = make([][]float64, len(a))
	size := len(a[0]) - 1
	for j, row := range a {
		r := make([]float64, size)
		copy(r[:i], row[:i])
		copy(r[i:], row[i+1:])
		result[j] = r
	}
	return
}

/*
доп. задание:
1) Напишите функцию, которая отсортирует строки в массиве по среднему арифметическому значению элементов этих строк.

по вашему вопросу:
1) Сделайте валидацию, которая будет проверять, что подан отсортированный массив.
*/
func sortArithmeticMean(arr [][]float64) [][]float64 {
	flag := true
	for flag {
		flag = false
		for i := 1; i < len(arr); i++ {

			if arithmeticMean(arr[i]) < arithmeticMean(arr[i-1]) {
				flag = true
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
	}
	return arr
}

func arithmeticMean(a []float64) float64 {
	var value float64
	for _, val := range a {
		value += val
	}
	value = value / (float64)(len(a))
	return value
}

func isValidation(arr [][]float64) bool {
	var temp float64
	fmt.Println("массив =", arr)
	for i, row := range arr {
		for j, value := range row {
			if i == 0 && j == 0 {
				temp = value
				continue
			}
			if value < temp {
				return false
			}
			temp = value
		}
	}
	return true
}
