package main

import "fmt"

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

// 闭包演示
func closureDemo1() {
	c := a()
	// 依次打印1 2 3
	for i := 0; i < 3; i++ {
		c()
	}

	// 单独调用a()不会打印i
	a()
}

func x() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

// 闭包复制的是原对象指针
func closureDemo2() {
	f := x()
	fmt.Println("---")
	f()
}

// 外部引用函数参数局部变量
func closureAdd(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func closureDemo3() {
	tmp1 := closureAdd(10)
	fmt.Println(tmp1(1), tmp1(2))
	tmp2 := closureAdd(100)
	fmt.Println(tmp2(1), tmp2(2)) // 此时tmp1和tmp2不是一个实体了
	fmt.Println(tmp1(1), tmp1(2))
}

// 返回2个函数类型的返回值
func ret2Func(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	// 相加
	add := func(i int) int {
		base += i
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		return base
	}
	// 返回
	return add, sub
}

func closureDemo4() {
	f1, f2 := ret2Func(10)
	// base一直存在
	fmt.Println(f1(1), f2(2)) // 11 9
	fmt.Println(f1(3), f2(4)) // 12 8
}
