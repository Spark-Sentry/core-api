package app

import (
	"core-api/internal/app/handlers"
	"core-api/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *handlers.AuthHandler, accountHandler *handlers.AccountHandler, userHandler *handlers.UserHandler, buildingHandler *handlers.BuildingHandler) *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/login", authHandler.Login)

		authenticatedRoutes := apiV1.Group("/")
		authenticatedRoutes.Use(middleware.JWTAuthMiddleware())
		{
			authenticatedRoutes.GET("/securedata", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Secured page",
				})
			})
			authenticatedRoutes.POST("/register", authHandler.Register)
			authenticatedRoutes.GET("/users/me", userHandler.UsersMe)
			authenticatedRoutes.POST("/accounts", accountHandler.CreateAccount)
			authenticatedRoutes.POST("/accounts/users", accountHandler.AssociateUserToAccount)

			// Buildings
			authenticatedRoutes.POST("/buildings", buildingHandler.HandleCreateBuilding)
			authenticatedRoutes.GET("/buildings", buildingHandler.GetAllBuildings)
			// Buildings - System
			authenticatedRoutes.POST("/buildings/:building_id/systems", buildingHandler.AddSystem)
			authenticatedRoutes.GET("/buildings/:building_id/systems", buildingHandler.GetSystemsByBuildingID)

		}
	}

	return router
}
