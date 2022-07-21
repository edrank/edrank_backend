package main

import (
	"github.com/edrank/edrank_backend/config"
	"github.com/edrank/edrank_backend/db"
	"github.com/edrank/edrank_backend/routes"
	"github.com/edrank/edrank_backend/middlewares"
	"github.com/edrank/edrank_backend/utils"
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
	publicRoutes := router.Group("/api/" + utils.GetVersion() + "/")
	privateRoutes := router.Group("/api/" + utils.GetVersion() + "/")

	privateRoutes.Use(middlewares.JWTMiddleware())

	routes.InitRoutes(router)
	routes.InitPublicRoutes(publicRoutes)
	routes.InitPrivateRoutes(privateRoutes)

	// change to ip:port to make it available on the local network
	router.Run(":5000")
}