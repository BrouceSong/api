package main

import (
	routes "api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	route := routes.Router(r)
	route.Run(":8080")
}

func hello() string {
	return "hello"
}
