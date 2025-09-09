package main

import "fmt"

func main() {
	// создаем слайс
	sl := []int{10, 20, 30, 40, 50}
	// выбираем индекс
	i := 2
	// убираем ненужный индекс
	copy(sl[i:], sl[i+1:])
	// укорачиваем длину слайса
	sl = sl[:len(sl)-1]
	fmt.Println(sl)
}
