package main

import (
	"server/routes/admin"
)

func main() {

	// r := gin.Default()

	r := admin.Admin()

	r.Run(":3000")

}
