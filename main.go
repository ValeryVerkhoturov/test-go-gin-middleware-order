package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//2024/08/30 21:20:37 1 Before
//2024/08/30 21:20:37 2 Before
//2024/08/30 21:20:37 3 Before
//2024/08/30 21:20:37 4 Before
//2024/08/30 21:20:37 5 Before
//2024/08/30 21:20:37 Calling ping
//2024/08/30 21:20:37 5 After
//2024/08/30 21:20:37 4 After
//2024/08/30 21:20:37 3 After
//2024/08/30 21:20:37 2 After
//2024/08/30 21:20:37 1 After

func main() {
	r := gin.Default()
	r.Use(middleware(1))
	r.Use(middleware(2))
	r.Use(middleware(3))
	r.Use(middleware(4), middleware(5))
	r.GET("/ping", func(c *gin.Context) {
		log.Default().Printf("Calling ping")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8090")
}

func middleware(num int) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Default().Printf("%v Before\n", num)
		c.Next()
		log.Default().Printf("%v After\n", num)
	}
}
