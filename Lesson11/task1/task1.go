package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
Задание 1. Определение количества слов начинающихся с большой буквы

Что нужно сделать:

В начале программы объявите строку и присвойте ей: Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software

Напишите программу, которая выведет количество слов начинающихся с большой буквы. Попробуйте ввести другие строки, например, пустую строку.
*/

func main() {

	fmt.Println("Программа Определение количества слов начинающихся с большой буквы")
	fmt.Print("Введите строку: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	cnt := numberWordCapitalLetter(text)
	fmt.Printf("Количество слов с заглавной буквы = %d. \n", cnt)
}

func numberWordCapitalLetter(text string) (cnt int) {
	t := strings.Split(text, " ")
	for _, v := range t {
		t := []rune(v)
		if unicode.IsUpper(t[0]) {
			cnt++
		}
	}
	return
}
