package main

import (
	"fmt"
)

// поиск пересечений
func intersection(a, b []int) []int {
	m := make(map[int]struct{})
	var result []int

	// записываем все элементы из а в мапу
	for _, v := range a {
		m[v] = struct{}{}
	}

	// проверяем элементы из b
	for _, v := range b {
		if _, ok := m[v]; ok {
			result = append(result, v)
			delete(m, v) //удаляем дубликаты
		}
	}

	return result
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	fmt.Println(intersection(A, B))

}
