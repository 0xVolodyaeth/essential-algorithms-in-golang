package heap

type Heap struct {
	slice *[]int
}

func NewHeap(slice *[]int) *Heap {
	h := &Heap{slice}
	h.makeHeap()
	return h
}

func (h *Heap) Len() int {
	return len(*h.slice)
}

func (h *Heap) makeHeap() {

	arr := *h.slice
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

func (h *Heap) RemoveTopItem() int {
	arr := *h.slice
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

	*h.slice = arr[:count-1]
	return result
}
