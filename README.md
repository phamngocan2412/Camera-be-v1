# Camera BE V1 - Security Camera Backend API

A robust, scalable backend service for security camera management built with Golang using the Gin web framework and GORM.

## 🚀 Features

- **Authentication**: JWT-based secure authentication system (Register, Login).
- **User Management**: Profile updates, password changes, and user data management.
- **API Documentation**: Interactive Swagger UI for easy endpoint testing.
- **Database Migrations**: Automatic schema management with `golang-migrate`.
- **Performance**: Redis integration for caching and high performance.
- **Logging**: High-speed, structured logging using `zap`.

## 🛠 Tech Stack

- **Language**: [Go](https://golang.org) (v1.24.0)
- **Web Framework**: [Gin Gonic](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io)
- **Database**: [PostgreSQL](https://www.postgresql.org)
- **Cache**: [Redis](https://redis.io)
- **Config**: [Viper](https://github.com/spf13/viper)
- **Documentation**: [Swagger](https://swaggo.github.io/swag/)
- **Infrastructure**: [Docker](https://www.docker.com) & [Docker Compose](https://docs.docker.com/compose/)

## 📂 Project Structure

```text
.
├── cmd/api/            # Entry point of the application
├── configs/            # Configuration files (YAML)
├── docs/               # Auto-generated Swagger documentation
├── internal/           # Private application and library code
│   ├── handlers/       # API route handlers
│   ├── models/         # GORM database models
│   ├── repository/     # Database access layer
│   └── service/        # Business logic layer
├── migrations/         # SQL migration files
└── run.sh              # Helper script for local development
```

## 🏗 Setup & Installation

### Option 1: Using Docker (Recommended)

1. **Install Docker and Docker Compose**.
2. **Build and run**:
   ```bash
   docker compose up --build
   ```

### Option 2: Local Development

1. **Prerequisites**: Install Go 1.24+, PostgreSQL, and Redis.
2. **Configure**: Update `configs/config.yaml` with your local database credentials.
3. **Run migrations**:
   ```bash
   go run cmd/api/main.go migrate
   ```
4. **Run the API**:
   ```bash
   ./run.sh
   # OR
   go run cmd/api/main.go
   ```

## 📚 API Documentation

Once the server is running, you can access the Swagger UI at:
`http://localhost:8080/swagger/index.html`

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
