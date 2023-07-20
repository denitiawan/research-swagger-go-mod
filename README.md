## About

Research how to implement swagger on go (go.mod) using gin

## Librarries

```
# Web Framework & Router
github.com/gin-gonic/gin v1.9.1

# Validator DTO
github.com/go-playground/validator/v10 v10.14.1

# Log
github.com/rs/zerolog v1.29.1

# API Docs
github.com/swaggo/files v1.0.1
github.com/swaggo/gin-swagger v1.6.0

# Database & Driver
database/sql

# Database Migrate
github.com/rubenv/sql-migrate
go install github.com/rubenv/sql-migrate/...@latest
go mod download github.com/gobuffalo/packr/v2
go mod tidy

# Database Server
Mysql
```

## Run APP

[docker-compose.yml](https://github.com/denitiawan/research-swagger-gomod-gin/blob/main/docker-compose/docker-compose.yml)

```
# run syntax
docker-compose up -d

# docker will create docker image (swg_go_mysql)
host: localhost
port: 3399      
user: user
password: password
database : database
```

## postman

[postman](https://github.com/denitiawan/research-swagger-gomod-gin/blob/main/postman/Swagger-GO.18-v%201.0.0.postman_collection.json)

## How to setup swagger on (go.mod & GIN)
[Setup Swagger](https://github.com/denitiawan/research-swagger-gomod-gin/blob/main/readme/setup_swagger.md)
