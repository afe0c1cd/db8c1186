package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	posrgresRepository "github.com/afe0c1cd/db8c1186/database/postgres"
	"github.com/afe0c1cd/db8c1186/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := command(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func command() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}
	db, err := gorm.Open(postgres.Open(dbURL))
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	repo := posrgresRepository.NewPostgresRepository(db)
	srv := server.NewServer(repo)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	errCh := make(chan error, 1)
	go func() {
		if err := srv.Start(port); err != nil {
			errCh <- fmt.Errorf("failed to start server: %w", err)
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-sigCh:
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to shutdown server: %w", err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			return fmt.Errorf("failed to get database connection: %w", err)
		}
		if err := sqlDB.Close(); err != nil {
			return fmt.Errorf("failed to close database connection: %w", err)
		}
		return nil
	}
}
