package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	stopByFlag()
	stopByChannel()
	stopByContext()
	stopByGoexit()
}

// 1. Остановка по условию
func stopByFlag() {
	stop := false

	go func() {
		for {
			if stop {
				fmt.Println("stop by flag")
				return
			}
			fmt.Println("start by flag")
			time.Sleep(300 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
	stop = true
	time.Sleep(500 * time.Millisecond)
}

// 2. Остановка через канал
func stopByChannel() {
	stop := make(chan struct{})

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("stop by channel")
				return
			default:
				fmt.Println("start by channel")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	close(stop)
	time.Sleep(500 * time.Millisecond)
}

// 3. Остановка через context
func stopByContext() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stop by context")
				return
			default:
				fmt.Println("start by context")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
}

// 4. Принудительная остановка
func stopByGoexit() {
	go func() {
		fmt.Println("stop by goexit")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Goexit")
		runtime.Goexit()
	}()

	time.Sleep(1 * time.Second)
}
