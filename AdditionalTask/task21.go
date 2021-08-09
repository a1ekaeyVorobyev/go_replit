package main

import "fmt"

func main() {
	f()
	fmt.Println("Done")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Востановление работы. Значение i=%v в функции test при возникновение panic. \n", r)
		}
	}()
	fmt.Println("Вызов функции test.")
	test(0)
	fmt.Println("Возврат к нормальной работы функции test.")
}

func test(i int) {
	if i > 3 {
		fmt.Println("Создание Panick!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Printf("Defer в функции test, значение i=%v.\n", i)
	fmt.Printf("Значение i=%v в функции test.\n", i)
	test(i + 1)
}
