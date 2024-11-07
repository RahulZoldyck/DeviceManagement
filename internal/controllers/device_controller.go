package controllers

import (
	"device-management-api/internal/models"
	"device-management-api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddDevice(c *gin.Context) {
	var device models.Device
	if err := c.BindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newDevice, err := services.AddDevice(&device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create device"})
		return
	}
	c.JSON(http.StatusCreated, newDevice)
}

func GetDevice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device ID"})
		return
	}

	device, err := services.GetDevice(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}

	c.JSON(http.StatusOK, device)
}

func ListDevices(c *gin.Context) {
	devices, err := services.ListDevices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve devices"})
		return
	}

	c.JSON(http.StatusOK, devices)
}

func UpdateDevice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device ID"})
		return
	}

	var updatedDevice models.Device
	if err := c.BindJSON(&updatedDevice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	device, err := services.UpdateDevice(id, &updatedDevice)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}

	c.JSON(http.StatusOK, device)
}

func DeleteDevice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device ID"})
		return
	}

	err = services.DeleteDevice(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func SearchDevicesByBrand(c *gin.Context) {
	brand := c.Query("brand")
	if brand == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Brand parameter is required"})
		return
	}

	devices, err := services.SearchDevicesByBrand(brand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve devices"})
		return
	}

	c.JSON(http.StatusOK, devices)
}

// Implement other handlers (Get, Update, Delete, Search by brand, etc.)
