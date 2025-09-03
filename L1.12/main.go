package main

import "fmt"

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(uniqueStrings(str))
}

// Функция проверяем уникальность слов
func uniqueStrings(strs []string) []string {
	// для записи слов в мапу
	m := map[string]bool{}
	result := []string{}
	for _, str := range strs {
		// записываем только те элементы которых еще не было
		if !m[str] {
			m[str] = true
		}
	}
	for str := range m {
		// добаляем результирующий слайс
		result = append(result, str)
	}
	return result

}
