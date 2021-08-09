package main

import (
	"fmt"
)

/*
Задание 6*. Движение лифта
В доме 24 этажа. Лифт должен ходить по кругу, пока не доставит всех пассажиров на первый этаж. Три пассажира ждут на четвёртом, седьмом и десятом этажах. При движении вверх лифт не должен останавливаться, при движении вниз — собирать всех, но не более двух человек в лифте. При этом лифт каждый раз доезжает до самого верхнего этажа и только после этого начинает движение вниз. Напишите программу, которая доставит всех пассажиров на первый этаж.
*/

func main() {
	amountRepeat := 0
  numberOfPeopleElevator := 0

	fmt.Println("Программа  Движение лифта")
	amountFloors := 24
	floorNumber := 24
	arr := make([]int, amountFloors+1)
	arr[4] = 3
	arr[7] = 3
	arr[10] = 3
	
  for {
		if floorNumber == 1 {
			amountRepeat += 1
			fmt.Println("Этаж 1. Все пассажиры вышли.")
			if numberOfPeopleElevator < 2 {
				break
			}
			numberOfPeopleElevator = 0
			floorNumber = amountFloors
			continue
		}
	
  	if numberOfPeopleElevator == 2 {
			fmt.Printf("Этаж %v. Лифт заполнен едем вниз.\n", floorNumber)
			floorNumber -= 1
			continue
		}
	
  	if arr[floorNumber] == 0 {
			fmt.Printf("На этаже %v никого нет. Везем %v\n", floorNumber, numberOfPeopleElevator)
			floorNumber -= 1
			continue
		}
	
  	if arr[floorNumber] > 0 {
			for numberOfPeopleElevator < 2 && arr[floorNumber] != 0 {
				numberOfPeopleElevator += 1
				arr[floorNumber] -= 1
			}
			continue
		}
	}

	fmt.Printf("Мы спустили всех людей. Лифт спустился %v раз. \n", amountRepeat)
}
