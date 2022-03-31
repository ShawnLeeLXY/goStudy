package main

import (
	"fmt"
	"sort"
)

func mapDemo() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	key := "Class"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "Alice", "Betty", "Cindy")
	sliceMap[key] = value
	fmt.Println("after init...")
	fmt.Println(sliceMap)
}

// 实现map有序输出
func mapOutput() {
	map1 := make(map[int]string, 5)
	map1[3] = "Jenny"
	map1[1] = "Alice"
	map1[4] = "Nancy"
	map1[5] = "Tina"
	map1[2] = "Claire"
	var s1 []int
	for k := range map1 {
		s1 = append(s1, k)
	}
	sort.Ints(s1)
	for i := 0; i < len(map1); i++ {
		fmt.Println(map1[s1[i]])
	}
}
