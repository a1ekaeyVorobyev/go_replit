package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/*
Урок №2 Работа с файлами

Написать программу, которая на вход получала бы строку, введенную пользователем, а в файл писала: № строки, дата и сообщение в формате:
2020-02-10 15:00:00 продам гараж

при вводе слова “выход” производился бы выход из программы
*/

func main() {
	var result string
	f, err := os.Create("task1.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Программа Работа с файлами")
	i := 1
	for {
		fmt.Scan(&result)
		if strings.ToLower(result) == "выход" {
			break
		}
		dt := time.Now()
		f.WriteString(fmt.Sprintf("%d. %s %s.\n", i, dt.Format("2006-02-01 15:04:05"), result))
		i++
	}
}
