package main

import (
	
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	initDB()
	defer db.Close()

	result, _ := startCompareImage("./public/bob degueulasse.jpg")
	log.Println(result)
	
	r := gin.Default()
	r.POST("/add", routeAddImage)
	r.Run() // listen and serve on 0.0.0.0:8080
}
