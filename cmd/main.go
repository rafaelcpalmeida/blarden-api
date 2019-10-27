package main

import (
	"blarden-api/src/api"
	"blarden-api/src/api/door"
	"blarden-api/src/api/users"
	"blarden-api/src/db"
	"blarden-api/src/db/models"
	"blarden-api/src/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	runMigrations()
	router := SetupRouter()
	_ = router.Run()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", api.DefaultHandler)

	router.GET("/open-door", middleware.CheckUserPermissions(), door.Open)
	router.GET("/garage-gate", middleware.CheckUserPermissions(), door.GarageGate)
	router.GET("/outside-gate", middleware.CheckUserPermissions(), door.OutsideGate)

	usersRoute := router.Group("/users")
	{
		usersRoute.GET("/", users.All)
		usersRoute.GET("/:id", users.WithId)
		usersRoute.POST("/", users.Create)
		usersRoute.PUT("/:id", users.Update)
		usersRoute.DELETE("/:id", users.Delete)
	}

	return router
}

func runMigrations() {
	DbConnection := db.DatabaseHandler()
	DbConnection.AutoMigrate(&models.User{})
	DbConnection.Model(&models.User{}).AddUniqueIndex("idx_user_phone_number", "phone_number")

	fmt.Println("Migrated")

	defer DbConnection.Close()
}