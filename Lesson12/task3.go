package main

import (
	"fmt"
	"os"
)

/*
Урок №4 уровни доступа

Создать файл только для чтения и проверить, что в него нельзя записать данные
*/

func main() {
	fileName := "task3.txt"
	_, err := os.Stat(fileName)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(err)
			return
		}
	} else {
		if err := os.Chmod(fileName, 0666); err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Программа уровни доступа")

	Write(fileName)

	if err := os.Chmod(fileName, 0444); err != nil {
		fmt.Println(err)
		return
	}
	Read(fileName)
	Write(fileName)

}

func Read(fileName string) {
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
}

func Write(fileName string) {
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	f.WriteString("Test string")
}
