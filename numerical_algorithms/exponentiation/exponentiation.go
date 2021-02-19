package main

import "fmt"

func main() {

	fmt.Println(RaiseToPower(7, 6))
}

func RaiseToPower(a, n int) int {

	// k := a
	for i := 1; i < n-1; i++ {
		// k = a ^ i
		fmt.Println(a * a)
	}

	return 0
}
