package main

import (
	"device-management-api/internal/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080") // Start server on port 8080
}
