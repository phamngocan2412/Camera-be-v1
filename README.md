# Camera BE V1 - Backend API Hệ Thống Camera An Ninh

Backend API mạnh mẽ, có khả năng mở rộng cho hệ thống quản lý camera an ninh, được xây dựng với Golang sử dụng Gin web framework và GORM.

## 🚀 Tính Năng

### Xác Thực & Phân Quyền
- **JWT Authentication**: Hệ thống xác thực bảo mật với JSON Web Tokens
- **Đăng ký/Đăng nhập**: Quản lý ngườ dùng với mã hóa mật khẩu bcrypt
- **Xác minh OTP**: Gửi và xác minh mã OTP qua email/SMS
- **Quên mật khẩu**: Khôi phục mật khẩu với xác minh OTP

### Quản Lý Ngườ Dùng
- **CRUD Users**: Tạo, đọc, cập nhật, xóa ngườ dùng
- **Cập nhật hồ sơ**: Quản lý thông tin cá nhân
- **Đổi mật khẩu**: Cập nhật mật khẩu an toàn
- **Kiểm tra tồn tại**: Kiểm tra email/phone đã tồn tại

### API & Documentation
- **Swagger UI**: Tài liệu API tương tác tự động
- **RESTful API**: Chuẩn REST API đầy đủ
- **Validation**: Validate dữ liệu đầu vào chặt chẽ
- **Localization**: Hỗ trợ đa ngôn ngữ (i18n)

### Cơ Sở Dữ Liệu
- **PostgreSQL**: Cơ sở dữ liệu chính
- **Redis**: Cache và session management
- **Migrations**: Quản lý schema với golang-migrate
- **GORM ORM**: Tương tác database dễ dàng

### Logging & Monitoring
- **Zap Logger**: Logging có cấu trúc, hiệu suất cao
- **Multiple Levels**: Debug, Info, Warn, Error
- **Request ID**: Theo dõi request qua các layer

## 🛠 Công Nghệ Sử Dụng

- **Ngôn ngữ**: [Go](https://golang.org) (v1.24.0)
- **Web Framework**: [Gin Gonic](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io)
- **Database**: [PostgreSQL](https://www.postgresql.org)
- **Cache**: [Redis](https://redis.io)
- **Authentication**: [JWT](https://github.com/golang-jwt/jwt)
- **Config**: [Viper](https://github.com/spf13/viper)
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate)
- **Validation**: [phonenumbers](https://github.com/nyaruka/phonenumbers)
- **Documentation**: [Swagger](https://swaggo.github.io/swag/)
- **Logging**: [Zap](https://github.com/uber-go/zap)
- **Infrastructure**: [Docker](https://www.docker.com) & [Docker Compose](https://docs.docker.com/compose/)

## 📂 Cấu Trúc Dự Án

```
camera-be-v1/
├── cmd/
│   ├── api/              # Entry point API server
│   │   └── main.go
│   └── migrate/          # Database migration tool
│       └── main.go
├── configs/              # Cấu hình (YAML)
│   └── config.yaml
├── docs/                 # Swagger documentation (auto-generated)
│   ├── swagger.json
│   ├── swagger.yaml
│   └── docs.go
├── internal/             # Private application code
│   ├── config/          # Config loader
│   ├── handlers/        # HTTP handlers (controllers)
│   ├── middleware/      # JWT middleware, logging, etc.
│   ├── models/          # GORM database models
│   ├── platform/        # Platform abstractions
│   │   ├── db/         # Database connection
│   │   ├── i18n/       # Internationalization
│   │   └── logger/     # Logger setup
│   ├── repository/      # Data access layer
│   └── service/         # Business logic layer
├── migrations/           # SQL migration files
│   ├── 0001_create_users_table.up.sql
│   └── 0001_create_users_table.down.sql
├── Dockerfile           # Docker build file
├── docker-compose.yml   # Docker Compose config
├── go.mod              # Go module definition
├── go.sum              # Go dependencies checksum
├── run.sh              # Helper script for local dev
└── README.md           # This file
```

## 🏗 Cài Đặt Và Chạy

### Yêu Cầu

- Go 1.24.0+
- PostgreSQL 14+
- Redis 7+
- Docker & Docker Compose (optional)

### Option 1: Sử Dụng Docker (Khuyến nghị)

1. **Cài đặt Docker và Docker Compose**

2. **Build và chạy**:
   ```bash
   docker compose up --build
   ```

3. **Chạy migrations**:
   ```bash
   docker compose exec api ./api migrate
   ```

4. **Truy cập API**:
   - API: http://localhost:8080
   - Swagger UI: http://localhost:8080/swagger/index.html

### Option 2: Local Development

1. **Cài đặt dependencies**:
   - Go 1.24+
   - PostgreSQL
   - Redis

2. **Clone repository**:
   ```bash
   git clone <repository-url>
   cd camera-be-v1
   ```

3. **Cài đặt Go dependencies**:
   ```bash
   go mod download
   ```

4. **Cấu hình database**:
   
   Chỉnh sửa `configs/config.yaml`:
   ```yaml
   database:
     url: "host=localhost user=postgres password=yourpassword dbname=camera_security port=5432 sslmode=disable"
   
   redis:
     addr: "localhost:6379"
     password: ""
     db: 0
   ```

5. **Tạo database**:
   ```bash
   createdb camera_security
   ```

6. **Chạy migrations**:
   ```bash
   go run cmd/migrate/main.go up
   # Hoặc
   go run cmd/api/main.go migrate
   ```

7. **Chạy API server**:
   ```bash
   # Sử dụng script
   ./run.sh
   
   # Hoặc chạy trực tiếp
   go run cmd/api/main.go
   ```

8. **Truy cập**:
   - API: http://localhost:8080
   - Swagger: http://localhost:8080/swagger/index.html

## ⚙️ Cấu Hình

### File `configs/config.yaml`

```yaml
server:
  port: ":8080"                    # Port API server

jwt:
  secret: "your-secret-key"        # JWT secret key (đổi trong production!)
  expiration: 24h                  # Thờ gian hết hạn token

database:
  url: "host=postgres user=postgres password=postgres dbname=camera_security port=5432 sslmode=disable"

redis:
  addr: "redis:6379"              # Redis address
  password: ""                    # Redis password
  db: 0                           # Redis database

log:
  level: "debug"                  # debug, info, warn, error

smtp:
  host: "smtp.gmail.com"          # SMTP server
  port: 587                       # SMTP port
  email: "your-email@gmail.com"   # Email gửi OTP
  password: "your-app-password"   # App password
```

### Environment Variables

Có thể override config bằng environment variables:

```bash
export SERVER_PORT=:8080
export DATABASE_URL="host=localhost user=postgres password=secret dbname=camera"
export JWT_SECRET="your-super-secret-key"
export REDIS_ADDR="localhost:6379"
```

## 📚 API Documentation

### Authentication Endpoints

| Method | Endpoint | Mô tả |
|--------|----------|-------|
| POST | `/api/v1/auth/register` | Đăng ký tài khoản mới |
| POST | `/api/v1/auth/login` | Đăng nhập |
| POST | `/api/v1/auth/forgot-password` | Yêu cầu quên mật khẩu |
| POST | `/api/v1/auth/verify-otp` | Xác minh OTP |
| POST | `/api/v1/auth/reset-password` | Đặt lại mật khẩu |

### User Endpoints (Protected)

| Method | Endpoint | Mô tả |
|--------|----------|-------|
| GET | `/api/v1/users/me` | Lấy thông tin user hiện tại |
| PUT | `/api/v1/users/me` | Cập nhật hồ sơ |
| PUT | `/api/v1/users/me/password` | Đổi mật khẩu |
| GET | `/api/v1/users/check-email` | Kiểm tra email tồn tại |
| GET | `/api/v1/users/check-phone` | Kiểm tra phone tồn tại |

### Truy Cập Swagger UI

Khi server đang chạy:
```
http://localhost:8080/swagger/index.html
```

## 🧪 Testing

```bash
# Chạy tất cả tests
go test ./...

# Chạy test với coverage
go test -cover ./...

# Chạy test cụ thể
go test ./internal/service/...

# Benchmark
go test -bench=. ./...
```

## 🗄 Database Migrations

### Tạo Migration Mới

```bash
# Tạo file migration mới
migrate create -ext sql -dir migrations -seq create_cameras_table
```

### Chạy Migrations

```bash
# Up - chạy migrations
go run cmd/migrate/main.go up

# Down - rollback
go run cmd/migrate/main.go down

# Version - xem version hiện tại
go run cmd/migrate/main.go version

# Force - force version cụ thể
go run cmd/migrate/main.go force 1
```

## 📝 Coding Conventions

### Naming Conventions
- **Files**: `snake_case.go`
- **Packages**: `lowercase`
- **Interfaces**: `PascalCase` (ví dụ: `UserRepository`)
- **Structs**: `PascalCase` (ví dụ: `AuthService`)
- **Functions**: `PascalCase` cho exported, `camelCase` cho private
- **Variables**: `camelCase`
- **Constants**: `SCREAMING_SNAKE_CASE`

### Layer Architecture
```
HTTP Request
    ↓
Handler (Input validation, HTTP response)
    ↓
Service (Business logic)
    ↓
Repository (Data access)
    ↓
Database
```

### Error Handling
- Sử dụng custom error types
- Không bao giờ swallow errors
- Log errors với context đầy đủ
- Trả về HTTP status codes phù hợp

## 🔒 Security

- **JWT**: Sử dụng HS256 algorithm
- **Password**: Bcrypt với cost 10
- **Input Validation**: Validate tất cả inputs
- **CORS**: Cấu hình CORS phù hợp
- **Rate Limiting**: Giới hạn requests
- **SQL Injection**: Sử dụng GORM parameterized queries

## 🚀 Deployment

### Docker Production

```bash
# Build image
docker build -t camera-be:latest .

# Run container
docker run -d \
  -p 8080:8080 \
  -e DATABASE_URL="postgresql://..." \
  -e JWT_SECRET="..." \
  --name camera-be \
  camera-be:latest
```

### Kubernetes (ví dụ)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: camera-be
spec:
  replicas: 3
  selector:
    matchLabels:
      app: camera-be
  template:
    metadata:
      labels:
        app: camera-be
    spec:
      containers:
      - name: api
        image: camera-be:latest
        ports:
        - containerPort: 8080
        env:
        - name: DATABASE_URL
          valueFrom:
            secretKeyRef:
              name: db-secret
              key: url
```

## 📊 Monitoring & Logging

### Log Levels
- **Debug**: Chi tiết phát triển
- **Info**: Thông tin hoạt động bình thường
- **Warn**: Cảnh báo không nghiêm trọng
- **Error**: Lỗi cần xử lý

### Structured Logging
```json
{
  "level": "info",
  "timestamp": "2026-02-06T10:30:00Z",
  "caller": "auth_handler.go:45",
  "msg": "user login successful",
  "user_id": "123",
  "ip": "192.168.1.1"
}
```

## 🤝 Đóng Góp

1. Fork repository
2. Tạo feature branch: `git checkout -b feature/tinh-nang-moi`
3. Commit changes: `git commit -m 'feat: thêm tính năng mới'`
4. Push lên branch: `git push origin feature/tinh-nang-moi`
5. Tạo Pull Request

### Commit Convention
- `feat:` - Tính năng mới
- `fix:` - Sửa lỗi
- `docs:` - Cập nhật tài liệu
- `refactor:` - Tái cấu trúc code
- `test:` - Thêm/cập nhật tests
- `chore:` - Công việc bảo trì

## 📄 License

Dự án này được cấp phép theo [MIT License](LICENSE).

## 📞 Liên Hệ

- **Author**: Your Name
- **Email**: your.email@example.com
- **GitHub**: [@username](https://github.com/username)
- **Project Link**: [camera-be-v1](https://github.com/username/camera-be-v1)

## 🙏 Cảm Ơn

- [Gin Web Framework](https://gin-gonic.com)
- [GORM](https://gorm.io)
- [Swaggo](https://github.com/swaggo/swag)
- [Uber Zap](https://github.com/uber-go/zap)

---

**Lưu ý bảo mật**: 
- Không bao giờ commit file `.env` hoặc secrets
- Luôn đổi JWT secret trong production
- Sử dụng HTTPS trong production
- Cập nhật dependencies thường xuyên
