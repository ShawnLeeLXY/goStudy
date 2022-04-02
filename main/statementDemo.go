package main

import (
	"fmt"
	"time"
)

func conditionDemo() {
	var num int
	for {
		print("请输入一个整数：")
		fmt.Scan(&num)
		if strs := []string{"大于", "小于", "等于"}; num > 0 {
			println("您输入的数", strs[0], "0！")
		} else if num < 0 {
			println("您输入的数", strs[1], "0！")
		} else {
			println("您输入的数", strs[2], "0！循环结束")
			break
		}
	}

	var grade string
	fmt.Print("请输入：")
	fmt.Scan(&grade)
	switch grade {
	case "A":
		fmt.Println("优秀！")
	case "B":
		fmt.Println("良好！")
	case "C":
		fmt.Println("及格！")
	case "D":
		fmt.Println("差！")
	default:
		fmt.Println("请输入正确的格式！")
	}
}

func cycleDemo() {
	for {
		fmt.Println("Hello!")
		time.Sleep(time.Second / 2)
	}
}

// range演示
func rangeDemo() {
	str := "abc"
	for i := range str {
		fmt.Println(string(str[i]))
	}
	for _, c := range str {
		fmt.Println(string(c))
	}

	m1 := map[int]string{1: "abc", 2: "def", 3: "ghi"}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
