package main

import "fmt"

func main() {
	str := "sun dog snow"
	fmt.Println(reverseWords(str))
}

// разворот слова
func reverse(run []rune, left, right int) {
	for left < right {
		run[left], run[right] = run[right], run[left]
		left++
		right--
	}
}

// разворот всех слов
func reverseWords(str string) string {
	run := []rune(str)
	n := len(run)
	reverse(run, 0, n-1)
	left := 0
	for i := 0; i < n; i++ {
		if run[i] == ' ' {
			reverse(run, left, i-1)
			left = i + 1
		}
	}
	reverse(run, left, n-1)
	return string(run)
}
