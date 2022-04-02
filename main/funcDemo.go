package main

import "fmt"

func funcTest(fn func() int) int {
	return fn()
}

// 定义函数类型
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func funcDemo() {
	// 直接将匿名函数当参数
	s1 := funcTest(func() int { return 100 })

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

	println(s1, s2)
}

// 变参本质上是一个slice
func argsHandle(str string, nums ...int) string {
	var sum int
	for _, i := range nums {
		sum += i
	}
	return fmt.Sprintf(str, sum)
}

func argsDemo() {
	fmt.Println(argsHandle("sum: %d", 1, 2, 3))
}
