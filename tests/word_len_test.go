package tests

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestWordLength(t *testing.T) {
	str := "我是 Leo!"
	fmt.Println("str", str, len(str), utf8.RuneCountInString(str))
}
