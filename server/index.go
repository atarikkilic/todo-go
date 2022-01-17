package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {

	c.String(http.StatusOK, "hello world")

}

func main() {

	r := gin.Default()

	r.GET("/", helloWorld)

	r.Run(":3000")

}
