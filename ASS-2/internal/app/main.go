package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"ASS-2/internal/handler"
	"ASS-2/internal/repository"
	_postgres "ASS-2/internal/repository/postgres"
	"ASS-2/internal/usecase"
	"ASS-2/pkg/modules"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbConfig := initPostgreConfig()
	db := _postgres.NewPGXDialect(ctx, dbConfig)

	repos := repository.NewRepositories(db)
	usecases := usecase.NewUsecases(repos)
	h := handler.NewHandler(usecases)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      h.InitRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server starting on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func initPostgreConfig() *modules.PostgreConfig {
	return &modules.PostgreConfig{
		Host:        "localhost",
		Port:        "5432",
		Username:    "dosyana",
		Password:    "maminpapin",
		DBName:      "dosyana",
		SSLMode:     "disable",
		ExecTimeout: 5 * time.Second,
	}
}
