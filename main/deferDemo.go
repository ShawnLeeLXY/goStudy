package main

import "fmt"

func deferDemo() {
	var whatever [5]struct{}

	for i := range whatever {
		defer fmt.Println(i)
	}
}
