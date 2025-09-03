package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// выбираем опорный элемент
	pivot := arr[len(arr)/2]

	var left, right []int
	for i, v := range arr {
		if i == len(arr)/2 {
			continue // пропускаем опорный элемент
		}
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	// рекурсия
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	data := []int{10, -3, 5, 0, 8, 2, 1, 9}
	fmt.Println("Было:", data)

	sorted := quickSort(data)
	fmt.Println("Стало:", sorted)
}
