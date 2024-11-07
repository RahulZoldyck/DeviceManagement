package routes

import (
	"device-management-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/devices", controllers.AddDevice)
	router.GET("/devices/:id", controllers.GetDevice)
	router.GET("/devices", controllers.ListDevices)
	router.PUT("/devices/:id", controllers.UpdateDevice)
	router.DELETE("/devices/:id", controllers.DeleteDevice)
	router.GET("/devices/search", controllers.SearchDevicesByBrand)

	return router
}
