package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
Проверить на примерах:
1 1 результат uint8
1 -1 результат int8
640 100 результат uint16
-640 100 результат int32
3000 3000 результат uint32
-3000 3000 результат int32
*/

func TestTypeMin(t *testing.T) {

	var summa int64
	summa = 1 * 1
	text := typeMin(summa)
	require.Equal(t, text, "uint8")

	summa = -1 * 1
	text = typeMin(summa)
	require.Equal(t, text, "int8")

	summa = 640 * 100
	text = typeMin(summa)
	require.Equal(t, text, "uint16")

	summa = -640 * 100
	text = typeMin(summa)
	require.Equal(t, text, "int32")

	summa = 3000 * 3000
	text = typeMin(summa)
	require.Equal(t, text, "uint32")

	summa = -3000 * 3000
	text = typeMin(summa)
	require.Equal(t, text, "int32")
}
