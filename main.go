package main

import (
	"github.com/edrank/edrank_backend/config"
	"github.com/edrank/edrank_backend/db"
	"github.com/edrank/edrank_backend/routes"
	"github.com/edrank/edrank_backend/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	// load config
	dbConfig := config.LoadConfig()
	
	// init db
	db.Init(dbConfig)

	// initialize router
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// attach cors middlware
	router.Use(middlewares.CORSMiddleware())

	// initialize routes
	routes.InitRoutes(router)

	// change to ip:port to make it available on the local network
	router.Run(":5000")
}