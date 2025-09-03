package main

import (
	"fmt"
	"time"
)

func main() {
	timer := 10
	ch := make(chan int)
	//горутина пишет
	go func() {
		defer close(ch)
		count := 1
		for {
			ch <- count
			count++
			time.Sleep(300 * time.Millisecond) //Иммитаця работы
		}
	}()
	//горутина читает
	go func() {
		for val := range ch {
			fmt.Printf("Получено: %d\n", val)
		}
	}()
	<-time.After(time.Duration(timer) * time.Second)
	fmt.Println("Время вышло")

}
