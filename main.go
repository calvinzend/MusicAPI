package main

import (
	routers "week2/Routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routers.RegisterRoutes(router)

	router.Run()
}
