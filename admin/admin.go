package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"kraken/events"
	"kraken/persistence"
)

type Admin struct {
	Queries *persistence.Queries
	ReloadEvent events.Event
}

func (admin Admin) New(ctx context.Context, bind string) error {
	router := gin.Default()
	
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:			[]string{"*"},
	}))

	router.GET("/services", admin.ListServices(ctx))
	router.GET("/services/:service", admin.GetService(ctx))
	router.POST("/services", admin.CreateService(ctx))
	// router.PUT("/services/:service", UpdateService)
	router.DELETE("/services/:service", admin.DeleteService(ctx))

	router.GET("/routes", admin.ListRoutes(ctx))
	router.GET("/routes/:route", admin.GetRoute(ctx))
	router.POST("/routes", admin.CreateRoute(ctx))
	router.PUT("/routes/:route", admin.UpdateRoute(ctx))
	router.DELETE("/routes/:route", admin.DeleteRoute(ctx))

	return router.Run(bind)
}