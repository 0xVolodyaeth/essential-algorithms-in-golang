package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(isPrime(27, 4))
	fmt.Println(isPrime(18, 3))
	fmt.Println(isPrime(17, 1))
	fmt.Println(isPrime(18, 1))
	fmt.Println(isPrime(12, 10))
	fmt.Println(isPrime(7, 10))
	fmt.Println(isPrime(3, 10))
}

func isPrime(p, maxTest int) bool {
	for test := 0; test < maxTest; test++ {
		seed := rand.NewSource(time.Now().Unix())
		r := rand.New(seed)

		var randomNum int
		for {
			if randomNum > 1 && randomNum < p {
				break
			}
			randomNum = r.Intn(p)
		}

		pow := int(math.Pow(float64(randomNum), float64(p-1)))
		mod := pow % p
		if mod != 1 {
			return false
		}
	}
	return true
}
