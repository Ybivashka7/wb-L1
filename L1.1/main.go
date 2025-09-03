package main

import "fmt"

// Родительская структура
type Human struct {
	name string
	age  int
}

// Метод структуры Human
func (h *Human) Speak() {
	fmt.Println("Привет " + h.name)
}

// Метод структуры Human
func (h *Human) Walk() {
	fmt.Println(h.name + " ходит")
}

// Дочерняя структура с  Human
type Action struct {
	Human
	Role string
}

func main() {
	a := Action{
		Human: Human{
			name: "Геннадий",
			age:  21,
		},
		Role: "Разработчик",
	}
	// Методы Human доступны напрямую у Action
	a.Speak()
	a.Walk()
	// Доступ к полям Human тоже напрямую
	fmt.Printf("%s работает как %s\n", a.name, a.Role)
}
