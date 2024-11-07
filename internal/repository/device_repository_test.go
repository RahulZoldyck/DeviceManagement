package repository

import (
	"device-management-api/internal/models"
	"testing"
)

func TestAddDevice(t *testing.T) {
	device := &models.Device{Name: "Phone", Brand: "BrandX"}
	addedDevice, err := AddDevice(device)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if addedDevice.ID == 0 {
		t.Fatalf("Expected valid device ID, got 0")
	}
}
