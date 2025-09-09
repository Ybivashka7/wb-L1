package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаём большие числа
	a := big.NewInt(0).SetUint64(1_500_000)
	b := big.NewInt(0).SetUint64(2_000_000)

	// Результат
	sum := new(big.Int).Add(a, b)  // a + b
	diff := new(big.Int).Sub(a, b) // a - b
	prod := new(big.Int).Mul(a, b) // a * b
	quot := new(big.Int).Div(a, b) // a / b

	// Вывод
	fmt.Println("a =", a, "b =", b)
	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", diff)
	fmt.Println("Произведение:", prod)
	fmt.Println("Деление:", quot)
}
