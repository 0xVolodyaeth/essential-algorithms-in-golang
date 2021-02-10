package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findPrimes(14)) // [4 6 8 9 10 12 14]
}

// complexity - O(N x log(log(N)))
func findPrimes(maxNum int) []int {
	isComposite := make([]bool, maxNum+1)
	for i := 4; i <= maxNum; i = i + 2 {
		isComposite[i] = true
	}

	nextPrime := 3
	stopAt := int(math.Sqrt(float64(maxNum)))
	for nextPrime <= stopAt {
		for i := nextPrime * 2; i < maxNum; i = i + nextPrime {
			isComposite[i] = true
		}

		nextPrime = nextPrime + 2
		for nextPrime <= maxNum && isComposite[nextPrime] {
			nextPrime = nextPrime + 2
		}

	}

	primes := make([]int, 0)
	for i := 2; i < maxNum+1; i++ {
		if isComposite[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
