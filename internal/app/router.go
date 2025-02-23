package app

import (
	"core-api/internal/app/handlers"
	"core-api/internal/app/middleware"
	"core-api/internal/infrastructure/repository"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures all routes for the API.
// CHANGES:
//   - Added new parameters "projectHandler", "efficiencyMeasureHandler", "contractorHandler", "bmsHandler",
//     "weatherStationHandler", "weatherUploadHandler", "regressionHandler", "meterHandler", "targetHandler",
//     "subsidyHandler", "billHandler", "equipmentHandler" and "categoryHandler" for the corresponding endpoints.
//   - For Regression:
//     POST   "/regressions"         -> CreateRegression
//     GET    "/regressions"         -> ListRegressions
//     GET    "/regressions/:id"     -> GetRegressionByID
//     PUT    "/regressions/:id"     -> UpdateRegression
//     DELETE "/regressions/:id"     -> DeleteRegression
//   - For Equipment:
//     POST   "/equipments"          -> CreateEquipment
//     GET    "/equipments"          -> ListAllEquipments
//     GET    "/equipments/:id"      -> GetEquipmentByID
//     PUT    "/equipments/:id"      -> UpdateEquipment
//     DELETE "/equipments/:id"      -> DeleteEquipment
//   - For Category:
//     GET    "/categories"          -> ListCategories
func SetupRouter(
	authHandler *handlers.AuthHandler,
	accountHandler *handlers.AccountHandler,
	userHandler *handlers.UserHandler,
	buildingHandler *handlers.BuildingHandler, // Building endpoints
	projectHandler *handlers.ProjectHandler, // Project endpoints
	efficiencyMeasureHandler *handlers.EfficiencyMeasureHandler, // Efficiency Measure endpoints
	contractorHandler *handlers.ContractorHandler, // Contractor endpoints
	bmsHandler *handlers.BmsHandler, // Bms endpoints
	weatherStationHandler *handlers.WeatherStationHandler, // WeatherStation endpoints
	weatherUploadHandler *handlers.WeatherUploadHandler, // Weather CSV upload endpoint
	regressionHandler *handlers.RegressionHandler, // Regression endpoints
	meterHandler *handlers.MeterHandler, // Meter endpoints
	targetHandler *handlers.TargetHandler, // Target endpoints
	subsidyHandler *handlers.SubsidyHandler, // Subsidy endpoints
	billHandler *handlers.BillHandler, // Bill endpoints
	equipmentHandler *handlers.EquipmentHandler, // Equipment endpoints
	categoryHandler *handlers.CategoryHandler, // Category endpoints
	parameterHandler *handlers.ParameterHandler, // Parameter endpoints
	independantVariableHandler *handlers.IndependantVariableHandler, // IndependantVariable endpoints
	userRepo *repository.UserRepository,
	collectHandler *handlers.CollectHandler,
	trendlogsHandler *handlers.TrendlogsHandler,
	savingsHandler *handlers.SavingsHandler,
) *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ready"})
	})

	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/login", authHandler.Login)
		apiV1.POST("/register", authHandler.Register)

		authRoutes := apiV1.Group("/")
		authRoutes.Use(middleware.JWTAuthMiddleware(*userRepo))
		{
			// User and Account routes
			authRoutes.GET("/securedata", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Secured page"})
			})
			authRoutes.GET("/users/me", userHandler.UsersMe)
			authRoutes.POST("/accounts", accountHandler.CreateAccount)
			authRoutes.GET("/accounts", accountHandler.ListAllAccounts)
			authRoutes.GET("/accounts/:id", accountHandler.GetAccountByID)
			authRoutes.POST("/accounts/users", accountHandler.AssociateUserToAccount)

			// Building endpoints
			authRoutes.POST("/buildings", buildingHandler.CreateBuilding)
			authRoutes.GET("/buildings", buildingHandler.GetAllBuildings)

			// Project endpoints
			authRoutes.POST("/projects", projectHandler.CreateProject)
			authRoutes.GET("/projects", projectHandler.ListAllProjects)
			authRoutes.GET("/projects/:id", projectHandler.GetProjectByID)
			authRoutes.PUT("/projects/:id", projectHandler.UpdateProject)

			// Efficiency Measure endpoints
			authRoutes.POST("/efficiencymeasures", efficiencyMeasureHandler.CreateEfficiencyMeasure)
			authRoutes.GET("/efficiencymeasures", efficiencyMeasureHandler.ListEfficiencyMeasures)
			authRoutes.GET("/efficiencymeasures/:id", efficiencyMeasureHandler.GetEfficiencyMeasureByID)
			authRoutes.PUT("/efficiencymeasures/:id", efficiencyMeasureHandler.UpdateEfficiencyMeasure)

			// Contractor endpoints
			authRoutes.POST("/contractors", contractorHandler.CreateContractor)
			authRoutes.GET("/contractors", contractorHandler.ListAllContractors)
			authRoutes.GET("/contractors/:id", contractorHandler.GetContractorByID)
			authRoutes.PUT("/contractors/:id", contractorHandler.UpdateContractor)
			authRoutes.DELETE("/contractors/:id", contractorHandler.DeleteContractor)

			// Bms endpoints
			authRoutes.POST("/bms", bmsHandler.CreateBms)
			authRoutes.GET("/bms", bmsHandler.ListBms)
			authRoutes.GET("/bms/:id", bmsHandler.GetBmsByID)
			authRoutes.PUT("/bms/:id", bmsHandler.UpdateBms)
			authRoutes.DELETE("/bms/:id", bmsHandler.DeleteBms)

			// WeatherStation endpoints
			authRoutes.POST("/weatherstations", weatherStationHandler.CreateWeatherStation)
			authRoutes.GET("/weatherstations", weatherStationHandler.ListWeatherStations)
			authRoutes.GET("/weatherstations/:id", weatherStationHandler.GetWeatherStationByID)
			authRoutes.PUT("/weatherstations/:id", weatherStationHandler.UpdateWeatherStation)
			authRoutes.DELETE("/weatherstations/:id", weatherStationHandler.DeleteWeatherStation)

			// Weather CSV Upload endpoint
			authRoutes.POST("/weather/upload", weatherUploadHandler.UploadWeatherCSV)

			// Regression endpoints
			authRoutes.POST("/regressions", regressionHandler.CreateRegression)
			authRoutes.GET("/regressions", regressionHandler.ListRegressions)
			authRoutes.GET("/regressions/:id", regressionHandler.GetRegressionByID)
			authRoutes.PUT("/regressions/:id", regressionHandler.UpdateRegression)
			authRoutes.DELETE("/regressions/:id", regressionHandler.DeleteRegression)

			// Meter endpoints
			authRoutes.POST("/meters", meterHandler.CreateMeter)
			authRoutes.GET("/meters", meterHandler.ListAllMeters)
			authRoutes.GET("/meters/:id", meterHandler.GetMeterByID)
			authRoutes.PUT("/meters/:id", meterHandler.UpdateMeter)
			authRoutes.DELETE("/meters/:id", meterHandler.DeleteMeter)

			// Target endpoints
			authRoutes.POST("/targets", targetHandler.CreateTarget)
			authRoutes.GET("/targets", targetHandler.ListTargets)
			authRoutes.GET("/targets/:id", targetHandler.GetTargetByID)
			authRoutes.PUT("/targets/:id", targetHandler.UpdateTarget)
			authRoutes.DELETE("/targets/:id", targetHandler.DeleteTarget)

			// Subsidy endpoints
			authRoutes.POST("/subsidies", subsidyHandler.CreateSubsidy)
			authRoutes.GET("/subsidies", subsidyHandler.ListSubsidies)
			authRoutes.GET("/subsidies/:id", subsidyHandler.GetSubsidyByID)
			authRoutes.PUT("/subsidies/:id", subsidyHandler.UpdateSubsidy)
			authRoutes.DELETE("/subsidies/:id", subsidyHandler.DeleteSubsidy)

			// Bill endpoints
			authRoutes.POST("/bills", billHandler.CreateBill)
			authRoutes.GET("/bills", billHandler.ListBills)
			authRoutes.GET("/bills/:id", billHandler.GetBillByID)
			authRoutes.PUT("/bills/:id", billHandler.UpdateBill)
			authRoutes.DELETE("/bills/:id", billHandler.DeleteBill)

			// Equipment endpoints
			authRoutes.POST("/equipments", equipmentHandler.CreateEquipment)
			authRoutes.GET("/equipments", equipmentHandler.ListAllEquipments)
			authRoutes.GET("/equipments/:id", equipmentHandler.GetEquipmentByID)
			authRoutes.PUT("/equipments/:id", equipmentHandler.UpdateEquipment)
			authRoutes.DELETE("/equipments/:id", equipmentHandler.DeleteEquipment)

			// Category endpoints
			authRoutes.GET("/categories", categoryHandler.ListCategories)

			// Parameter endpoints
			authRoutes.POST("/parameters", parameterHandler.CreateParameter)
			authRoutes.GET("/parameters", parameterHandler.ListParameters)
			authRoutes.GET("/parameters/:id", parameterHandler.GetParameterByID)
			authRoutes.PUT("/parameters/:id", parameterHandler.UpdateParameter)
			authRoutes.DELETE("/parameters/:id", parameterHandler.DeleteParameter)

			// IndependantVariable endpoints
			authRoutes.POST("/independantvariables", independantVariableHandler.CreateIndependantVariable)
			authRoutes.GET("/independantvariables", independantVariableHandler.ListIndependantVariables)
			authRoutes.GET("/independantvariables/:id", independantVariableHandler.GetIndependantVariableByID)
			authRoutes.PUT("/independantvariables/:id", independantVariableHandler.UpdateIndependantVariable)
			authRoutes.DELETE("/independantvariables/:id", independantVariableHandler.DeleteIndependantVariable)

			// Other routes
			authRoutes.POST("/collect", collectHandler.CollectHandler)
			authRoutes.POST("/trendlogs", trendlogsHandler.GetTrendlogs)
			authRoutes.POST("/savings", savingsHandler.GetSavings)
		}
	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
