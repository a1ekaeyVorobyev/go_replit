package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserValidation(t *testing.T) {
	text := []string{
		"Главными фишками новинки являются сквозное шифрование и отсутствие какой-либо рекламы, пишет The Verge.",
		"При этом для того, чтобы начать чат с каким-либо пользователем, необходимо узнать его номер телефона.", "Это, к слову, создаёт дополнительный уровень защиты приватности.",
	}
	chars := []rune{'v', 'g', 'e', 'ф', 'и', 'a'}
	result := parseTest(text, chars)
	etalon := [][]int{{-1, 3, 1, -1, -1, -1},
		{-1, -1, -1, 8, -1, -1},
		{-1, -1, -1, -1, 4, -1}}
	require.Equal(t, result, etalon)
}

func TestUserIsInclude(t *testing.T) {
	s := "Этот мир"
	result := isInclude(s, "Это")
	require.Equal(t, result, true)

	result = isInclude(s, "ято")
	require.Equal(t, result, false)

	result = isInclude(s, "этоэ")
	require.Equal(t, result, false)
}
