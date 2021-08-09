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

func TestNumberWordCapitalLetter(t *testing.T) {

	text := `Перезапущенная после атаки хакеров крупнейшая на восточном побережье США сеть топливных трубопроводов Colonial Pipeline снова испытывает трудности с сетью: клиенты не могут оформить`
	result := numberWordCapitalLetter(text)
	require.Equal(t, result, 4)

	text = `Компания Mozilla объявила о начале массового тестирования в ночных сборках и бета-выпусках Firefox режима`
	result = numberWordCapitalLetter(text)
	require.Equal(t, result, 3)

	text = `Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software`
	result = numberWordCapitalLetter(text)
	require.Equal(t, result, 5)
}
