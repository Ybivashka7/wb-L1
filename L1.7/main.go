package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	m := make(map[string]int)
	var mu sync.Mutex     // для контроля записи
	var wg sync.WaitGroup // для того что бы дождаться завершения всех горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()              // блокируем запись
			defer mu.Unlock()      // разблокируем запись
			key := strconv.Itoa(i) // создаем ключ
			m[key] = i             // записываем
		}(i)
	}
	wg.Wait() // ждем завершения всех горутин
	fmt.Println(m)
}
