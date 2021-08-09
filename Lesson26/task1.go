package main

import (
	"fmt"
	"os"
	"strings"
)

/*
Программа должна получать на вход имена двух файлов и конкатенировать их, используя strings.Join.
go run task1.go file1.txt file2.txt result.txt
*/
const separetion = "\n"

func main() {
	var str string
	var fileResult, fileName2, fileName1 string
	switch len(os.Args) {
	case 4:
		fileResult = os.Args[3]
		if _, err := os.Stat(fileResult); err == nil {
			var x string
			fmt.Print("The file exists. Rewrite?(y/n)")
			fmt.Scan(&x)
			if strings.ToUpper(x) != "Y" {
				return
			}
		}
		fallthrough
	case 3:
		fileName2 = os.Args[2]
		if err := check(fileName2); err != nil {
			fmt.Println(err)
			return
		}
		str = Read(fileName2)
		fallthrough
	case 2:
		fileName1 = os.Args[1]
		if err := check(fileName1); err != nil {
			fmt.Println(err)
			return
		}
		str = strings.Join([]string{Read(fileName1), str}, separetion)
		if len(os.Args) == 4 {
			Write(fileResult, str)
		} else {
			fmt.Println(str)
		}
	default:
		fmt.Println("Invalid format. You must enter the name of file1, the name of file2, and the name of the file to write, if necessary.")
	}
}

func Read(fileName string) string {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	buff := make([]byte, fileInfo.Size())
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	f.Read(buff)
	return string(buff)
}

func Write(fileName, text string) {
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	f.WriteString(text)
	//Записываем на диск, так как может находится в кэше памяти
	if err = f.Sync(); err != nil {
		fmt.Println(err)
		return
	}
}

func check(fileName string) error {
	_, err := os.Stat(fileName)
	if err != nil {

		return err
	}
	return nil
}
