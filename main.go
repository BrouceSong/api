package main

import (
	"github.com/gin-gonic/gin"
	routes "api/routes"
)

func main() {
	r := gin.Default()
	route := routes.Router(r)
	route.Run(":8080")
}

func hello() string {
	return "hello"
}
