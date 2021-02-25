package stack

func ReverseArray(slice []int) ([]int, error) {
	st := NewStack()
	for _, i := range slice {
		st.Push(i)
	}

	for i := 0; i < len(slice); i++ {
		elem, err := st.Pop()
		if err != nil {
			return nil, err
		}
		slice[i] = elem
	}

	return slice, nil
}
