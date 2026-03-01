# Go REST API - Students API

A REST API built with Go for managing student records using SQLite database.

## Features

- Pure Go implementation (no CGO required)
- SQLite database using `modernc.org/sqlite` (pure Go driver)
- RESTful API endpoints for student management

## Building

The application can be built without CGO:

```bash
go build -o students-api.exe ./cmd/students-api
```

Or explicitly with CGO disabled:

```bash
set CGO_ENABLED=0
go build -o students-api.exe ./cmd/students-api
```

## Running

```bash
./students-api.exe
```

## Dependencies

- `modernc.org/sqlite` - Pure Go SQLite driver (no CGO required)
- Other dependencies listed in `go.mod`

## Note

This project previously used `github.com/mattn/go-sqlite3` which required CGO. It has been migrated to `modernc.org/sqlite` for better portability and to avoid CGO dependencies.


## Run CMD

go run cmd/students-api/main.go -config config/local.yaml