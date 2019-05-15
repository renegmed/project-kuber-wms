package main


import (
	"net/http"
	"os"
	//"strconv"
	//"fmt"
	"github.com/gin-gonic/gin"
)


func main() {
	
	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context){
		c.String(http.StatusOK, "pong")
	}) 

	engine.LoadHTMLGlob("./templates/*.html") 

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Warehouse System",
		})
	})
 
		// the hello message endpoint with JSON response from map
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Warehouse Management System."})
	})


	engine.Run(port())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8484"
	}
	return ":" + port
}