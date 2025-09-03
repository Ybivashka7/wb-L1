package main

import (
	"fmt"
	"reflect"
)

// обычный  swichtType
func swichtType(v interface{}) {
	switch val := v.(type) {
	case int, int8, int16, int32, int64:
		fmt.Printf("Целое число (%T): %v\n", val, val)
	case uint, uint8, uint16, uint32, uint64, uintptr:
		fmt.Printf("Беззнаковое целое (%T): %v\n", val, val)
	case float32, float64:
		fmt.Printf("Число с плавающей точкой (%T): %v\n", val, val)
	case complex64, complex128:
		fmt.Printf("Комплексное число (%T): %v\n", val, val)
	case string:
		fmt.Printf("Строка: %q\n", val)
	case bool:
		fmt.Printf("Булево: %v\n", val)
	case chan int, chan string, chan interface{}:
		fmt.Printf("Канал (%T)\n", val)
	case []int, []string, []interface{}:
		fmt.Printf("Срез (%T): %v\n", val, val)
	case map[string]int, map[int]string:
		fmt.Printf("Карта (%T): %v\n", val, val)
	case func():
		fmt.Println("Функция без аргументов")
	default:
		fmt.Printf("Неизвестный  тип: %T\n", val)
	}
}

func main() {
	swichtType(42)
	swichtType(uint16(7))
	swichtType(3.14)
	swichtType(make(chan int))
	swichtType([]string{"a", "b"})
	swichtType(map[string]int{"x": 1})
	swichtType(func() {})
	swichtType(struct{ X int }{10})
	swichtType("hello")
	swichtType(true)
	TypeReflect(42)
	TypeReflect(3.14)
	TypeReflect("hello")
	TypeReflect(true)
	TypeReflect(make(chan int))
	TypeReflect([]string{"a", "b"})
	TypeReflect(map[string]int{"x": 1})
	TypeReflect(func(x int) string { return "ok" })
	TypeReflect(struct{ X int }{10})
}

// Через рефлект
func TypeReflect(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Printf("Тип: %s, значение: %v\n", t.String(), v)
}
