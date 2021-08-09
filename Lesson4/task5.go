package main

import (
	"fmt"
)

/*
Задача 5. Ресторан

В ресторане действуют следующие правила:

по понедельникам должна применяться скидка 10% на всё меню, потому что понедельник — день тяжёлый!
по пятницам, если сумма чека превышает 10 000 руб., включается дополнительная скидка в размере 5%;
если число гостей в одной компании превышает 5 человек, автоматически включается надбавка на обслуживание 10%.
Напишите программу, которая запрашивает день недели, число гостей и сумму чека и рассчитывает сумму к оплате.
*/
func main() {
	var dayOfWeek, numberOfGuests, amountOfCheck int
	discount := 0

	fmt.Println("Программа Ресторан")
	fmt.Printf("Введите день недели( понедельник =1,вторник=2....воскресенье=7):\n")
	fmt.Scan(&dayOfWeek)

	fmt.Printf("Введите число гостей:\n")
	fmt.Scan(&numberOfGuests)

	fmt.Printf("сумму чека:\n")
	fmt.Scan(&amountOfCheck)

	if dayOfWeek < 1 || dayOfWeek > 7 {
		fmt.Printf("Вы не правельно ввели день недели. Должно быть от 1 до 7. Вы ввели %v:\n", dayOfWeek)
		return
	}
	if numberOfGuests < 1 {
		fmt.Printf("Вы не правельно ввели количество госте. Должен быть хоть 1 человек, а вы ввели %v:\n", numberOfGuests)
		return
	}

	if amountOfCheck < 1 {
		fmt.Printf("Неужели вы ничего не заказали?")
		return
	}

	if dayOfWeek == 1 {
		discount += 10
	}

	if amountOfCheck > 10000 {
		discount += 5
	}

	if numberOfGuests > 5 {
		discount -= 5
	}
	summaDiscount := amountOfCheck / 100 * discount

	if discount > 0 {
		fmt.Printf("Сумма чека %v, скидка %v. Итого: %v \n", amountOfCheck, summaDiscount, amountOfCheck-summaDiscount)
	} else if discount == 0 {
		fmt.Printf("Сумма чека %v, К сожлению Вам не положенна скидка. Итого: %v \n", amountOfCheck, amountOfCheck)
	} else {
		fmt.Printf("Сумма чека %v, надбавка на обслуживание %v. Итого: %v \n", amountOfCheck, summaDiscount, amountOfCheck+summaDiscount)
	}
}
