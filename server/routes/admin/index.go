package admin

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func helloWorld(c *gin.Context) {

	c.String(http.StatusOK, "hello world")

}

type create_json struct {
	Name string

	Message string

	Age int
}

func json_struct(c *gin.Context) {

	data := create_json{

		Name: "Tarik",

		Message: "Hello World",

		Age: 24,
	}

	c.JSON(http.StatusOK, data)

}

func Admin(e *gin.Engine) {

	//r := gin.Default()

	e.GET("/admin", helloWorld)

	//e.GET("/admin/json", helloJson)

	e.GET("/admin/json_struct", json_struct)
	// return r

}
