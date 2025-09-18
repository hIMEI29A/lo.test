package main

import (
	"context"
	"lo.test/api"
	"lo.test/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Инициализация репозитория
	repos := repository.New()

	// Канал для логирования
	logChan := make(chan string)
	defer close(logChan)

	// Запуск логгера
	go logger(logChan)

	// Инициализация обработчиков
	taskHandler := &api.TaskHandler{
		Repos:   repos,
		LogChan: logChan,
	}

	// Настройка маршрутов
	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", taskHandler.GetTasks)
	mux.HandleFunc("GET /tasks/{id}", taskHandler.GetTaskByID)
	mux.HandleFunc("POST /tasks", taskHandler.CreateTask)

	// Настройка сервера
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Запуск сервера в горутине
	go func() {
		logChan <- "Server starting on :8080"
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logChan <- "Server error: " + err.Error()
		}
	}()

	// Обработка graceful shutdown
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logChan <- "Shutting down server..."

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logChan <- "Server forced to shutdown: " + err.Error()
	}

	logChan <- "Server exiting"
	time.Sleep(100 * time.Millisecond) // Время на отправку последних логов
}

func logger(logChan <-chan string) {
	for msg := range logChan {
		log.Println("[LOG]", msg)
	}
}
