package main

import "fmt"

func interfaceDemo() {
	interfaceTest1()
}

type speaker interface {
	speak()
}

type chinese struct{}

type british struct{}

func (c chinese) speak() {
	fmt.Println("Chinese speak Chinese.")
}

func (b british) speak() {
	fmt.Println("British speak English.")
}

func interfaceTest1() {
	var zhangsan speaker = chinese{}
	var jack speaker = british{}
	zhangsan.speak()
	jack.speak()
}
