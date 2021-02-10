package main

import "fmt"

func main() {

	var k int = 0
	for i := 0; i < 10; i++ {
		k = lgm(7, k, 5, 11)
		fmt.Println()
	}
}

func lgm(a, x, b, m int) int {
	return ((a * x) + b) % m
}
