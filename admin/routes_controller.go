package admin

import (
	"context"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"

	"kraken/persistence"
)

func (admin Admin) ListRoutes(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error;
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		routes, err := admin.Queries.ListRoutes(ctx)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{
			"routes": routes,
		})
	
	}
}

func (admin Admin) GetRoute(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error;
	
		routeID, err := strconv.ParseInt(c.Param("route"), 10, 64)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		route, err := admin.Queries.GetRoute(ctx, routeID)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{
			"routes": route,
		})
	}
}

func (admin Admin) DeleteRoute(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error;

		routeID, err := strconv.ParseInt(c.Param("route"), 10, 64)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		err = admin.Queries.DeleteRoute(ctx, routeID)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{
			"message": "Route deleted",
		})

		admin.ReloadEvent.Emit()
	}
}

func (admin Admin) UpdateRoute(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error;

		// routeID, err := strconv.ParseInt(c.Param("route"), 10, 64)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		var route persistence.UpdateRouteParams
		err = c.BindJSON(&route)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		updatedRoute, err := admin.Queries.UpdateRoute(ctx, route)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{
			"routes": updatedRoute,
		})

		admin.ReloadEvent.Emit()
	}	
}

func (admin Admin) CreateRoute(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error;

		var route persistence.CreateRouteParams
		err = c.BindJSON(&route)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		createdRoute, err := admin.Queries.CreateRoute(ctx, route)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{
			"routes": createdRoute,
		})

		admin.ReloadEvent.Emit()
	}
}