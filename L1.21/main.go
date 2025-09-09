package main

import "fmt"

// OldLogger — старый логгер с интерфейсом
type OldLogger struct{}

// Старый метод логгера
func (l *OldLogger) WriteLog(msg string) {
	fmt.Println("OLD LOGGER:", msg)
}

// Новый интерфейс, который ожидает клиентский код
type Logger interface {
	Log(message string)
}

// Клиентский код — работает только с Logger
func Process(logger Logger) {
	logger.Log("Processing started")
	// какаето работа
	logger.Log("Processing finished")
}

// LoggerAdapter — адаптер, приводящий OldLogger к интерфейсу Logger
type LoggerAdapter struct {
	oldLogger *OldLogger
}

// Реализация метода Log
func (a *LoggerAdapter) Log(message string) {
	a.oldLogger.WriteLog(message)
}

func main() {
	// Создаём экземпляр старого логгера
	old := &OldLogger{}
	// Оборачиваем его адаптером
	adapter := &LoggerAdapter{oldLogger: old}
	// Передаём адаптер в клиентский код
	Process(adapter)
}
