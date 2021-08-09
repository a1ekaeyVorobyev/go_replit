
package main

import (
	"fmt"
)
/*
Задание 2: Проблемы округления процентов

Что нужно сделать

В связи с особенностями хранения дробных чисел, они не очень подходят для хранения денежных значений (может потеряться значимая часть при переполнении, да и дробная часть больше двух знаков не нужна). Но мы с вами попробуем решить задачу начисления процентов используя именно их. Пусть пользователь вводит сумму которую он кладет в банк, ежемесячный процент капитализации и количество лет на которое делается вклад. Необходимо помесячно пересчитывать сумму округляя ее до целого количества копеек вниз. Например, если после начисления процентов остаток включает 35,78 копеек, то оставляем только 35 копеек, а дробную часть отбрасываем. По окончании работы программы необходимо вывести какую сумму получит вкладчик на руки и какая сумма будет зачислена в пользу банка за счет округления копеек.

Для 1000 рублей, 1% и 1 года программа должна вывести 1126,78 и 0.04350000000022192 (возможно меньше знаков)
Для 1000 рублей, 1% и 10 лет программа должна вывести 3299.41 и 0.5216000000013992 (возможно меньше знаков)
*/

func main() {
  var cntYear,interestOnDeposit int
  var amount float64 = 0.0
  fmt.Println("Программа проблемы округления процентов")
  fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)
  fmt.Print("Введите количество лет: ")
	fmt.Scan(&cntYear)
  fmt.Print("Введите процент по депазиту: ")
	fmt.Scan(&interestOnDeposit)

  total,bank:= interestСalculation(amount,cntYear,interestOnDeposit)

  fmt.Printf("amount = %6.2f\n",total)
  fmt.Printf("bank = %2.20f\n",bank)
}

func  interestСalculation(amount float64, cntYear,interestOnDeposit int ) (total,bank float64){
  total = amount
  for i:=0;i<cntYear*12;i++{
    p := total/100*float64(interestOnDeposit)
    c := float64(int(p*100))/100
    total += c
    bank += p - c
  }
  return
}