package main

import (
	"fmt"
	"os"
)

/*
Урок №3 интерфейс io.Reader

Написать программу, которая полностью вычитывала бы файл из предыдущего домашнего задания без использования ioutil

Подсказка: для получения размера файла у файла есть метод Stat(), который возвращает информацию о файле и ошибку.
*/

func main() {
	var fileName string

	fmt.Println("Программа интерфейс io.Reader")
	fmt.Println("Введите имя файла:")
	fmt.Scan(&fileName)
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	buff := make([]byte, fileInfo.Size())
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Read(buff)
	fmt.Printf("%s", buff)
}
