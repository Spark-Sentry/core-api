package app

import (
	"core-api/internal/app/handlers"
	"core-api/internal/app/middleware"
	"core-api/internal/infrastructure/repository"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	authHandler *handlers.AuthHandler,
	accountHandler *handlers.AccountHandler,
	userHandler *handlers.UserHandler,
	buildingHandler *handlers.BuildingHandler,
	userRepo *repository.UserRepository,
	collectHandler *handlers.CollectHandler,
	trendlogsHandler *handlers.TrendlogsHandler,
	savingsHandler *handlers.SavingsHandler,
) *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ready",
		})
	})
	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/login", authHandler.Login)
		apiV1.POST("/register", authHandler.Register)

		authenticatedRoutes := apiV1.Group("/")
		authenticatedRoutes.Use(middleware.JWTAuthMiddleware(*userRepo))
		{
			authenticatedRoutes.GET("/securedata", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Secured page",
				})
			})
			authenticatedRoutes.GET("/users/me", userHandler.UsersMe)
			authenticatedRoutes.POST("/accounts", accountHandler.CreateAccount)
			authenticatedRoutes.GET("/accounts", accountHandler.ListAllAccounts)
			authenticatedRoutes.GET("/accounts/:id", accountHandler.GetAccountByID)
			authenticatedRoutes.POST("/accounts/users", accountHandler.AssociateUserToAccount)

			// Building Routes
			// Create a new building
			authenticatedRoutes.POST("/buildings", buildingHandler.HandleCreateBuilding) // "Create a new building."
			// List all buildings for the authenticated account
			authenticatedRoutes.GET("/buildings", buildingHandler.GetAllBuildings) // "Retrieve all buildings associated with the authenticated account."

			// Area Management Routes
			// Create an area within a specific building
			authenticatedRoutes.POST("/buildings/:building_id/areas", buildingHandler.AddArea) // "Add a new area to a specific building."
			// Retrieve all areas of a specific building
			authenticatedRoutes.GET("/buildings/:building_id/areas", buildingHandler.GetAreasByBuildingID) // "Retrieve all areas associated with a specific building."
			// Update a specific area (if needed)
			authenticatedRoutes.PUT("/areas/:area_id", buildingHandler.UpdateArea) // "Update details of a specific area."
			// Delete a specific area (if needed)
			authenticatedRoutes.DELETE("/areas/:area_id", buildingHandler.DeleteArea) // "Delete a specific area."

			// System Management Routes
			// Add a new system to a specific area within a building
			authenticatedRoutes.POST("/buildings/:building_id/areas/:area_id/systems", buildingHandler.AddSystem) // "Create a new system within a specific area of a building."
			// Retrieve all systems associated with a specific area within a building
			authenticatedRoutes.GET("/buildings/:building_id/areas/:area_id/systems", buildingHandler.GetSystemsByAreaID) // "List all systems within a specific area of a building."
			// Update a specific system (if needed)
			authenticatedRoutes.PUT("/systems/:system_id", buildingHandler.UpdateSystem) // "Update details of a specific system."
			// Delete a specific system
			authenticatedRoutes.DELETE("/systems/:system_id", buildingHandler.DeleteSystem) // "Remove a specific system."

			// Equipment Management Routes
			// Add new equipment to a specific system
			authenticatedRoutes.POST("/systems/:system_id/equipments", buildingHandler.AddEquipmentToSystem) // "Add new equipment to a specific system."
			// Retrieve all equipment associated with a specific system
			authenticatedRoutes.GET("/systems/:system_id/equipments", buildingHandler.GetEquipmentsBySystemID) // "List all equipments within a specific system."
			// Update a specific piece of equipment (if needed)
			authenticatedRoutes.PUT("/equipments/:equipment_id", buildingHandler.UpdateEquipment) // "Update details of a specific piece of equipment."
			// Delete a specific piece of equipment
			authenticatedRoutes.DELETE("/equipments/:equipment_id", buildingHandler.DeleteEquipment) // "Remove a specific piece of equipment."

			// Parameter Management Routes
			// Add a new parameter to a specific equipment
			authenticatedRoutes.POST("/equipments/:equipment_id/parameters", buildingHandler.AddParameterToEquipment) // "Add a new parameter to a specific equipment."
			// Retrieve all parameters associated with a specific equipment
			authenticatedRoutes.GET("/equipments/:equipment_id/parameters", buildingHandler.GetParametersByEquipmentID) // "List all parameters associated with a specific equipment."
			// Collect Data Routes
			// Collect data handler main
			authenticatedRoutes.POST("/collect", collectHandler.CollectHandler) // "Collect data for a specific parameter."

			// Trendlogs route
			// Retrieve timeseries data
			authenticatedRoutes.POST("/trendlogs", trendlogsHandler.GetTrendlogs)

			// Savings route
			// Retrieve savings data
			authenticatedRoutes.POST("/savings", savingsHandler.GetSavings)

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
