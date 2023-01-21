# Gin Clean Architecture

This project was created to learn golang with gin framework

## How To Run

1. Run docker compose with command `docker compose up`
2. Run Migration DB `migrate -database "mysql://root:root@tcp(localhost:3306)/gin_clean_architecture" -path db/migrations up`
3. Run application with command `go run main.go`

## Feature

- [x] Database sqlx
- [x] Database Relational
- [x] Json Validation
- [ ] JWT Security
- [x] Database migration
- [x] Docker Support
- [ ] Open API / Swagger
- [x] Integration Test
- [x] Http Client
- [ ] Error Handling
- [x] Logging
- [x] Cache