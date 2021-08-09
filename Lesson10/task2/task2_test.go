package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

/*
Для 1000 рублей, 1% и 1 года программа должна вывести 1126,78 и 0.04350000000022192 (возможно меньше знаков)
Для 1000 рублей, 1% и 10 лет программа должна вывести 3299.41 и 0.5216000000013992 (возможно меньше знаков
*/

func TestEPowX(t *testing.T) {

    total,bank:= interestСalculation(1000,1,1)
    require.Equal(t, total, 1126.78)
    require.

    total,bank = interestСalculation(1000,1,1)
    require.Equal(t, total, 3299.41)
}
