package main

import (
	//"github.com/gin-gonic/gin"

	"server/routes"
	"server/routes/admin"
	"server/routes/web"
)

/*
type Option func(*gin.Engine)

var options = []Option{}

func include(opts ...Option) {

	options = append(options, opts...)

}

func arrange() *gin.Engine {

	r := gin.New()

	for _, opt := range options {

		opt(r)

	}

	return r

}
*/
func main() {

	// r := gin.Default()

	// r := admin.Admin()

	routes.Include(admin.Admin, web.Web)

	r := routes.Arrange()

	r.Run(":3000")

}
