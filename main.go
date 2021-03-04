package main

import (
	"fmt"
)

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	makeHeapFromSlice(slice)
	fmt.Println(slice)

	fmt.Println(RemoveTopItem(&slice))
	fmt.Println(slice)

	fmt.Println(RemoveTopItem(&slice))
	fmt.Println(slice)

	fmt.Println(RemoveTopItem(&slice))
	fmt.Println(slice)

	fmt.Println(RemoveTopItem(&slice))
	fmt.Println(slice)

	fmt.Println(RemoveTopItem(&slice))
	fmt.Println(slice)

	fmt.Println(RemoveTopItem(&slice))
	fmt.Println(slice)

	fmt.Println(RemoveTopItem(&slice))
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

func RemoveTopItem(arrPtr *[]int) int {

	arr := *arrPtr
	result := arr[0]
	count := len(arr)

	arr[0] = arr[count-1]
	index := 0

	for true {
		child1 := 2*index + 1
		child2 := 2*index + 2

		if child1 >= count {
			child1 = index
		}

		if child2 >= count {
			child2 = index
		}

		if arr[index] >= arr[child1] && arr[index] >= arr[child2] {
			break
		}

		var swapChild int
		if arr[child1] > arr[child2] {
			swapChild = child1
		} else {
			swapChild = child2
		}

		arr[index], arr[swapChild] = arr[swapChild], arr[index]
		index = swapChild
	}

	*arrPtr = arr[:count-1]
	return result
}
