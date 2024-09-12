package main

import (
	"fmt"
	"workout-evo/internal/adapter/api"
	"workout-evo/internal/core/service"
	"workout-evo/internal/port/repository"
)

func main() {
	// Выводим простой текст для проверки работы
	fmt.Println("Запуск приложения Workout-Evo...")

	// Выводим ASCII-арт
	fig := figure.NewFigure("SPORTS", "", true)
	fig.Print()

	// Инициализируем зависимости
	workoutRepo := repository.NewWorkoutRepository()
	workoutService := service.NewWorkoutService()
	workoutAPI := api.NewWorkoutAPI()

	// Выводим сообщение об успешном запуске
	fmt.Println("Приложение успешно запущено!")
}
