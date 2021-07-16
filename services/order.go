package services

import (
	"fmt"
	"sort"
)

func Calculate(quantity int) []int {
	items := []int{250, 500, 1000, 2000, 5000}
	length := len(items)
	sort.Sort(sort.Reverse(sort.IntSlice(items)))

	var result []int
	for index, item := range items {
		if quantity == 0 {
			break
		}

		remainder := quantity % item
		if remainder != quantity {
			value := (quantity - remainder) / item

			for i := 0; i < value; i++ {
				result = append(result, item)
			}

			quantity = remainder
			continue
		}

		if quantity < item && lastValue(length, index) {
			result = append(result, item)
			quantity = 0
		}
	}

	fmt.Printf("%v\n", result)
	return result
}

func lastValue(length int, index int) bool {
	return length == index+1
}
