package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcd"
	str1 := "abCdefAaf"
	str2 := "aabcd"
	fmt.Println(uniqueStr(str))
	fmt.Println(uniqueStr(str1))
	fmt.Println(uniqueStr(str2))

}
func uniqueStr(str string) bool {
	// приводим все к одному регистру
	s := strings.ToLower(str)
	seen := make(map[rune]bool)
	// записываем в мапу и проверяем есть ли одинаковые элементы
	for _, v := range s {
		if !seen[v] {
			seen[v] = true
		} else {
			return false
		}

	}

	return true
}
