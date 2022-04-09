package main

import "fmt"

func methodDemo() {
	methodTest1()
	fmt.Println("----------")
	methodTest2()
	fmt.Println("----------")
	methodTest3()
}

type Data struct {
	x int
}

func (d Data) ValueTest() {
	fmt.Printf("Value: %p\n", &d)
}

func (d *Data) PointerTest() {
	fmt.Printf("Pointer: %p\n", d)
}

func (d *Data) MixTest() {
	fmt.Printf("ptr: %p, val: %v\n", d, d)
}

func methodTest1() {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)

	d.ValueTest()   // 操作副本
	d.PointerTest() // 操作指针

	p.ValueTest()   // 操作副本
	p.PointerTest() // 操作指针
}

func methodTest2() {
	data := Data{1}
	data.MixTest()
	fmt.Printf("ptr: %p, val: %v\n", &data, &data)
	num := 1
	fmt.Printf("ptr: %p, val: %v\n", &num, &num)
}

type User struct {
	id   int
	name string
}

type Manager struct {
	User
	title string
}

func (u *User) ToString() string {
	return fmt.Sprintf("User: %p, %v", u, u)
}

func (m *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", m, m)
}

// 通过匿名字段实现继承的覆盖能力
func methodTest3() {
	m := Manager{User{1, "Tom"}, "Administrator"}
	fmt.Println(m.ToString())
	fmt.Println(m.User.ToString())
}
