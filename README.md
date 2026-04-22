🚀 Simple Go Project

Minimal va toza arxitekturaga ega Go backend loyiha. API-first yondashuv, tez ishga tushadi, ortiqcha “enterprise drama” yo‘q.

📌 Overview

Ushbu loyiha REST API xizmatini taqdim etadi. Strukturasi oddiy, kengaytirish oson, va production’ga moslab yozilgan.

⚙️ Tech Stack
Language: Go
HTTP Router: net/http (yoki chi / gin qo‘shishingiz mumkin)
Database: PostgreSQL (optional)
Config: .env
Build: Go modules
📁 Project Structure
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
🚀 Getting Started
1. Clone repository
git clone https://github.com/your-username/your-project.git
cd your-project
2. Install dependencies
go mod tidy
3. Setup environment

.env fayl yarating:

APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=mydb
4. Run project
go run cmd/server/main.go

Yoki build qilib:

go build -o app ./cmd/server
./app
📡 API Endpoints
Method	Endpoint	Description
GET	/health	Health check
GET	/api/items	Get all items
POST	/api/items	Create new item
🧠 Architecture
handler — HTTP layer (request/response)
service — business logic
repository — database access
model — data structures

Clean separation: handler → service → repository

🧪 Testing
go test ./...
📦 Build for Production
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/server
🐳 Docker (optional)
FROM golang:1.22-alpine

WORKDIR /app
COPY . .

RUN go build -o app ./cmd/server

CMD ["./app"]
🛠️ Future Improvements
JWT authentication
Middleware (logging, rate limit)
Swagger docs
CI/CD pipeline
Caching (Redis)
📄 License

MIT License
