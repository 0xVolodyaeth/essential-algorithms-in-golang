package sorting_algorithms

func InsertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if j < i && arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	return arr
}

func SelectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		currentMin := arr[i]

		for j := i; j < len(arr); j++ {
			if currentMin > arr[j] {
				currentMin = arr[j]
				arr[j], arr[i] = arr[i], arr[j]
			}

		}
	}

	return arr
}

func BubbleSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	return arr
}

func HeapSort(arr []int) []int {

	res := make([]int, len(arr))
	makeHeapFromSlice(arr)

	for i := len(res) - 1; i != -1; i-- {
		top := RemoveTopItem(&arr)
		res[i] = top
	}

	return res
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

func MergeSort(slice []int) []int {

	if len(slice) == 1 {
		return slice
	}

	midpoint := int(len(slice) / 2)

	left := make([]int, midpoint)
	right := make([]int, len(slice)-midpoint)

	for i := 0; i < len(slice); i++ {
		if i < midpoint {
			left[i] = slice[i]
		} else {
			right[i-midpoint] = slice[i]
		}
	}

	return merge(MergeSort(left), MergeSort(right))
}

func merge(sliceLeft []int, sliceRight []int) []int {

	mergedSlice := make([]int, len(sliceLeft)+len(sliceRight))
	i := 0
	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(sliceLeft) && rightIndex < len(sliceRight) {

		if sliceLeft[leftIndex] < sliceRight[rightIndex] {
			mergedSlice[i] = sliceLeft[leftIndex]
			leftIndex++
		} else {
			mergedSlice[i] = sliceRight[rightIndex]
			rightIndex++
		}
		i++
	}

	for j := leftIndex; j < len(sliceLeft); j++ {
		mergedSlice[i] = sliceLeft[j]
		i++
	}

	for j := rightIndex; j < len(sliceRight); j++ {
		mergedSlice[i] = sliceRight[j]
		i++
	}

	return mergedSlice
}
