
# Device Management API

Device Management API is a RESTful service built with Go, Gin, and GORM that manages information about devices, allowing CRUD operations such as creating, reading, updating, and deleting device records. It also supports querying devices by brand.

## Features

- Add a new device
- Retrieve a device by ID
- List all devices, with optional search by brand
- Update a device by ID
- Delete a device by ID

## Technologies Used

- **Golang**: Main programming language
- **Gin**: Web framework for handling HTTP requests
- **GORM**: ORM library for database operations
- **MySQL**: Relational database for storing device data
- **Docker**: Containerization for easy deployment and development

## Project Structure

```
DeviceManagement/
├── cmd/
│   └── main.go              # Main entry point
├── config/
│   └── config.go            # Configuration handling
├── internal/
│   ├── controllers/         # Controller logic
│   ├── models/              # Database models
│   ├── repository/          # Database operations
│   └── routes/              # API routes setup
├── Dockerfile               # Docker image setup
├── docker-compose.yml       # Docker Compose configuration
├── .env                     # Environment variables
└── README.md                # Project documentation
```

## Getting Started

### Prerequisites

Ensure you have the following installed:
- Docker
- Docker Compose
- Golang (for local development)

### Installation and Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/RahulZoldyck/DeviceManagement.git
   cd DeviceManagement
   ```

2. **Environment Variables**: Configure database and app settings in `.env`.
   ```
   PORT=8080
   DB_HOST=mysql
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=pass
   DB_NAME=devicedb
   ```

3. **Run Docker Compose**:
   Start the application and database services with Docker Compose.
   ```bash
   docker-compose up --build
   ```

   This will build and run both the MySQL and API services. The API will be accessible at `http://localhost:8080`.

4. **Database Initialization**:
   The application will automatically set up tables with GORM's auto-migration feature.

### API Endpoints

- **Create Device**: `POST /devices`
- **Get Device by ID**: `GET /devices/:id`
- **List All Devices**: `GET /devices`
- **Search Devices by Brand**: `GET /devices?brand={brand}`
- **Update Device**: `PUT /devices/:id`
- **Delete Device**: `DELETE /devices/:id`

### Example Requests

Here’s a sample `curl` request for each endpoint:

1. **Create a Device**
   ```bash
   curl -X POST http://localhost:8080/devices -H "Content-Type: application/json" -d '{
     "name": "Smartphone",
     "brand": "Apple",
     "price": 999
   }'
   ```

2. **Get a Device by ID**
   ```bash
   curl -X GET http://localhost:8080/devices/1
   ```

3. **List All Devices**
   ```bash
   curl -X GET http://localhost:8080/devices
   ```

4. **Search Devices by Brand**
   ```bash
   curl -X GET "http://localhost:8080/devices?brand=Apple"
   ```

5. **Update a Device**
   ```bash
   curl -X PUT http://localhost:8080/devices/1 -H "Content-Type: application/json" -d '{
     "name": "Updated Device",
     "brand": "Samsung",
     "price": 850
   }'
   ```

6. **Delete a Device**
   ```bash
   curl -X DELETE http://localhost:8080/devices/1
   ```

## Troubleshooting

- **Database Connection Error**: Ensure your `.env` database credentials match the MySQL setup in `docker-compose.yml`.
- **Port Conflicts**: Change the port in `.env` if `8080` is already in use.
- **Docker Issues**: Use `docker-compose down` and `docker-compose up --build` to restart from scratch.
