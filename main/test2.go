package main

import "fmt"

func test2() {
	s1 := "hello"
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "一个字符串"
	runeS2 := []rune(s2)
	runeS2[0] = '两'
	fmt.Println(string(runeS2))
}
