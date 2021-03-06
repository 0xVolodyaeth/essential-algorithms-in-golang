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

func QuickSort(arr []int, start, end int) {

	// log.Println(partition(arr, 0, len(arr)-1))

	if start < end {
		pi := partition(arr, start, end)

		QuickSort(arr, start, pi-1)
		QuickSort(arr, pi+1, end)
	}

	// return nil
}

func partition(arr []int, start, end int) int {

	pivot := arr[end]
	i := start - 1

	for j := 0; j <= end-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[end] = arr[end], arr[i+1]
	return i + 1
}
