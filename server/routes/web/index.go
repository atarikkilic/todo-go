package web

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func helloWorld(c *gin.Context) {

	c.String(http.StatusOK, "hello world")

}

func Web(e *gin.Engine) {

	// r := gin.Default()

	e.GET("/web", helloWorld)

	// return r

}
