package main

import "fmt"

func test3() {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%v\n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}

	fmt.Println("----------")

	a := 10
	fmt.Println("a =", a)
	ptr := &a
	*ptr = 20
	fmt.Println("a =", a)
}
