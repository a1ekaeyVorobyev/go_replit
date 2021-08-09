package main

import (
	"fmt"
	"math"
)

func main() {
	var numberOfStops, comeBus, exitBus, total, cnt int
	var profit, revenue float64
	ticket := 20
	fmt.Println("Программа Симулятор маршрутки")
	fmt.Println("Введите количество остановок:")
	fmt.Scan(&numberOfStops)
	numberOfStops = 4
	for i := 1; i <= numberOfStops; i++ {
		fmt.Println("Сколько пассажиров зашло на остановке?")
		fmt.Scan(&comeBus)
		fmt.Println("Сколько пассажиров вышло на остановке?")
		fmt.Scan(&exitBus)
		cnt += comeBus - exitBus
		total += comeBus
		fmt.Printf("Отправляемся с %v остановки. В салоне пассажиров: %v \n", i, cnt)
	}

	revenue = (float64)(total * ticket)
	profit = revenue - math.Round(revenue*0.25*1000)/1000 - 3*math.Round(revenue*0.2*1000)/1000
	fmt.Printf("Всего заработали: %v руб.\n", revenue)
	fmt.Printf("Зарплата водителя: %v руб.\n", math.Round(revenue*0.25*1000)/1000)
	fmt.Printf("Расходы на топливо: %v руб.\n", math.Round(revenue*0.2*1000)/1000)
	fmt.Printf("Налоги: %v руб.\n", math.Round(revenue*0.2*1000)/1000)
	fmt.Printf("Расходы на ремонт машины: %v руб.\n", math.Round(revenue*0.2*1000)/1000)
	fmt.Print("-------")
	fmt.Printf("Итого доход: %v руб.\n", profit)
}
