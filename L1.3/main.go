package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// запуск go run main.go 3 -количество горутин
// worker – читает данные из канала и выводит
func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Воркер %d принял: %d\n", id, job)
	}
}

func main() {
	// количество воркеров
	if len(os.Args) < 2 {
		fmt.Println("напишите количество воркеров, например: go run main.go 3")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Неправильное число воркеров")
		return
	}

	jobs := make(chan int)

	// запуск воркеров
	for i := 1; i <= n; i++ {
		go worker(i, jobs)
	}

	// главная горутина бесконечно пишет в канал
	counter := 1
	for {
		jobs <- counter
		counter++
		time.Sleep(500 * time.Millisecond) // для наглядности
	}
}
