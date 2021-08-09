package main

import (
	"fmt"
)

/*
Задание 5*. Задача на постепенное наполнение корзин разной ёмкости (if и continue и break)
Представьте, что у вас есть три корзины разной ёмкости. Пользователю предлагается ввести, какое количество яблок помещается в каждую корзину. После этого программа должна заполнить все корзины по очереди, учитывая, какие корзины уже заполнены, строго соблюдая очерёдность заполнения и добавляя по одному яблоку в каждой итерации.
Пример: пользователь решил, что у корзин будет ёмкость на шесть, четыре и девять яблок. Программа должна заполнить корзину 1 и в следующей итерации перейти к корзине 2, далее в следующей итерации перейти к корзине 3. Если очередная корзина уже заполнена, программа должна переходить к следующей по очереди, и так по кругу, пока не заполнит все.
*/

func main() {
	var basketСapacity1, basketСapacity2, basketСapacity3 int
	fmt.Println("Программа  Задача на постепенное наполнение корзин разной ёмкости")

	fmt.Printf("Введите ёмкость 1 корзины:\n")
	fmt.Scan(&basketСapacity1)
	fmt.Printf("Введите ёмкость 2 корзины:\n")
	fmt.Scan(&basketСapacity2)
	fmt.Printf("Введите ёмкость 3 корзины:\n")
	fmt.Scan(&basketСapacity3)

	i := 0
	for {
		i += 1
		if basketСapacity1 > 0 {
			basketСapacity1 -= 1
			continue
		}
		if basketСapacity2 > 0 {
			basketСapacity2 -= 1
			continue
		}
		if basketСapacity3 > 0 {
			basketСapacity3 -= 1
			continue
		}
		break
	}
	fmt.Printf("Мы заполнили карзины за %v итераций.\n", i)
}