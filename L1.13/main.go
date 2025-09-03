package main

import "fmt"

func main() {
	a := 5
	b := 10

	fmt.Println(xor(a, b))
	fmt.Println(rearrange(a, b))
	fmt.Println(mathSwap(a, b))
}

// xor выполняет обмен значениями без использования временной переменной
func xor(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

// rearrange меняет местами значения с помощью множественного присваивания
func rearrange(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

// mathSwap меняет местами значения через арифметику
func mathSwap(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}
