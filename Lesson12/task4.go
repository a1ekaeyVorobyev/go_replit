package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

/*
Урок №6 пакет ioutil

переписать задачи из урока 2 и 3 на пакет ioutil
*/

func main() {
	fileName := "task4.txt"
	fmt.Println("Программа Работа с файлами")
	Read(fileName)
	Write(fileName)
}

func Read(fileName string) {
	var result string
	var b bytes.Buffer
	i := 1
	for {
		fmt.Scan(&result)
		if strings.ToLower(result) == "выход" {
			break
		}
		dt := time.Now()
		b.WriteString(fmt.Sprintf("%d. %s %s.\n", i, dt.Format("2006-02-01 15:04:05"), result))
		i++
	}
	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		fmt.Println(err)
		return
	}
}

func Write(fileName string) {
	f, err := os.Open(fileName)
	defer f.Close()
	result, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", result)
}
