package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
доп. задание:
1) Напишите простой генератор IP-адресов.
*/

func main() {
	val := 0
	fmt.Println("Программа  Герератор IP адресов")
	rand.Seed(time.Now().UnixNano())
  fmt.Print("Введите количество сгенерированных IP адресов: ")
	fmt.Scan(&val)
	for i:=0;i<val;i++{
    fmt.Printf("%d. %d.%d.%d.%d\n",i+1,rand.Intn(253)+1,rand.Intn(253)+1,rand.Intn(253)+1,rand.Intn(253)+1)
	}
}