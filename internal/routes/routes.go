package routes

import (
	"device-management-api/internal/controllers"
	"device-management-api/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	deviceRepo := repository.NewDeviceRepository(db)
	deviceController := controllers.NewDeviceController(deviceRepo)

	router.POST("/devices", deviceController.AddDevice)
	router.GET("/devices/:id", deviceController.GetDevice)
	router.GET("/devices", deviceController.ListDevices)
	router.PUT("/devices/:id", deviceController.UpdateDevice)
	router.DELETE("/devices/:id", deviceController.DeleteDevice)

	return router
}
