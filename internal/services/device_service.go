package services

import (
	"device-management-api/internal/models"
	"device-management-api/internal/repository"
)

func AddDevice(device *models.Device) (*models.Device, error) {
	return repository.AddDevice(device)
}

func GetDevice(id int) (*models.Device, error) {
	return repository.GetDevice(id)
}

func ListDevices() ([]*models.Device, error) {
	return repository.ListDevices()
}
func UpdateDevice(id int, updatedDevice *models.Device) (*models.Device, error) {
	return repository.UpdateDevice(id, updatedDevice)
}
func DeleteDevice(id int) error {
	return repository.DeleteDevice(id)
}
func SearchDevicesByBrand(brand string) ([]*models.Device, error) {
	return repository.SearchDevicesByBrand(brand)
}

// Implement other service functions (Get, Update, Delete, Search by brand, etc.)
