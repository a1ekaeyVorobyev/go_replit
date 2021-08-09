package main

import (
	"fmt"
	"math"
)

/*
1 задание:
Что нужно сделать
В данном модуле мы с вами рассмотрели примеры по целочисленным типам, их размерам в памяти и что происходит при ее переполнении. Теперь давайте напишем программу, которая будет в цикле с использованием встроенных констант (на предельные значения целых чисел, в пакете math) подсчитывать сколько приходится переполнений чисел типа uint8, uint16 в диапазоне от 0 до uint32.
*/

func main() {

	cntUint8 := 0
	cntUint16 := 0
	numberUint8 := 0
	numberUint16 := 0
	fmt.Println("Программа Задание 9.1.")

	for i := 0; i <= math.MaxUint32; i++ {
		if numberUint8 == math.MaxUint8 {
			cntUint8++
			numberUint8 = 0
		} else {
			numberUint8++
		}

		if numberUint16 == math.MaxUint16 {
			cntUint16++
			numberUint16 = 0
		} else {
			numberUint16++
		}
	}
	fmt.Printf("Переполнения типа uint8 было %d раз. \n", cntUint8)
	fmt.Printf("Переполнения типа uint16 было %d раз. \n", cntUint16)
}
