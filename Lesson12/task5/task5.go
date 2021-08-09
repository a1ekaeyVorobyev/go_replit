package main

import (
	"fmt"
)

/*
Дополнительно задание со звездочкой*

Написать программу, которая на вход принимала бы интовое число и для него генерировала бы все возможные комбинации круглых скобок, т.е.

на вход приходит число 3

на выходе:

["((()))","(()())","(())()","()(())","()()()"]

на вход 1

на выходе:

["()"]
*/

func main() {
	var cnt int
	fmt.Println("Программа Работа с файлами")
	fmt.Println("Количество сочетаний")
	fmt.Scan(&cnt)
	arr := make([]string, 0)
	str := make([]rune, cnt*2)
	FindBracket(&arr, cnt, cnt, str, 0)
	fmt.Println(arr)
}

func FindBracket(list *[]string, left int, right int, str []rune, cnt int) {
	if left < 0 || right < left {
		return // некорректное состояние
	}
	if left == 0 && right == 0 { /* нет больше левых скобок */
		*list = append(*list, string(str))
	} else {
		/* Добавляем левую скобку, если остались любые левые скобки */
		if left > 0 {
			str[cnt] = '('
			FindBracket(list, left-1, right, str, cnt+1)
		}

		/* Добавляем правую скобку, если выражение верно */
		if right > left {
			str[cnt] = ')'
			FindBracket(list, left, right-1, str, cnt+1)
		}
	}
}
