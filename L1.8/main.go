package main

import "fmt"

func main() {
	var num int64 = 5 //  0101
	fmt.Println("Исходное число", num)
	fmt.Println("Измененное число ", setBit(num, 0, 0)) //0100
}
func setBit(num int64, i uint, bit int) int64 {
	if bit == 1 {
		// Устанавливаем i-й бит в 1
		return num | (1 << i)
	} else {
		// Устанавливаем i-й бит в 0
		return num &^ (1 << i)
	}
}
