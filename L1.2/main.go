package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	// wg - для того что бы дождаться завершения всех горутин
	wg := sync.WaitGroup{}

	for _, num := range nums {
		// увеличиваем счетчик горутин
		wg.Add(1)
		go func(num int) {
			//уменьшаем счетчик горутин
			defer wg.Done()
			fmt.Println(num * num) // возводим в квадрат
		}(num) // передаём значение, чтобы не было замыкания
	}
	//ждем завершения горутин
	wg.Wait()

}
