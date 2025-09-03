package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}
	gen := generator(nums)
	mul := multiplier(gen)
	for res := range mul {
		fmt.Println(res)
	}

}

// генератор чисел
func generator(nums []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, n := range nums {
			ch <- n
		}
	}()
	return ch
}

// функция умножения на 2
func multiplier(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out

}
