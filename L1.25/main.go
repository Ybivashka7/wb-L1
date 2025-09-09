package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Начало")
	// пауза 2 секунды
	sleep(2 * time.Second)
	fmt.Println("Конец")
}

// Sleep блокирует выполнение текущей горутины
func sleep(duration time.Duration) {
	start := time.Now()
	for {
		if time.Since(start) > duration {
			break
		}
	}
}
