package main

import "fmt"

// 输出100以下的素数
func prime() {
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j*j <= i; j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j*j > i {
			fmt.Printf("%d 是素数\n", i)
		}
	}
}
