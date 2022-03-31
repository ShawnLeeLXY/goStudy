package main

import "fmt"

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) []student {
	//切片是引用传递，是可以改变值的
	ce[1].age = 19
	ce = append(ce, student{3, "Cindy", 20})
	return ce
}

// 引用传递测试
func test5() {
	//定义一个切片类型的结构体
	var ce []student
	ce = []student{
		{1, "Alice", 18},
		{2, "Betty", 21},
	}
	fmt.Println(ce)
	ce = demo(ce)
	fmt.Println(ce)

	m := make(map[int]student)
	m[1] = student{1, "Dora", 22}
	m[2] = student{2, "Erica", 23}
	fmt.Println(m)
	delete(m, 2)
	fmt.Println(m)
}
