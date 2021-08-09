package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLemonadeChange(t *testing.T) {

	orders := []int{5, 5, 5, 10, 20}
	flag := lemonadeChange(orders)
	require.Equal(t, flag, true)

	orders = []int{10, 10}
	flag = lemonadeChange(orders)
	require.Equal(t, flag, false)

	orders = []int{5, 5, 10, 10, 20}
	flag = lemonadeChange(orders)
	require.Equal(t, flag, false)

}

func TestCheckEnterArray(t *testing.T) {

	ordersTest := []int{5, 5, 5, 10, 20}
	orderText := "5,5,5,10,20"
	order, errorMesage := checkEnterArray(orderText)
	require.Equal(t, order, ordersTest)
	require.Equal(t, errorMesage, "")

	ordersTest = []int{10, 10}
	orderText = "10,10"
	order, errorMesage = checkEnterArray(orderText)
	require.Equal(t, order, ordersTest)
	require.Equal(t, errorMesage, "")

	ordersTest = []int{10, 0}
	orderText = "10,a"
	order, errorMesage = checkEnterArray(orderText)
	require.Equal(t, order, ordersTest)
	require.Equal(t, errorMesage, "Вы не верно ввели a.\nВы не верно ввели. Значение должно быть кратно 5,10,20 a.\n")

	ordersTest = []int{0}
	orderText = "10;10"
	order, errorMesage = checkEnterArray(orderText)
	require.Equal(t, order, ordersTest)
	require.Equal(t, errorMesage, "Вы не верно ввели 10;10.\nВы не верно ввели. Значение должно быть кратно 5,10,20 10;10.\n")

}
