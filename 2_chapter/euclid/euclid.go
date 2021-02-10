package main

import "fmt"

func main() {
	a, b := 60, 24
	fmt.Println(gdm(a, b)) // 12
}

// O(log(n))
func gdm(a, n int) int {
	for n != 0 {
		rem := a % n
		a = n
		n = rem
	}

	return a
}
