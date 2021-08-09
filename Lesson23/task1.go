package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int,10)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n * 10)
	}
  fmt.Println("---исходный---")
	fmt.Println(arr)
	a, b := f(arr)
  fmt.Println("---четный---")
	fmt.Println(a)
  fmt.Println("---не четный---")
	fmt.Println(b)
}

func f(arr []int) (a, b []int) {
	a = make([]int, 0, len(arr))
	b = make([]int, 0, len(arr))
	for _, v := range arr {
		if v%2 == 0 {
			a = append(a, v)
		} else {
			b = append(b, v)
		}
	}
	return
}
