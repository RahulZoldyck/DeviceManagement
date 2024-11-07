package repository

import (
	"device-management-api/internal/models"
	"gorm.io/gorm"
)

type DeviceRepository struct {
	DB *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) *DeviceRepository {
	return &DeviceRepository{DB: db}
}

func (repo *DeviceRepository) AddDevice(device *models.Device) (*models.Device, error) {
	if err := repo.DB.Create(device).Error; err != nil {
		return nil, err
	}
	return device, nil
}

func (repo *DeviceRepository) GetDevice(id uint) (*models.Device, error) {
	var device models.Device
	if err := repo.DB.First(&device, id).Error; err != nil {
		return nil, err
	}
	return &device, nil
}

func (repo *DeviceRepository) ListDevices() ([]*models.Device, error) {
	var devices []*models.Device
	if err := repo.DB.Find(&devices).Error; err != nil {
		return nil, err
	}
	return devices, nil
}

func (repo *DeviceRepository) UpdateDevice(id uint, updatedDevice *models.Device) (*models.Device, error) {
	var device models.Device
	if err := repo.DB.First(&device, id).Error; err != nil {
		return nil, err
	}

	device.Name = updatedDevice.Name
	device.Brand = updatedDevice.Brand

	if err := repo.DB.Save(&device).Error; err != nil {
		return nil, err
	}
	return &device, nil
}

func (repo *DeviceRepository) DeleteDevice(id uint) error {
	if err := repo.DB.Delete(&models.Device{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (repo *DeviceRepository) SearchDevicesByBrand(brand string) ([]*models.Device, error) {
	var devices []*models.Device
	if err := repo.DB.Where("brand = ?", brand).Find(&devices).Error; err != nil {
		return nil, err
	}
	return devices, nil
}
