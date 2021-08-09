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

* в качестве дополнительного задания можно во втором задании найти еще числа представленные в восьмеричной и шестнадцатеричной системах
*/

func main() {
	fmt.Println("Программа Вывод чисел содержащихся в строке")
	fmt.Print("Введите строку: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	cnt := numberWordToInt(text)
	fmt.Printf("Количество чисел = %d. \n", cnt)
  cnt := cntHexVakue(text,8)
	fmt.Printf("Количество чисел в 8 системе = %d. \n", cnt)
   cnt := cntHexVakue(text,16)
	fmt.Printf("Количество слов в 16 системе = %d. \n", cnt)
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

func cntHexVakue(text string,main int)(cnt int){
  t := strings.Split(text, " ")
	for _, v := range t {
		_, err := strconv.ParseInt(hex, main, 64)
		if err == nil {
			cnt++
		}
	}
	return
}