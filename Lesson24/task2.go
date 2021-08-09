package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Напишите функцию, которая на вход получает сколько угодно интов, сортирует их пузырьком и переворачивается (либо сразу сортирует в обратном порядке, как посчитаете нужным).
*/

const sizeArray = 10

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 10)
	for i := 0; i < sizeArray; i++ {
		arr[i] = rand.Intn(100)
	}

	sortAndTurnOver := func(arr ...int) (result []int) {
		result = make([]int, len(arr))
		copy(result, arr)
		bubbleSort(result)
		size := len(result) / 2
		for i := 0; i < size; i++ {
			result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
		}
		return result
	}

	t := sortAndTurnOver(arr...)
	fmt.Println("--до сортировки--")
	fmt.Println(arr)
	fmt.Println("--после сортировки--")
	fmt.Println(t)

	t = sortAndTurnOver(1, 9, 2, 4, 3)
	fmt.Println("--до сортировки--")
	fmt.Println("1, 9, 2, 4, 3")
	fmt.Println("--после сортировки--")
	fmt.Println(t)
}

func bubbleSort(arr []int) {
	flag := true
	for flag {
		flag = false
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				flag = true
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
	}
}
