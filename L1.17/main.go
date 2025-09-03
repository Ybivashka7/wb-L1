package main

import (
	"fmt"
)

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7
	fmt.Println(binarySearch(sl, target))
}

// бинарный поиск
func binarySearch(sl []int, target int) int {

	for i, v := range sl {
		// самый левый элемент
		left := i
		// самый правый элемент элемент
		right := len(sl)
		// середина
		mid := (right - left) / 2
		// если нашли результат возвращаем индекс
		if v == target {
			return i
			// если искомое число больше данного то сдвигаем в центр в лево
		} else if v < target {
			left = mid
			// если искомое число меньше данного то сдвигаем в центр в право
		} else {
			right = mid
		}
	}
	// не нашли элемент
	return -1
}
