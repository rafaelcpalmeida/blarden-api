package main

import (
	"blarden-api/api"
	"blarden-api/api/users"
	"blarden-api/db"
	"blarden-api/db/models"
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