package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"proj/pkg/api/service"
	"proj/pkg/db"
	"proj/pkg/server"
)

func main() {

	if err := db.Init(service.PathDbManual()); err != nil {
		log.Fatal("База данных не открылась. Запуск сервера отменен.")
	}
	defer db.DBOpen.Close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	httpServer := server.ServerRun()
	<-sigCh

	log.Println("Получен сигнал остановки сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("Ошибка при остановке сервера: %v", err)
		return
	}

	log.Println("Сервер остановлен.")
}
