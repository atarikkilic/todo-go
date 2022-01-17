package admin

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func helloWorld(c *gin.Context) {

	c.String(http.StatusOK, "hello world")

}

func Admin(e *gin.Engine) {

	//r := gin.Default()

	e.GET("/admin", helloWorld)

	// return r

}
