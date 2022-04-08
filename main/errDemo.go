package main

import (
	"errors"
	"fmt"
)

func errDemo() {
	errTest1()
	// 输出"panic error!"

	errTest2()
	// 输出：
	// defer inner
	// <nil>
	// test panic

	errTest3()
	// 输出"test panic"

	errTest4()
	// 输出"divided by zero"
}

func errTest1() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 将interface{}转型为具体类型
		}
	}()

	panic("panic error!")
}

func errTest2() {
	defer func() {
		fmt.Println(recover()) // 有效
	}()
	defer recover()              // 无效！
	defer fmt.Println(recover()) // 无效！
	defer func() {
		func() {
			println("defer inner")
			recover() // 无效！
		}()
	}()

	panic("test panic")
}

func errTest3() {
	defer except() // 有效
	panic("test panic")
}

func except() {
	fmt.Println(recover())
}

var ErrDivByZero = errors.New("divided by zero")

func divTest(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}

func errTest4() {
	defer func() {
		fmt.Println(recover())
	}()
	switch z, err := divTest(10, 0); err {
	case nil:
		println(z)
	case ErrDivByZero:
		panic(err)
	}
}
