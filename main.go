package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	router := SetupRouter()
	_ = router.Run()
}
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	return router
}