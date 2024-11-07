#!/bin/bash

# Build and run the Docker container
docker build -t device-management-api .
docker run -p 8080:8080 device-management-api
