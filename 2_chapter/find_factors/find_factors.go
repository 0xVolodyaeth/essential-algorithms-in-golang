package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findFactors(27))
}

// complexity - O(sqrt(N))
func findFactors(num int) []int {

	factors := make([]int, 0)
	for num%2 == 0 {
		factors = append(factors, 2)
		num = num / 2
	}

	i := 3
	maxFactor := int(math.Sqrt(float64(num)))
	for {
		if i > maxFactor {
			break
		}

		for {
			if num%i != 0 {
				break
			}
			factors = append(factors, i)
			num = num / i

			maxFactor = int(math.Sqrt(float64(maxFactor)))
		}
		i = i + 2
	}

	return factors
}
