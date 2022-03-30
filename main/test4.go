package main

import "fmt"

func test4() {
	var a byte
	var b rune
	fmt.Printf("type of a: %T\n", a)
	fmt.Printf("type of b: %T\n", b)

	var c struct {
		name string
		age  uint8
	}
	c.name = "Alice"
	c.age = 18
	fmt.Printf("%v\n", c)
}
