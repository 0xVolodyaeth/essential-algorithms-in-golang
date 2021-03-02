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

	heap := makeHeap(arr)
	return heap
}

func makeHeap(arr []int) []int {

	// проходимся по всему массиву
	for i := 0; i < len(arr); i++ {

		idx := i

		// проходимся до корня
		for idx != 0 {
			// родитель узла c индексом idx
			parent := (idx - 1) / 2

			// потомок меньше родителя - выходим
			if arr[idx] <= arr[parent] {
				break
			}

			// потомок больше - свопаем
			arr[idx], arr[parent] = arr[parent], arr[idx]
			idx = parent
		}
	}
	return arr
}
