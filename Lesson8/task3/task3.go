package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Задание 3*:
В киоске с лимонадом, каждый лимонад стоит 5 долларов.
Клиенты стоят в очереди, чтобы купить у вас, и заказывают по одному (в порядке, указанном в счетах).
Каждый покупатель может купить только один лимонад и заплатить купюрой на 5, 10 или 20 долларов. Вы должны дать каждому покупателю сдачу с его купюры.

Обратите внимание, что сначала у вас нет сдачи.
*/
func main() {
	fmt.Println("Программа ")
	var orderText string
	fmt.Println("Введите заказы через ,. например 5,10,20:")
	fmt.Scan(&orderText)

	orders,  errorMessage := checkEnterArray(orderText)

	if errorMessage !="" {
		fmt.Println("Вы не верно ввели значения заказа.")
		fmt.Println(errorMessage)
	}
	fmt.Println(lemonadeChange(orders))
}

func checkEnterArray(orderText string) (orders []int, errorMessage string) {
	split := strings.Split(orderText, ",")
	orders = make([]int, len(split))
	for i, val := range split {
		value, err := strconv.Atoi(val)
		if err != nil {
			errorMessage += fmt.Sprintf("Вы не верно ввели %v.\n", val)
		}
		if value == 5 || value == 10 || value == 20 {
			orders[i] = value
		} else {
			errorMessage += fmt.Sprintf("Вы не верно ввели. Значение должно быть кратно 5,10,20 %v.\n", val)
		}

	}
	return
}

func lemonadeChange(bills []int) bool {
	summa := 0
	cnt5 := 0
	cnt10 := 0
	cnt20 := 0

	for _, v := range bills {
		if v > 5 && summa-v < 0 {
			return false
		}
		summa += 5
		switch {
		case v == 5:
			cnt5++
		case v == 10 && cnt5*5 >= 10:
			cnt5--
			cnt10++
		case v == 20 && cnt5*5+cnt10*10 >= 20:
			cnt20++
			if cnt10 >= 2 {
				cnt10 -= 2
			} else {
				cnt5 -= (4 - cnt10*2)
        cnt10 -= cnt10
			}
		default:
			return false
		}
	}
	return true
}
