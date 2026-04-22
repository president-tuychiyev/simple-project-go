# 🚀 Simple Go Project

Minimal and clean Go (Golang) backend project. API-first approach, fast startup, no unnecessary complexity.

---

## 📌 Overview

This project provides a simple REST API service. The structure is clean, scalable, and production-ready.

---

## ⚙️ Tech Stack

- Language: Go
- HTTP Router: net/http (can be replaced with Gin / Chi)
- Database: PostgreSQL (optional)
- Config: `.env`
- Build: Go Modules

---

## 📁 Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   └── model/
├── pkg/
│   └── utils/
├── configs/
│   └── config.go
├── .env
├── go.mod
└── README.md
```

---

## 🚀 Getting Started

### 1. Clone repository

```bash
git clone https://github.com/your-username/your-project.git
cd your-project
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Setup environment

Create a `.env` file:

```env
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=mydb
```

### 4. Run project

```bash
go run cmd/server/main.go
```

Or build and run:

```bash
go build -o app ./cmd/server
./app
```

---

## 📡 API Endpoints

| Method | Endpoint     | Description        |
|--------|-------------|--------------------|
| GET    | /health     | Health check       |
| GET    | /api/items  | Get all items      |
| POST   | /api/items  | Create new item    |

---

## 🧠 Architecture

- `handler` — HTTP layer (request/response)
- `service` — business logic
- `repository` — database layer
- `model` — data structures

Flow: `handler → service → repository`

---

## 🧪 Testing

```bash
go test ./...
```

---

## 📦 Build for Production

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/server
```

---

## 🐳 Docker (optional)

```dockerfile
FROM golang:1.22-alpine

WORKDIR /app
COPY . .

RUN go build -o app ./cmd/server

CMD ["./app"]
```

---

## 🛠️ Future Improvements

- JWT authentication
- Middleware (logging, rate limiting)
- Swagger documentation
- CI/CD pipeline
- Redis caching

---

## 📄 License

MIT License
