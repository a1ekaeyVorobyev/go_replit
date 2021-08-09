package main

import (
	"fmt"
)

func main() {
	var a = [][]float64{
		{2, 1, 1},
		{3, 0, 1},
	}

	var b = [][]float64{
		{3, 1},
		{2, 1},
		{1, 0},
	}

	c, ok := multiplication(a, b)
	fmt.Println("Умножаем:")
	fmt.Printf("матрицу: %v \n", a)
	fmt.Printf("на матрицу: %v \n", b)
	if ok {
		fmt.Printf("Результат = %v \n", c)
	} else {
		fmt.Println("Переумножить матрицы не получилось")
	}
}

func multiplication(a, b [][]float64) (result [][]float64, ok bool) {
	/*
	  проверка условий, что Произведение двух матриц АВ имеет смысл только в том случае, когда число столбцов матрицы А совпадает с числом строк матрицы В
	*/
	if len(a[0]) != len(b) {
		return
	}
	ok = true
	/*
	  В произведении матриц АВ число строк равно числу строк матрицы А , а число столбцов равно числу столбцов матрицы В .
	*/
	result = make([][]float64, len(a))
	size := len(b[0])
	fmt.Println("Размер матрицы ", size, len(a))
	for i := 0; i < len(a); i++ {
		r := make([]float64, size)
		for j := 0; j < size; j++ {
			//
			for k := 0; k < len(b); k++ {
				r[j] += a[i][k] * b[k][j]
			}
		}
		result[i] = r
	}
	return
}
