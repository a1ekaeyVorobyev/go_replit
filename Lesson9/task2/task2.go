package main

import (
	"fmt"
)

/*
2 задание:
Что нужно сделать
Достаточно часто, при передаче по сети или сохранении больших объемов данных, приходится выбирать тип с минимальным размером в памяти, чтобы экономить трафик или место на диске. Напишите программу, в которую пользователь вводит 2 числа (int16), а программа выводит в какой минимальный тип данных можно сохранить результат умножения этих чисел.
*/
const (
	MaxInt8   = int64(1<<7 - 1)
	MinInt8   = int64(-1 << 7)
	MaxInt16  = int64(1<<15 - 1)
	MinInt16  = int64(-1 << 15)
	MaxInt32  = int64(1<<31 - 1)
	MinInt32  = int64(-1 << 31)
	MaxInt64  = int64(1<<63 - 1)
	MinInt64  = int64(-1 << 63)
	MaxUint8  = int64(1<<8 - 1)
	MaxUint16 = int64(1<<16 - 1)
	MaxUint32 = int64(1<<32 - 1)
	//MaxUint64 = int64(1<<64 - 1)
)

func main() {
	var firstNumber, secondNumber int16
	var summa int64
	fmt.Println("Программа Задание 9.2.")

	fmt.Printf("Введите первое число:\n")
	fmt.Scan(&firstNumber)

	fmt.Printf("Введите второе число:\n")
	fmt.Scan(&secondNumber)
	summa = int64(firstNumber * secondNumber)
	t := typeMin(summa)
	fmt.Printf("Минимальный тип для знаяения %d является тип %s. \n", summa, t)
}
func typeMin(summa int64) string {
	if summa > 0 {
		switch {
		case summa < MaxUint8:
			return "uint8"
		case summa < MaxUint16:
			return "uint16"
		case summa < MaxUint32:
			return "uint32"
		default:
			return "uint64"
		}
	}
	switch {
	case summa >= MinInt8 && summa <= MaxUint8:
		return "int8"
	case summa >= MinInt16 && summa <= MaxUint16:
		return "int16"
	case summa >= MinInt32 && summa <= MaxUint32:
		return "int32"
	default:
		return "int64"
	}
}
