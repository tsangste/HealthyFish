package services

import (
	"sort"
)

func Calculate(items []int, quantity int) []int {
	length := len(items)
	sort.Sort(sort.Reverse(sort.IntSlice(items)))

	var result []int
	previous := items[0]
	for index, item := range items {
		if quantity == 0 {
			break
		}

		remainder := quantity % item

		if previous > quantity && quantity > item && lastValue(length, index) {
			result = append(result, previous)
		} else if remainder != quantity {
			value := (quantity - remainder) / item

			for i := 0; i < value; i++ {
				result = append(result, item)
			}

			quantity = remainder
		} else if quantity < item && lastValue(length, index) {
			result = append(result, item)
			quantity = 0
		}

		previous = item
	}

	return result
}

func lastValue(length int, index int) bool {
	return length == index+1
}
