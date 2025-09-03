package main

import (
	"fmt"
	"sync"
)

// Счётчик с защитой от одновременного доступа
type Counter struct {
	mu    sync.Mutex
	value int
}

// Увеличить счётчик на 1
func (c *Counter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// Получить текущее значение
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	numGoroutines := 10

	// Запускаем 10 горутин, каждая увеличит счётчик
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(counter.Value()) // Должно быть 10
}
