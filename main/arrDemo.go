package main

import "fmt"

// 一维数组
var arr1 [5]int = [5]int{1, 2, 3} // "[5]int"可省略
var arr2 = [5]int{1, 2, 3, 4, 5}
var arr3 = [...]int{1, 2, 3, 4, 5, 6}
var arr4 = [5]string{3: "hello world", 4: "tom"}
var arr5 = [...]struct {
	key   string
	value int32
}{
	{"key1", 123},
	3: {"key2", 666},
	{"key3", 789},
}

// 多维数组
var arr6 [3][2]int
var arr7 = [...][3]int{{1, 2, 3}, {4, 5, 6}}

func arrDemo() {
	// 一维数组
	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用索引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}
	// 多维数组
	e := [...][2]int{{1, 1}, {2, 2}, {3, 3}}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)
	fmt.Println(a, b, c, d, e)
	// len()和cap()都可返回数组长度
	fmt.Println(len(arr1), len(arr2), len(arr3), len(arr4), len(arr5))
	fmt.Println(len(a), len(b), len(c), len(d), len(e))
	fmt.Println(cap(a), cap(b), cap(c), cap(d), cap(e))

	// 多维数组遍历
	for k1, v1 := range e {
		for k2, v2 := range v1 {
			fmt.Printf("(%d, %d)=%d ", k1, k2, v2)
		}
		fmt.Println()
	}

	// 数组传参
	f := [5]int32{1, 2, 3, 4, 5}
	fmt.Println("f =", f)
	printArr(&f)
}

func printArr(arr *[5]int32) {
	arr[1] = 10
	fmt.Println("arr =", arr)
}

/*
输出结果：
[1 2 3 0 0]
[1 2 3 4 5]
[1 2 3 4 5 6]
[   hello world tom]
[{key1 123} { 0} { 0} {key2 666} {key3 789}]
[[0 0] [0 0] [0 0]]
[[1 2 3] [4 5 6]]
[1 2 0] [1 2 3 4] [0 0 100 0 200] [{user1 10} {user2 20}] [[1 1] [2 2] [3 3]]
5 5 6 5 5
3 4 5 2 3
3 4 5 2 3
(0, 0)=1 (0, 1)=1
(1, 0)=2 (1, 1)=2
(2, 0)=3 (2, 1)=3
f = [1 2 3 4 5]
arr = &[1 10 3 4 5]
*/
