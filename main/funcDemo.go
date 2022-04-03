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

func add(a, b int) (c int) {
	c = a + b
	return
}

func cal(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2

	return
}

func retDemo1() {
	var a, b = 1, 2
	c := add(a, b)
	sum, avg := cal(a, b)
	fmt.Println(a, b, c, sum, avg)
}

func op(x, y int) (z int) {
	defer func() {
		println(z) // 输出: 203
	}()

	z = x + y
	return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
}

func retDemo2() {
	println("z =", op(1, 2)) // 输出: z = 203
}
