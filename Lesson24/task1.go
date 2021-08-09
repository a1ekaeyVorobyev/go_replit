package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Напишите функцию, сортирующую массив длины 10 вставками.
*/

const sizeArray = 10

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 10)
	for i := 0; i < sizeArray; i++ {
		arr[i] = rand.Intn(100)
	}
  fmt.Println("--до сортировки--")
	fmt.Println(arr)
	InsertionSort(arr)
  fmt.Println("--после сортировки--")
	fmt.Println(arr)
}

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i
		for ; j >= 1 && arr[j-1] > x; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = x
	}
}
