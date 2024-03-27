package main

import (
	"context"
	"core-api/internal/app"
	"core-api/internal/app/handlers"
	"core-api/internal/domain/services"
	"core-api/internal/infrastructure/database"
	"core-api/internal/infrastructure/kafka"
	"core-api/internal/infrastructure/repository"
	"errors"
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

	// Init repo
	userRepo := repository.NewUserRepository(database.DB)
	accountRepo := repository.NewAccountRepository(database.DB)
	buildingRepo := repository.NewBuildingRepository(database.DB)
	systemRepo := repository.NewSystemRepository(database.DB)
	equipmentRepo := repository.NewEquipmentRepository(database.DB)
	areaRepo := repository.NewAreaRepository(database.DB)

	// Auth features
	authService := services.NewAuthService(*userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// Account features
	accountService := services.NewAccountService(*userRepo, *accountRepo)
	accountHandler := handlers.NewAccountHandler(accountService)

	// User features
	userService := services.NewUserService(*userRepo, *accountRepo)
	userHandler := handlers.NewUserHandler(userService)

	buildingService := services.NewBuildingService(*buildingRepo, *systemRepo, *equipmentRepo, areaRepo)
	buildingHandler := handlers.NewBuildingHandler(accountService, &buildingService)

	kafkaProducer, _ := kafka.NewKafkaProducer("localhost")
	collectService := services.NewCollectService(*kafkaProducer)
	collectHandler := handlers.NewCollectHandler(collectService)

	router := app.SetupRouter(authHandler, accountHandler, userHandler, buildingHandler, collectHandler, userRepo)

	srv := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
