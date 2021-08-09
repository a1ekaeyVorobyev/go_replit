package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContains(t *testing.T) {
	str := "привет мир!"
	substr := "вет"
	result := contains(str, substr)
	require.Equal(t, result, true)

	str = "Программирование - это просто"
	substr = "вание"
	result = contains(str, substr)
	require.Equal(t, result, true)

	str = "Программирование - это просто"
	substr = "корабль"
	result = contains(str, substr)
	require.Equal(t, result, false)
}
