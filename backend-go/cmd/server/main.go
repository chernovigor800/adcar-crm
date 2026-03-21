package main

import (
	"github.com/gin-gonic/gin"
	"adcar-crm/backend-go/internal/rest"
)

func main() {
	r := gin.Default()
	rest.Setup(r)
	r.Run(":8080")
}