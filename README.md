# User CRUD API — Go, Gin, GORM, PostgreSQL

A RESTful API built with **Go**, using the **Gin** web framework and **GORM** ORM against a **PostgreSQL** database. Implements full user management — authentication (register / login / password update) and protected CRUD operations guarded by an `X-API-KEY` header middleware.

---

## Stack

| | Technology |
|---|---|
| Language | Go 1.25 |
| Framework | Gin v1.11 |
| ORM | GORM v1.31 + pgx v5 driver |
| Database | PostgreSQL |
| Config | godotenv |
| Dev reload | Air (hot reload via `.air.toml`) |

---

## Project Structure

```
.
├── cmd/
│   └── main.go               # Entry point — boots env, DB, migrations, router
├── db/
│   ├── connections.go        # Opens PostgreSQL connection; auto-creates DB if missing
│   └── migrations.go         # GORM AutoMigrate for all models
├── modules/
│   ├── auth/
│   │   ├── controller.go     # Register, Login, UpdatePassword handlers
│   │   ├── model.go          # Request body structs with binding validation
│   │   └── routes.go         # Mounts /auth/* routes
│   └── user/
│       ├── controller.go     # CreateUser, GetUser, GetUserID, UpdateUser, DeleteUser
│       ├── middleware.go     # X-API-KEY guard — aborts 401 if header missing
│       ├── model.go          # User struct (soft-delete via gorm.DeletedAt)
│       └── routes.go         # Public + protected route groups
├── routes/
│   └── routes.go             # Top-level router — wires public and protected groups
├── utils/
│   ├── env.go                # Loads .env via godotenv
│   └── errors.go             # ErrBadRequest / ErrServerCrash helpers
├── .air.toml                 # Air hot-reload config
├── .env.example              # Environment variable reference
├── go.mod
└── go.sum
```

---

## API Endpoints

### Auth — Public (no authentication required)

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/auth/register` | Register a new user |
| `POST` | `/auth/login` | Login with email and password |
| `POST` | `/auth/update-password` | Update password for an existing user |

### User — Public

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/user/create` | Create a user (no API key needed) |

### User — Protected (requires `X-API-KEY` header)

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/user/` | Create a user |
| `GET` | `/user/all` | Get all users |
| `GET` | `/user/:id` | Get a single user by ID |
| `PUT` | `/user/:id` | Update a user by ID |
| `DELETE` | `/user/:id` | Soft-delete a user by ID |

> All protected routes require the header: `X-API-KEY: <your-key>`
> Requests without it receive `401 Unauthorized`.

---

## Request & Response Examples

### POST `/auth/register`
```json
// Request
{
  "name": "John Doe",
  "email": "john@example.com",
  "age": 28,
  "password": "secret123",
  "confirm_password": "secret123"
}

// Response 201
{
  "message": "User registered successfully",
  "user": { "id": 1, "name": "John Doe", "email": "john@example.com", "age": 28, "created_at": "..." }
}
```

### POST `/auth/login`
```json
// Request
{ "email": "john@example.com", "password": "secret123" }

// Response 200
{ "message": "User logged in successfully", "user": { "id": 1, "name": "John Doe", ... } }
```

### POST `/auth/update-password`
```json
// Request
{ "email": "john@example.com", "old_password": "secret123", "new_password": "newpass456" }

// Response 200
{ "message": "Password updated successfully", "user": { ... } }
```

### GET `/user/all` *(X-API-KEY required)*
```json
// Response 200
{
  "message": "Users fetched successfully",
  "users": [
    { "id": 1, "name": "John Doe", "email": "john@example.com", "age": 28, "created_at": "..." }
  ]
}
```

### PUT `/user/:id` *(X-API-KEY required)*
```json
// Request
{ "name": "Jane Doe", "age": 30 }

// Response 200
{ "message": "User updated successfully", "user": { "id": 1, "name": "Jane Doe", "age": 30, ... } }
```

---

## Data Model

### User

| Field | Type | Notes |
|-------|------|-------|
| `id` | `uint` | Primary key, auto-increment |
| `name` | `string` | — |
| `email` | `string` | Required, validated as email |
| `age` | `uint8` | — |
| `password` | `string` | Omitted from all JSON responses (`json:"-"`) |
| `created_at` | `time.Time` | Set by GORM |
| `updated_at` | `time.Time` | Set by GORM |
| `deleted_at` | `gorm.DeletedAt` | Soft delete — records are never hard-deleted |

---

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL running locally (or a connection string to a remote instance)

### 1. Clone & install dependencies

```bash
git clone https://github.com/basitkhan32/user-crud-api-go-gin-gorm.git
cd user-crud-api-go-gin-gorm
go mod download
```

### 2. Configure environment

```bash
cp .env.example .env
```

Edit `.env`:

```env
PORT=8080
GIN_MODE=release
DB_STRING=postgresql://<user>:<password>@localhost:5432/<dbname>?sslmode=disable
DB_NAME=<dbname>
```

### 3. Run

```bash
# Development with hot reload (requires Air)
go install github.com/air-verse/air@latest
air

# Or run directly
go run cmd/main.go
```

The server starts on `http://localhost:8080` (or the `PORT` in your `.env`).
GORM runs `AutoMigrate` on startup — the `users` table is created automatically.

---

## Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `PORT` | Port the server listens on | `8080` |
| `GIN_MODE` | `debug` or `release` | `release` |
| `DB_STRING` | Full PostgreSQL connection string | `postgresql://user:pass@localhost:5432/mydb?sslmode=disable` |
| `DB_NAME` | Database name (used to auto-create DB if missing) | `mydb` |

---

## Author

**Abdul Basit Khan** — [@basitkhan32](https://github.com/basitkhan32)
