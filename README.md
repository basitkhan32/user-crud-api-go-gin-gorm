# CRUD API with Gin Framework and GORM

A RESTful API built with Go, using the Gin web framework and GORM ORM for PostgreSQL database operations. This project implements complete CRUD (Create, Read, Update, Delete) operations for user management.

## Features

- **RESTful API** endpoints for user management
- **PostgreSQL** database integration using GORM
- **Gin Framework** for high-performance routing
- **Environment variables** configuration with godotenv
- **Clean architecture** with separated concerns (models, controllers, routes, connections)
- **Database migrations** support

## Prerequisites

Before running this application, make sure you have the following installed:

- Go 1.25.5 or higher
- PostgreSQL database
- Git

## Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/basitkhan32/crud-api-go-gin.git
   cd crud-api-go-gin
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   
   Create a `.env` file in the root directory with the following variables:
   ```env
   DB_STRING=host=localhost user=your_user password=your_password dbname=your_db port=5432 sslmode=disable
   PORT=8080
   ```

4. **Run database migrations** (if needed)
   
   The application uses GORM's AutoMigrate feature. You can run migrations using the migration package.

## Running the Application

1. **Build and run**
   ```bash
   go run cmd/main.go
   ```

2. **Or build the executable**
   ```bash
   go build -o main cmd/main.go
   ./main
   ```

The server will start on the port specified in your `.env` file (default: 8080).

## 📡 API Endpoints

### Base URL
```
http://localhost:8080
```

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/create` | Create a new user |
| GET | `/read` | Get all users |
| GET | `/read/:id` | Get a user by ID |
| PUT | `/update/:id` | Update a user by ID |
| DELETE | `/delete/:id` | Delete a user by ID |

### Request/Response Examples

#### 1. Create User
**Request:**
```bash
POST /create
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 30
}
```

**Response:**
```json
{
  "message": "User created successfully",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 30
  }
}
```

#### 2. Get All Users
**Request:**
```bash
GET /read
```

**Response:**
```json
{
  "message": "Users fetched successfully",
  "users": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "age": 30
    }
  ]
}
```

#### 3. Get User by ID
**Request:**
```bash
GET /read/1
```

**Response:**
```json
{
  "message": "User with id 1 fetched successfully",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 30
  }
}
```

#### 4. Update User
**Request:**
```bash
PUT /update/1
Content-Type: application/json

{
  "name": "Jane Doe",
  "email": "jane.doe@example.com",
  "age": 28
}
```

**Response:**
```json
{
  "message": "User updated successfully",
  "user": {
    "name": "Jane Doe",
    "email": "jane.doe@example.com",
    "age": 28
  }
}
```

#### 5. Delete User
**Request:**
```bash
DELETE /delete/1
```

**Response:**
```json
{
  "message": "User deleted successfully with id:1",
  "user": {}
}
```

## Project Structure

```
crud-api-gin/
├── cmd/
│   └── main.go              # Application entry point
├── connections/
│   └── db_connection.go     # Database connection setup
├── controllers/
│   ├── create_controller.go # Create user handler
│   ├── read_controller.go   # Read users handlers
│   ├── update_controller.go # Update user handler
│   └── delete_controller.go # Delete user handler
├── migrations/
│   └── db_migration.go      # Database migration logic
├── models/
│   └── user_model.go        # User data model
├── routes/
│   └── routes.go            # API routes definition
├── utils/
│   └── env.go               # Environment variables loader
├── tmp/
│   └── main                 # Compiled binary (dev)
├── go.mod                   # Go module dependencies
├── go.sum                   # Dependencies checksums
├── .env                     # Environment variables (not in repo)
└── README.md                # This file
```

## Database Schema

### User Model

| Field | Type | Description |
|-------|------|-------------|
| ID | uint | Primary key (auto-increment) |
| Name | string | User's name |
| Email | *string | User's email (nullable) |
| Age | uint8 | User's age |
| CreatedAt | time.Time | Record creation timestamp (GORM) |
| UpdatedAt | time.Time | Record update timestamp (GORM) |
| DeletedAt | gorm.DeletedAt | Soft delete timestamp (GORM) |

## Technologies Used

- **[Go](https://golang.org/)** - Programming language
- **[Gin](https://github.com/gin-gonic/gin)** - Web framework
- **[GORM](https://gorm.io/)** - ORM library
- **[PostgreSQL](https://www.postgresql.org/)** - Database
- **[godotenv](https://github.com/joho/godotenv)** - Environment variables management

## Development

### Adding New Routes

1. Create a new controller in `controllers/`
2. Define your handler function
3. Register the route in `routes/routes.go`

### Running Migrations

To run database migrations manually:

```go
import "github.com/basitkhan32/crud-api-go-gin/migrations"

migrations.DBMigration()
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is open source and available under the MIT License.

## 👤 Author

**Basit Khan**

- GitHub: [@basitkhan32](https://github.com/basitkhan32)

## Acknowledgments

- Gin Web Framework team
- GORM team
- Go community

---

**Note:** Remember to never commit your `.env` file or any sensitive credentials to version control. Add `.env` to your `.gitignore` file.
