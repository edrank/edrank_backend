package main

import (
	"github.com/edrank/edrank_backend/config"
	"github.com/edrank/edrank_backend/db"
	"github.com/edrank/edrank_backend/middlewares"
	"github.com/edrank/edrank_backend/routes"
	"github.com/edrank/edrank_backend/services"
	"github.com/edrank/edrank_backend/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// load config
	dbConfig := config.LoadConfig()

	// load aws env
	utils.LoadEnv()

	// connect to aws (create session)
	awsSession, err := services.ConnectAws()

	if err != nil {
		utils.PrintToConsole("Cannot connect to aws : "+err.Error(), "red")
		return
	}
	// init db
	db.Init(dbConfig)

	// initialize router
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// set aws session to router context
	router.Use(func(c *gin.Context) {
		c.Set("sess", awsSession)
		c.Next()
	})

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
