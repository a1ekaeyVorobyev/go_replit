package main

import (
	"fmt"
)

/*
Задание 3. Расчёт суммы скидки
Напишите программу, которая принимает на вход цену товара и скидку. Посчитайте и верните на экран сумму скидки.
Дополнительное условие: скидка не превышает 30% от цены товара и не больше 2 000 рублей.
*/

func main() {
	price := 0
	discount := 0
	fmt.Println("Программа Расчёт суммы скидки")
	for {
		fmt.Printf("Введите цену товара :\n")
		fmt.Scan(&price)
		if price > 0 && price < 2000 {
			break
		}
		fmt.Printf("Цена товара не может быть больше 2000. А вы ввели %v.\n", price)
		fmt.Println("Повторите ввод цены.")
	}

	for {
		fmt.Printf("Введите скидку :\n")
		fmt.Scan(&discount)
		if discount <= 30 {
			break
		}
		fmt.Printf("Скидка на товара не не превышаеть 30%. А вы ввели %v.\n", discount)
		fmt.Println("Повторите ввод скидки.")
	}

	summa := price * (100 - discount) / 100
	fmt.Println("Итоговая сумма = ", summa)

}
