package main

import (
	"fmt"
)

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)

	makeHeapFromSlice(slice)
	fmt.Println(slice)

}

func makeHeapFromSlice(arr []int) {

	for i := 0; i < len(arr); i++ {

		idx := i
		for idx != 0 {

			parent := (idx - 1) / 2

			if arr[idx] <= arr[parent] {
				break
			}

			temp := arr[idx]
			arr[idx] = arr[parent]
			arr[parent] = temp

			idx = parent
		}

	}
}
