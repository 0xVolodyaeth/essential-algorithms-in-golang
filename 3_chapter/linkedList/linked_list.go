package main

import "fmt"

type ListItem struct {
	Value    *int
	NextItem *ListItem
}

func (top *ListItem) Iterate() {

	for top != nil {
		if top.Value == nil {
			fmt.Println(top.Value, top.NextItem)
			top = top.NextItem
			continue
		}
		fmt.Println(*top.Value, top.NextItem)
		top = top.NextItem
	}
}

func (li *ListItem) FindCell(value int) *ListItem {

	for li != nil {
		if li.Value == nil {
			li = li.NextItem
			continue
		}
		if *li.Value == value {
			return li
		}

		li = li.NextItem
	}
	return nil
}

func (top *ListItem) FindCellBefore(target int) *ListItem {

	for top.NextItem != nil {
		if *top.NextItem.Value == target {
			return top
		}
		top = top.NextItem
	}

	return nil
}

func (top *ListItem) AddAtBeginning(newListItem *ListItem) {
	newListItem.NextItem = top.NextItem
	top.NextItem = newListItem
}

func (top *ListItem) AddAtEnd(newCell *ListItem) {
	for top.NextItem != nil {
		top = top.NextItem
	}
	top.NextItem = newCell
	return
}

func (top *ListItem) InsertCell(afterListItem *ListItem, newListItem *ListItem) {
	newListItem.NextItem = afterListItem.NextItem
	afterListItem.NextItem = newListItem
}

func (top *ListItem) DeleteAfter(afterMe *ListItem) {
	afterMe.NextItem = afterMe.NextItem.NextItem
}

func (top)

func newInt(a int) *int {
	return &a
}


func main() {

	forth := ListItem{newInt(4), nil}
	third := ListItem{newInt(3), &forth}
	second := ListItem{newInt(2), &third}
	first := ListItem{newInt(1), &second}

	// ограничитель, к которому обращаются все алгоритмы
	top := ListItem{nil, &first}

	top.Iterate()
	fmt.Println("=====")

	fmt.Println(top.FindCell(3))
	fmt.Println("===== FindCellBefore")

	fmt.Println(top.FindCellBefore(1))
	fmt.Println("===== AddAtBeginning")

	top.AddAtBeginning(&ListItem{newInt(5), nil})
	top.Iterate()
	fmt.Println("===== AddAtEnd")

	top.AddAtEnd(&ListItem{newInt(10), nil})
	top.Iterate()

	fmt.Println("===== InsertCell")
	top.InsertCell(&third, &ListItem{newInt(18), nil})
	top.Iterate()

	fmt.Println("===== DeleteAfter")
	top.DeleteAfter(&third)
	top.Iterate()

}
