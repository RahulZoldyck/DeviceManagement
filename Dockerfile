FROM golang:1.19-alpine

WORKDIR /app

# Install bash if you want to use bash
RUN apk add --no-cache bash

COPY . .
RUN chmod +x wait-for-it.sh  # Make the script executable

RUN go mod download
RUN go build -o main cmd/main.go

EXPOSE 8080

# Use bash instead of sh (if installed)
CMD ["bash", "./wait-for-it.sh", "mysql:3306", "--", "./main"]
