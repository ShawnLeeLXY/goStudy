package main

import "fmt"

var array = [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90}

func sliceDemo1() {
	// 声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("空")
	} else {
		fmt.Println("非空")
	}
	s2 := []int{}
	// make([]type, len)
	// make([]type, len, cap)
	var s3 = make([]int, 0)
	var s4 = make([]int, 0, 0)
	fmt.Println(s1, s2, s3, s4)
	s5 := []int{1, 2, 3} // 注意没有省略号
	fmt.Println(s5)
	// 从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	// 前包后不包
	s6 = arr[1:4]
	fmt.Println(s6)
	// 切片的len=high-low, cap=max-low
	fmt.Println("len =", len(s6), "cap =", cap(s6))

	fmt.Println("----------")

	fmt.Println(array[3:6])
	fmt.Println(array[:6])
	fmt.Println(array[3:])
	fmt.Println(array[:])
	fmt.Println(array[:len(array)-1])

	fmt.Println("----------")

	// 切片是数组的引用
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s7 := data[2:5]
	s7[0] += 100
	s7[1] += 200
	p := &s7[2]
	*p += 300
	fmt.Println(s7)
	fmt.Println(data)

	fmt.Println("----------")

	// 通过初始化表达式构造，可使用索引号
	s8 := []int{0, 1, 2, 3, 8: 100}
	fmt.Println(s8, len(s8), cap(s8))
	// 使用 make 创建，指定 len 和 cap 值
	s9 := make([]int, 6, 8)
	fmt.Println(s9, len(s9), cap(s9))
	// 省略 cap，相当于 cap = len
	s10 := make([]int, 6)
	fmt.Println(s10, len(s10), cap(s10))

	fmt.Println("----------")

	arrs := [5]struct {
		x int
	}{}
	s := arrs[:]
	arrs[1].x = 10
	s[2].x = 20
	fmt.Println(arrs)
	fmt.Printf("%p, %p\n", &arrs, &arrs[0])

}

/*
输出结果：
空
[] [] [] []
[1 2 3]
[2 3 4]
len = 3 cap = 4
----------
[40 50 60]
[10 20 30 40 50 60]
[40 50 60 70 80 90]
[10 20 30 40 50 60 70 80 90]
[10 20 30 40 50 60 70 80]
----------
[102 203 304]
[0 1 102 203 304 5 6 7 8 9]
----------
[0 1 2 3 0 0 0 0 100] 9 9
[0 0 0 0 0 0] 6 8
[0 0 0 0 0 0] 6 6
----------
[{0} {10} {20} {0} {0}]
0xc00000a540, 0xc00000a540
*/

func sliceDemo2() {
	// append: 向 slice 尾部添加数据
	// 若追加元素后len<=cap，则地址不变
	// 若追加元素后len>cap，则重新开辟2*cap的空间去存储新slice
	a := []int{1, 2, 3}
	fmt.Printf("slice a : %v\n", a)
	b := []int{4, 5, 6}
	fmt.Printf("slice b : %v\n", b)
	c := append(a, b...)
	fmt.Printf("slice c : %v\n", c)
	d := append(c, 7)
	fmt.Printf("slice d : %v\n", d)
	e := append(d, 8, 9, 10)
	fmt.Printf("slice e : %v\n", e)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c)
	fmt.Printf("%p\n", &d)
	fmt.Printf("%p\n", &e)

	fmt.Println("----------")

	// 复制slice
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("arr =", arr)
	s1 := arr[:5]
	s2 := arr[8:]
	fmt.Printf("s1 = %v\n", s1)
	fmt.Printf("s2 = %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied s1 = %v\n", s1)
	fmt.Printf("copied s2 = %v\n", s2)
	fmt.Println("copied arr =", arr)

	fmt.Println("----------")

	str := "hello world"
	s3 := []byte(str)
	fmt.Println("s3 =", string(s3))
	fmt.Printf("str: %p, s3: %p\n", &str, s3)
	s3 = s3[:8]
	s3[6] = 'G'
	fmt.Println("s3 =", string(s3))
	fmt.Printf("str: %p, s3: %p\n", &str, s3)
	s3 = append(s3, '!')
	fmt.Println("s3 =", string(s3))
	fmt.Printf("str: %p, s3: %p\n", &str, s3)
}

/*
输出结果：
slice a : [1 2 3]
slice b : [4 5 6]
slice c : [1 2 3 4 5 6]
slice d : [1 2 3 4 5 6 7]
slice e : [1 2 3 4 5 6 7 8 9 10]
0xc000004078
0xc0000040a8
0xc0000040d8
0xc000004108
0xc000004138
----------
arr = [0 1 2 3 4 5 6 7 8 9]
s1 = [0 1 2 3 4]
s2 = [8 9]
copied s1 = [0 1 2 3 4]
copied s2 = [0 1]
copied arr = [0 1 2 3 4 5 6 7 0 1]
----------
s3 = hello world
str: 0xc00004c250, s3: 0xc00001a220
s3 = hello Go
str: 0xc00004c250, s3: 0xc00001a220
s3 = hello Go!
str: 0xc00004c250, s3: 0xc00001a220
*/

func sliceDemo3() {
	arr := [3]int{100, 200, 300}
	fmt.Printf("arr: %p, %v\n", &arr, arr)
	slc := arr[:]
	fmt.Printf("slc: %p, %v\n", &slc, slc)
}
