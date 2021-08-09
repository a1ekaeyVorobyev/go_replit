package main

import (
	"fmt"
)

/*
4. Сумма без сдачи

Что нужно сделать

Программное обеспечение банкоматов постоянно решает задачу, как имеющимися купюрами сформировать сумму, введенную пользователем. Давайте попробуем решить похожую задачу, и определить, сможет ли пользователь заплатить за товар без сдачи или нет. Для этого он будет вводить стоимость товара и номиналы трех монет.
*/
func main() {
	var summa, cnt5, cnt2, cnt1 int
	var amount5, amount2, amount1 int
	fmt.Println("Программа Сумма без сдачи")
	fmt.Printf("Введите количество 5 монет:\n")
	fmt.Scan(&amount5)

	fmt.Printf("Введите количество 2 монет:\n")
	fmt.Scan(&amount2)

	fmt.Printf("Введите количество 1 монет:\n")
	fmt.Scan(&amount1)

	fmt.Printf("Введите сумму:\n")
	fmt.Scan(&summa)

	if summa < 0 || amount5 < 0 || amount2 < 0 || amount1 < 0 {
		fmt.Println("Все введенный значения должны быть больше 0")
		return
	}

	temp := summa
	cnt5 = temp / 5
	if amount5 > 0 && cnt5 > 0 {
		if amount5 < cnt5 {
			cnt5 = amount5
		}
		temp -= cnt5 * 5
	} else {
		cnt5 = 0
	}

	cnt2 = temp / 2
	if amount2 > 0 && cnt2 > 0 {
		if amount2 < cnt2 {
			cnt2 = amount2
		}
		temp -= cnt2 * 2
	} else {
		cnt2 = 0
	}
	cnt1 = temp / 1
	if amount1 > 0 && cnt1 > 0 {
		if amount1 < cnt1 {
			cnt1 = amount1
		}
		temp -= cnt1
	} else {
		cnt1 = 0
	}

	if summa != (cnt5*5 + cnt2*2 + cnt1) {
		fmt.Println("Пользователь не может оплатить без сдачи")
		return
	}

	fmt.Println("Пользователь может оплатить без сдачи")
	fmt.Printf("сумму: %v\n", summa)
	fmt.Println("Монетами:")
	if cnt5 > 0 {
		fmt.Printf("монетой номиналом 5 - %v шт.\n", cnt5)
	}

	if cnt2 > 0 {
		fmt.Printf("монетой номиналом 2 - %v шт.\n", cnt2)
	}
	if cnt1 > 0 {
		fmt.Printf("монетой номиналом 1 - %v шт.\n", cnt1)
	}

}
