package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// init 操作
	r := gin.Default()

	initRouter(r)

	r.Run("127.0.0.1:8080")
}
