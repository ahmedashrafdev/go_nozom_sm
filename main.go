package main

import (
	"hand_held/db"
	_ "hand_held/docs" // docs is generated by Swag CLI, you have to import it.
	"hand_held/handler"
	"hand_held/router"

	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

// @title Swagger Example API
// @version 1.0
// @description Conduit API
// @title Conduit API

// @host 127.0.0.1:8585
// @BasePath /api

// @schemes http https
// @produce	application/json
// @consumes application/json

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	r := router.New()

	r.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := r.Group("/api")
	db.InitDatabase()
	db := db.DBConn

	h := handler.NewHandler(db)
	h.Register(v1)
	r.Logger.Fatal(r.Start(":8585"))
}
