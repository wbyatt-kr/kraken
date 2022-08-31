package admin

import (
	"context"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func (admin Admin) ListServices(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error;
	
		services, err := admin.Queries.ListServices(ctx)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{
			"services": services,
		})
	}
}

func (admin Admin) GetService(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error;
	
		serviceID, err := strconv.ParseInt(c.Param("service"), 10, 64)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		service, err := admin.Queries.GetService(ctx, serviceID)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{
			"services": service,
		})
	}
}

type CreateServiceParams struct {
	Backend string `json:"backend" binding:"required"`
}

func (admin Admin) CreateService(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error;
	
		var serviceParams CreateServiceParams
		c.BindJSON(&serviceParams)
		
		service, err := admin.Queries.CreateService(ctx, serviceParams.Backend)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{
			"service": service,
		})
		
		admin.ReloadEvent.Emit()
	}
}

func (admin Admin) DeleteService(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error;
	
		serviceID, err := strconv.ParseInt(c.Param("service"), 10, 64)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		err = admin.Queries.DeleteService(ctx, serviceID)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{
			"message": "Service deleted",
		})

		admin.ReloadEvent.Emit()
	}
}