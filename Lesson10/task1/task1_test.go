package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
для x = 0 и n = 1 вывод должен быть 1
для x = 1 и n = 3 вывод должен быть 2.7182539682539684
для x = 1 и n = 5 вывод должен быть 2.7182815255731922
*/

func TestEPowX(t *testing.T) {

	result := ePowX(0, 1)
	require.Equal(t, result, 1.0)

	result = ePowX(1, 3)
	require.Equal(t, result, 2.7182539682539684)

	result = ePowX(1, 5)
	require.Equal(t, result, 2.7182815255731922)

}
