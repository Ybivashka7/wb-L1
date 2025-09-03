package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// запуск go run main.go 3 -количество горутин
// worker – читает данные из канала и выводит
func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d завершает работу\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Воркер %d: канал закрыт\n", id)
				return
			}
			fmt.Printf("Воркер %d принял: %d\n", id, job)
		}
	}
}

func main() {
	// количество воркеров передаём параметром
	if len(os.Args) < 2 {
		fmt.Println("Укажите количество воркеров, например: go run main.go 3")
		return
	}
	// проверка количество воркеров и ошибку
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Неправильное число воркеров")
		return
	}

	jobs := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// запуск воркеров
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	// обработка Ctrl+C, ждем сигнала
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigs // отправляем сигнал о завершении
		fmt.Println("\n  сигнал завершения, останавливаемся")
		cancel()    // уведомляем воркеров о завершении работы
		close(jobs) // закрываем канал
	}()

	// горутина пишет в канал
	counter := 1
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			fmt.Println("Программа остановлена")
			return
		default:
			jobs <- counter
			counter++
			time.Sleep(500 * time.Millisecond) // для наглядности
		}
	}
}
