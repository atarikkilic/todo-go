package admin

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func helloWorld(c *gin.Context) {

	c.String(http.StatusOK, "hello world")

}

func Admin() *gin.Engine {

	r := gin.Default()

	r.GET("/admin", helloWorld)

	return r

}
