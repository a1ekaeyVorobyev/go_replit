package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Задание 2. Вывод чисел содержащихся в строке

Что нужно сделать:

В начале программы объявите строку и присвойте ей: a10 10 20b 20 30c30 30 dd

Напишите программу, которая выведет все части строки, разделенные пробелами, которые можно привести к целому числу. Попробуйте ввести другие строки, например, пустую строку..
*/

func main() {
	fmt.Println("Программа Определение количества слов начинающихся с большой буквы")
	fmt.Print("Введите строку: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	cnt := numberWordToInt(text)
	fmt.Printf("Количество слов с заглавной буквы = %d. \n", cnt)
}

func numberWordToInt(text string) (cnt int) {
	t := strings.Split(text, " ")
	for _, v := range t {
		_, err := strconv.Atoi(v)
		if err == nil {
			cnt++
		}
	}
	return
}
