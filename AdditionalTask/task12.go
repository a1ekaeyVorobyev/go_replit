package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

/*
Урок №6 пакет ioutil

Доп. задание: попробуйте передавать путь к файлу через аргумент, c которым запускаете программу. Подсказка - можно это сделать с помощью пакета flag)

go run task12.go -filename=task12.txt -path=./task11
*/

func main() {
	var fileNames string
	fileName := flag.String("filename", "", "file name")
	path := flag.String("path", "", "file path")

	flag.Parse()
	if *path != "" {
		if _, err := os.Stat(*path); os.IsNotExist(err) {
			fmt.Println("err", err)
			fmt.Println("Вы не ввели путь к файлу")
			return
		}
	}
	if *fileName == "" {
		fmt.Println("Вы не ввели имя файла")
		return
	}
	if string(*path)[0] != '/' || string(*path)[0] != '.' {
		fileNames = `./` + *path + `/` + *fileName
	} else {
		fileNames = *path + `/` + *fileName
	}
	fmt.Println("Программа Работа с файлами")
	Write(fileNames)
	Read(fileNames)
}

func Write(fileName string) {
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

func Read(fileName string) {
	f, err := os.Open(fileName)
	defer f.Close()
	result, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", result)
}
