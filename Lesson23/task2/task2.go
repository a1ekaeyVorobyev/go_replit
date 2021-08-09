package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
Задание 2. Поиск символов в нескольких строках
Что нужно сделать

Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune, а возвращает 2D-массив, где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее слово в предложении i (строку надо разбить на слова и взять последнее). То есть сигнатура следующая:
*/
func main() {
	t := "Этот мир"
	flag := isInclude(t, "это")
	fmt.Println(flag)

	text := []string{
		"Главными фишками новинки являются сквозное шифрование и отсутствие какой-либо рекламы, пишет The Verge.",
		"При этом для того, чтобы начать чат с каким-либо пользователем, необходимо узнать его номер телефона.", "Это, к слову, создаёт дополнительный уровень защиты приватности.",
	}
  chars := []rune{'v', 'g', 'e', 'ф', 'и', 'a'}
	result := parseTest(text, chars)
	fmt.Println(result)
}

func parseTest(sentences []string, chars []rune) (result [][]int) {

	result = make([][]int, len(sentences))
	for i := range result {
		result[i] = make([]int, len(chars))
	}

	for i, v := range sentences {
		temp := strings.Split(v, " ")
		text := temp[len(temp)-1]
		for j, v := range chars {
			index := strings.IndexRune(text, v)
			result[i][j] = index
		}
	}
	return
}

func isInclude(s, v string) bool {
	arr := []rune(s)
	for _, v := range v {
		flag := false
		for i, t := range arr {
			if unicode.ToUpper(t) == unicode.ToUpper(v) {
				arr[i] = 0
				flag = true
				break
			}
		}
		if flag == false {
			return false
		}
	}
	return true
}
