package main

import (
	"context"
	"core-api/internal/app"
	"core-api/internal/domain/services"
	"core-api/internal/infrastructure/database"
	"core-api/internal/infrastructure/repository"
	"core-api/internal/interfaces/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if os.Getenv("DEBUG_MODE") == "true" {
		log.Println("🐛 Debug mode is enabled")
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	log.Println("🚀 Starting server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()

	userRepo := repository.NewUserRepository(database.DB)

	authService := services.NewAuthService(*userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	accountService := services.NewAccountService(*userRepo)
	accountHandler := handlers.NewAccountHandler(accountService)

	router := app.SetupRouter(authHandler, accountHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Listen: %s\n", err)
		}
	}()
	log.Println("🔊 Server is listening on :8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Received shutdown signal")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server forced to shutdown: %v", err)
	} else {
		log.Println("✅ Server shutdown gracefully.")
	}

	log.Println("👋 Server exiting")
}
