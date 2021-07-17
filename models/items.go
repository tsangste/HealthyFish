package models

import "sort"

var items []int

func init() {
	if items == nil {
		items = []int{250, 500, 1000, 2000, 5000}
	}
}

func GetAll() []int {
	sort.Ints(items)

	return items
}

func Add(item int) {
	items = append(items, item)
}

func Delete(item int) {
	found := indexOf(items, item)
	if found == -1 {
		return
	}

	items = append(items[:found], items[found+1:]...)
}

func indexOf(data []int, element int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
