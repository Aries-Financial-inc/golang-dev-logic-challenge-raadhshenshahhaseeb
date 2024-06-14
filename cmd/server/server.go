package server

import (
	"fmt"
	"io"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/controllers"
)

type Container struct {
	r           *gin.Engine
	controllers *controllers.Controllers
}

func Start() error {
	container := new(Container)

	err := container.initControllers()
	if err != nil {
		return fmt.Errorf("unable to initialize controllers: %w", err)
	}

	container.r = SetupRouter()
	container.allowedPaths()

	return container.r.Run("localhost:8080")
}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()

	allowedCORS(router)

	return router
}

func allowedCORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func (c *Container) allowedPaths() {
	public := c.r.Group("/api/v1")
	{
		public.POST("/analyze", c.controllers.AnalysisHandler)
	}
}

func (c *Container) initControllers() error {
	var err error

	c.controllers, err = controllers.New()
	if err != nil {
		return fmt.Errorf("unable to init new controllers: %w", err)
	}

	return nil
}
