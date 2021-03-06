package main

import "log"

func main() {

	// slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	//fmt.Println(slice)

	//makeHeapFromSlice(slice)
	//fmt.Println("heap: ", slice)

	//for len(slice) != 0 {
	//fmt.Println(RemoveTopItem(&slice))
	//}

	//fmt.Println(slice)

	slice := []int{4, 3, 9, 5, 7, 6, 1, 10, 2, 8}

	QuickSort(slice, 0, len(slice)-1)

	log.Println(slice)
}

// func makeHeapFromSlice(arr []int) {

// 	for i := 0; i < len(arr); i++ {

// 		idx := i
// 		for idx != 0 {

// 			parent := (idx - 1) / 2

// 			if arr[idx] <= arr[parent] {
// 				break
// 			}

// 			temp := arr[idx]
// 			arr[idx] = arr[parent]
// 			arr[parent] = temp

// 			idx = parent
// 		}

// 	}
// }

// func RemoveTopItem(arrPtr *[]int) int {

// 	arr := *arrPtr
// 	result := arr[0]
// 	count := len(arr)

// 	arr[0] = arr[count-1]
// 	index := 0

// 	for true {
// 		child1 := 2*index + 1
// 		child2 := 2*index + 2

// 		if child1 >= count {
// 			child1 = index
// 		}

// 		if child2 >= count {
// 			child2 = index
// 		}

// 		if arr[index] >= arr[child1] && arr[index] >= arr[child2] {
// 			break
// 		}

// 		var swapChild int
// 		if arr[child1] > arr[child2] {
// 			swapChild = child1
// 		} else {
// 			swapChild = child2
// 		}

// 		arr[index], arr[swapChild] = arr[swapChild], arr[index]
// 		index = swapChild
// 	}

// 	*arrPtr = arr[:count-1]
// 	return result
// }

func QuickSort(slice []int, start, end int) {

	if start < end {
		pivot := pivotFunc(slice, start, end)

		QuickSort(slice, start, pivot-1)
		QuickSort(slice, pivot+1, end)
	}

}

func pivotFunc(slice []int, start, end int) int {

	pivot := slice[end]
	i := start - 1

	for j := start; j <= end-1; j++ {

		if slice[j] < pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[end] = slice[end], slice[i+1]
	return i + 1
}
