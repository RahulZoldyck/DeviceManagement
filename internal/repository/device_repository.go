package repository

import (
	"device-management-api/internal/models"
	"errors"
	"sync"
	"time"
)

var (
	devices    = make(map[int]*models.Device)
	idCounter  = 1
	devicesMux sync.Mutex
)

func AddDevice(device *models.Device) (*models.Device, error) {
	devicesMux.Lock()
	defer devicesMux.Unlock()

	device.ID = idCounter
	device.CreatedAt = time.Now()
	devices[idCounter] = device
	idCounter++

	return device, nil
}
func GetDevice(id int) (*models.Device, error) {
	devicesMux.Lock()
	defer devicesMux.Unlock()

	device, exists := devices[id]
	if !exists {
		return nil, errors.New("device not found")
	}

	return device, nil
}
func ListDevices() ([]*models.Device, error) {
	devicesMux.Lock()
	defer devicesMux.Unlock()

	deviceList := make([]*models.Device, 0, len(devices))
	for _, device := range devices {
		deviceList = append(deviceList, device)
	}

	return deviceList, nil
}
func UpdateDevice(id int, updatedDevice *models.Device) (*models.Device, error) {
	devicesMux.Lock()
	defer devicesMux.Unlock()

	device, exists := devices[id]
	if !exists {
		return nil, errors.New("device not found")
	}

	// Update device fields
	device.Name = updatedDevice.Name
	device.Brand = updatedDevice.Brand
	device.CreatedAt = time.Now() // Or keep the original creation time if preferred

	return device, nil
}
func DeleteDevice(id int) error {
	devicesMux.Lock()
	defer devicesMux.Unlock()

	if _, exists := devices[id]; !exists {
		return errors.New("device not found")
	}

	delete(devices, id)
	return nil
}
func SearchDevicesByBrand(brand string) ([]*models.Device, error) {
	devicesMux.Lock()
	defer devicesMux.Unlock()

	var result []*models.Device
	for _, device := range devices {
		if device.Brand == brand {
			result = append(result, device)
		}
	}

	return result, nil
}
