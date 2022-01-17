package admin

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func helloWorld(c *gin.Context) {

	c.String(http.StatusOK, "hello world")

}

func helloJson(c *gin.Context) {

	data := gin.H{

		"name": "Tarik",

		"message": "Hello World",

		"age": 24,
	}

	c.JSON(http.StatusOK, data)

}

func Admin(e *gin.Engine) {

	//r := gin.Default()

	e.GET("/admin", helloWorld)

	e.GET("/admin/json", helloJson)

	// return r

}
