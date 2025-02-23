package main

import (
	"context"
	"core-api/internal/app"
	"core-api/internal/app/handlers"
	"core-api/internal/domain/services"
	"core-api/internal/infrastructure/database"
	"core-api/internal/infrastructure/influxdb"
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

var (
	influxClient   influxdb.ClientInfluxDBClient
	collectService *services.CollectService
)

func main() {
	if os.Getenv("DEBUG_MODE") == "true" {
		log.Println("🐛 Debug mode is enabled")
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	log.Println("🚀 Starting server...")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()

	// Initialize repositories
	userRepo := repository.NewUserRepository(database.DB)
	accountRepo := repository.NewAccountRepository(database.DB)
	buildingRepo := repository.NewBuildingRepository(database.DB)
	projectRepo := repository.NewProjectRepository(database.DB)
	effMeasureRepo := repository.NewEfficiencyMeasureRepository(database.DB)
	contractorRepo := repository.NewContractorRepository(database.DB)
	bmsRepo := repository.NewBmsRepository(database.DB)
	weatherStationRepo := repository.NewWeatherStationRepository(database.DB)
	meterRepo := repository.NewMeterRepository(database.DB)
	targetRepo := repository.NewTargetRepository(database.DB)
	subsidyRepo := repository.NewSubsidyRepository(database.DB)
	billRepo := repository.NewBillRepository(database.DB)
	equipmentRepo := repository.NewEquipmentRepository(database.DB)
	categoryRepo := repository.NewCategoryRepository(database.DB)
	regressionRepo := repository.NewRegressionRepository(database.DB)
	parameterRepo := repository.NewParameterRepository(database.DB)
	independantVariableRepo := repository.NewIndependantVariableRepository(database.DB)

	// Auth features
	authService := services.NewAuthService(*userRepo, *accountRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// Account features
	accountService := services.NewAccountService(*userRepo, *accountRepo)
	accountHandler := handlers.NewAccountHandler(accountService)

	// User features
	userService := services.NewUserService(*userRepo, *accountRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Building features
	buildingService := services.NewBuildingService(buildingRepo)
	buildingHandler := handlers.NewBuildingHandler(buildingService)

	// Project features
	projectService := services.NewProjectService(projectRepo)
	projectHandler := handlers.NewProjectHandler(projectService)

	// Efficiency Measure features
	effMeasureService := services.NewEfficiencyMeasureService(effMeasureRepo)
	efficiencyMeasureHandler := handlers.NewEfficiencyMeasureHandler(effMeasureService)

	// Contractor features
	contractorService := services.NewContractorService(contractorRepo)
	contractorHandler := handlers.NewContractorHandler(contractorService)

	// Bms features
	bmsService := services.NewBmsService(bmsRepo)
	bmsHandler := handlers.NewBmsHandler(bmsService)

	// WeatherStation features
	weatherStationService := services.NewWeatherStationService(weatherStationRepo)
	weatherStationHandler := handlers.NewWeatherStationHandler(weatherStationService)

	independantVariableService := services.NewIndependantVariableService(independantVariableRepo)
	independantVariableHandler := handlers.NewIndependantVariableHandler(independantVariableService)

	// Weather CSV Upload features
	influxClient = influxdb.NewClient()
	if influxClient == nil {
		log.Fatal("Failed to initialize InfluxDB client")
	}
	weatherUploadHandler := handlers.NewWeatherUploadHandler(influxClient)

	// Regression features
	regressionService := services.NewRegressionService(regressionRepo)
	regressionHandler := handlers.NewRegressionHandler(regressionService)

	// Meter features
	meterService := services.NewMeterService(meterRepo)
	meterHandler := handlers.NewMeterHandler(meterService)

	// Target features
	targetService := services.NewTargetService(targetRepo)
	targetHandler := handlers.NewTargetHandler(targetService)

	// Subsidy features
	subsidyService := services.NewSubsidyService(subsidyRepo)
	subsidyHandler := handlers.NewSubsidyHandler(subsidyService)

	// Bill features
	billService := services.NewBillService(billRepo)
	billHandler := handlers.NewBillHandler(billService)

	// Equipment features
	equipmentService := services.NewEquipmentService(equipmentRepo)
	equipmentHandler := handlers.NewEquipmentHandler(equipmentService)

	// Category features
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Collect features
	trendlogsService := services.NewTrendlogsService(influxClient)
	trendlogsHandler := handlers.NewTrendlogsHandler(trendlogsService)
	collectService = services.NewCollectService(influxClient)
	collectHandler := handlers.NewCollectHandler(collectService)
	savingsService := services.NewSavingsService(influxClient)
	savingsHandler := handlers.NewSavingsHandler(savingsService)

	// Instantiate Parameter features
	parameterService := services.NewParameterService(parameterRepo)
	parameterHandler := handlers.NewParameterHandler(parameterService)

	// Setup router with all handlers.
	router := app.SetupRouter(
		authHandler,
		accountHandler,
		userHandler,
		buildingHandler,
		projectHandler,
		efficiencyMeasureHandler,
		contractorHandler,
		bmsHandler,
		weatherStationHandler,
		weatherUploadHandler,
		regressionHandler,
		meterHandler,
		targetHandler,
		subsidyHandler,
		billHandler,
		equipmentHandler,
		categoryHandler,
		parameterHandler,
		independantVariableHandler,
		userRepo,
		collectHandler,
		trendlogsHandler,
		savingsHandler,
	)

	srv := &http.Server{
		Addr:    ":8080",
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
