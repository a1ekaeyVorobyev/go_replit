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

func TestNumberWordToInt(t *testing.T) {

	text := `a10 10 20b 20 30c30 30 dd`
	result := numberWordToInt(text)
	require.Equal(t, result, 3)

	text = `13 ff 7f7 98`
	result = numberWordToInt(text)
	require.Equal(t, result, 2)
}
