package main

import (
	"flag"
	"fmt"
	"strings"
)

/*
Написать программу для нахождения подстроки в кириллической подстроке. Программа должна запускаться с помощью команды:

 go run main.go --str "строка для поиска" --substr "поиска"

Для реализации такой работы с флагами воспользуйтесь пакетом flags, а для поиска подстроки в строке вам понадобятся руны.
*/
func main() {

	str := flag.String("str", "åÅäÄ", "This's the string")
	substr := flag.String("substr", "Ää", "this's the substring")
	flag.Parse()
	if *str == "" || *substr == "" {
		fmt.Println("invalid format. You must enter a string and a substring.")
		return
	}

	flag.Parse()
	flag := strings.Contains(*str, *substr)
	fmt.Println("flag:", flag)
	flag = contains(*str, *substr)
	fmt.Println("flag:", flag)
}

func contains(str, substr string) bool {
	astr := []rune(str)
	index := 0
Exit:
	for _, v := range substr {
		for i := index; i < len(astr); i++ {
			if v == astr[i] {
				index = i + 1
				continue Exit
			}
		}
		return false
	}
	return true
}
