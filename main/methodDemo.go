package main

import "fmt"

type Data struct {
	x int
}

func (d Data) ValueTest() {
	fmt.Printf("Value: %p\n", &d)
}

func (d *Data) PointerTest() {
	fmt.Printf("Pointer: %p\n", d)
}

func methodDemo() {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)

	d.ValueTest()   // 操作副本
	d.PointerTest() // 操作指针

	p.ValueTest()   // 操作副本
	p.PointerTest() // 操作指针
}
