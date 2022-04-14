package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func getDefault(c *gin.Context) {
	processCall(c, 0)
}

func get(c *gin.Context) {

	numberOfSleeps, _ := strconv.Atoi(c.Param("numberOfSleeps"))

	processCall(c, numberOfSleeps)
}

func processCall(c *gin.Context, numberOfSleeps int) {
	for i := 0; i < numberOfSleeps; i++ {
		time.Sleep(100 * time.Millisecond)
	}

	c.JSON(http.StatusOK, c.Request.URL.Path)
}

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	router := gin.New()

	gin.SetMode(gin.ReleaseMode)
	router.GET("/", get)
	router.GET("/:numberOfSleeps", get)

	log.Fatal(router.Run(":" + httpPort))
}
