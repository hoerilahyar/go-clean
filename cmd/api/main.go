package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/bootstrap"
	"github.com/hoerilahyar/go-clean/internal/handler/middleware"
	"github.com/hoerilahyar/go-clean/internal/routes"
)

func main() {
	app := bootstrap.NewApplication()
	defer app.DB.Close()

	r := gin.New()

	r.Use(
		gin.Recovery(),
		middleware.Logger(),
	)

	routes.Register(r, app)

	log.Printf("Server running on :%s", app.Config.AppPort)

	log.Fatal(r.Run(":" + app.Config.AppPort))
}
