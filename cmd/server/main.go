package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Mukam21/io-bound-task-api_Golang/pkg/handlers"
	"github.com/Mukam21/io-bound-task-api_Golang/pkg/router"
	"github.com/Mukam21/io-bound-task-api_Golang/pkg/service"
)

func main() {
	repo := handlers.NewInMemoryTaskRepo(10000)
	taskService := service.NewTaskService(repo, 5)
	router := router.NewRouter(taskService)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Println("Server starting on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Forced shutdown: %v", err)
	}
	log.Println("Server stopped")
}
