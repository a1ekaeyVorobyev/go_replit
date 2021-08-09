package main

import (
	"fmt"
	"strconv"
)

/*
1. Зеркальные билеты

Что нужно сделать

Вывести сколько билетов находится среди всех шестизначных номеров от 100000 до 999999.
*/
func main() {

	minNumber := 100000
	maxNumber := 999999
	cnt := 0
	fmt.Println("Программа Зеркальные билеты")

	for i := minNumber; i <= maxNumber; i++ {
		text := strconv.Itoa(i)
    numberOfDigits := len(text)
    if numberOfDigits%2>0{
      continue
    }
		leftPart := text[:numberOfDigits/2]
    var rigthPart string
    for j:=numberOfDigits-1;j>=numberOfDigits/2;j--{
 		  rigthPart += string(text[j])
    }
 		if leftPart == rigthPart {
			cnt++
		}
	}
	fmt.Printf("В диапазоне билетов %v и %v, находится %v зеркальных.\n", minNumber, maxNumber, cnt)
}
