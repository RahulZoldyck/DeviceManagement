package controllers

import (
	"device-management-api/internal/models"
	"device-management-api/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeviceController struct {
	Repo *repository.DeviceRepository
}

func NewDeviceController(repo *repository.DeviceRepository) *DeviceController {
	return &DeviceController{Repo: repo}
}

// AddDevice: Adds a new device to the database
func (ctrl *DeviceController) AddDevice(c *gin.Context) {
	var device models.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newDevice, err := ctrl.Repo.AddDevice(&device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create device"})
		return
	}

	c.JSON(http.StatusCreated, newDevice)
}

// GetDevice: Retrieves a device by its ID
func (ctrl *DeviceController) GetDevice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device ID"})
		return
	}

	device, err := ctrl.Repo.GetDevice(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}

	c.JSON(http.StatusOK, device)
}

// ListDevices: Lists all devices in the database
func (ctrl *DeviceController) ListDevices(c *gin.Context) {
	devices, err := ctrl.Repo.ListDevices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve devices"})
		return
	}

	c.JSON(http.StatusOK, devices)
}

// UpdateDevice: Updates an existing device by its ID
func (ctrl *DeviceController) UpdateDevice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device ID"})
		return
	}

	var updatedDevice models.Device
	if err := c.ShouldBindJSON(&updatedDevice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	device, err := ctrl.Repo.UpdateDevice(uint(id), &updatedDevice)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}

	c.JSON(http.StatusOK, device)
}

// DeleteDevice: Deletes a device by its ID
func (ctrl *DeviceController) DeleteDevice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device ID"})
		return
	}

	err = ctrl.Repo.DeleteDevice(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// SearchDevicesByBrand: Searches for devices by their brand
func (ctrl *DeviceController) SearchDevicesByBrand(c *gin.Context) {
	brand := c.Query("brand")
	if brand == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Brand parameter is required"})
		return
	}

	devices, err := ctrl.Repo.SearchDevicesByBrand(brand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve devices"})
		return
	}

	c.JSON(http.StatusOK, devices)
}
