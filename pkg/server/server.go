package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	_ "modernc.org/sqlite"

	"proj/pkg/api/handlers"
	"proj/pkg/api/middleware"
	"proj/pkg/api/service"
)

func ServerRun() *http.Server {
	r := chi.NewRouter()

	r.Post("/api/signin", handlers.LoginHandle)
	r.Get("/api/nextdate", handlers.NextDayHandler)

	pass := service.AutherENV()

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(pass))
		r.Post("/api/task", handlers.AddTaskHandler)
		r.Get("/api/tasks", handlers.GetTaskHandler)
		r.Get("/api/task", handlers.GetIDHandler)
		r.Put("/api/task", handlers.PutTaskHandler)
		r.Post("/api/task/done", handlers.DonePostHandler)
		r.Delete("/api/task", handlers.DeleteTaskHandler)
	})

	fs := http.FileServer(http.Dir("./web"))
	r.Handle("/*", fs)

	server := &http.Server{
		Addr:         ":" + service.PortGlobal(),
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("Запускаем сервер...")
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Ошибка сервера: %v", err)
			return
		}
	}()

	return server

}
