package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindBracket(t *testing.T) {
	//1 ["()"]
	cnt := 1
	arr := make([]string, 0)
	str := make([]rune, cnt*2)

	FindBracket(&arr, cnt, cnt, str, 0)
	test := []string{"()"}
	require.Equal(t, arr, test)

	//2
	cnt = 2
	arr = make([]string, 0)
	str = make([]rune, cnt*2)

	FindBracket(&arr, cnt, cnt, str, 0)
	test = []string{"(())", "()()"}
	require.Equal(t, arr, test)

	//3 ["((()))","(()())","(())()","()(())","()()()"]
	cnt = 3
	arr = make([]string, 0)
	str = make([]rune, cnt*2)

	FindBracket(&arr, cnt, cnt, str, 0)
	test = []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	require.Equal(t, arr, test)

}
