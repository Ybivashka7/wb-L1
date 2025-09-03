package main

import "fmt"

func main() {
	str := "главрыба"
	fmt.Println(lineReversal(str))
}
func lineReversal(str string) string {
	// конвертируем в слайс рун
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		// меняем местами первый и последний элемент
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
