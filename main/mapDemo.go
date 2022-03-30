package main

import "fmt"

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
