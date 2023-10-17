package main

import (
	"github.com/aifuxi/gin-ranking/models"
	"github.com/aifuxi/gin-ranking/router"
	"log"
)

func main() {
	models.Init()

	r := router.New()
	log.Fatal(r.Run("127.0.0.1:8080")) // listen and serve on 0.0.0.0:8080
}
