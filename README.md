# REST API in Go

A comprehensive REST API built with Go, featuring complete event management with JWT authentication, user registration, and event registration system using SQLite database integration.

Based on the Udemy course by [Maximilian Schwarzmüller](https://www.udemy.com/user/maximilian-schwarzmuller/)

## 🚀 Features

### Core Functionality
- **Event Management**: Full CRUD operations for events with ownership validation
- **User Authentication**: JWT-based authentication with 2-hour token expiration
- **User Registration**: Account creation with bcrypt password hashing (cost factor 14)
- **Event Registration**: Users can register/unregister for events they don't own
- **Authorization**: Ownership-based access control for event modifications

### Technical Features
- **Database Integration**: SQLite with automatic table creation (users, events, registrations)
- **RESTful Design**: Clean API endpoints following REST principles with proper HTTP status codes
- **Middleware Authentication**: JWT token validation with context-based user identification
- **Input Validation**: Comprehensive request validation using Gin binding tags
- **Password Security**: Production-grade bcrypt hashing for secure credential storage
- **Route Groups**: Organized protected and public endpoints with middleware separation

## 🏗️ Architecture

This is a **complete 4-layer Go REST API** with the following structure:

- **`main.go`**: Entry point - initializes database and starts Gin server on `:8080`
- **`db/`**: Global SQLite connection (`db.DB`) with auto table creation
- **`models/`**: Data models with pointer receiver methods (e.g., `func (e *Event) Save()`)
- **`routes/`**: Gin handlers split by domain (`events.go`, `users.go`, `register.go`) + registration
- **`middlewares/`**: Authentication middleware for JWT token validation
- **`utils/`**: Shared utilities (bcrypt password hashing, JWT token management)

## 📁 Project Structure

```
rest_api_golang/
├── main.go                 # Application entry point
├── api.db                  # SQLite database (auto-generated)
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── api-test/               # HTTP test files with JWT tokens
│   ├── create-event.http
│   ├── create-user.http
│   ├── delete-event.http
│   ├── get-events.http
│   ├── get-single-event.http
│   ├── login-user.http
│   ├── register-for-event.http
│   ├── signup-user.http
│   └── update-event.http
├── db/
│   └── db.go              # Database connection & table creation
├── models/
│   ├── event.go           # Event model & methods
│   └── user.go            # User model & methods
├── routes/
│   ├── routes.go          # Route registration & middleware setup
│   ├── events.go          # Event CRUD handlers with ownership validation
│   ├── register.go        # Event registration handlers
│   └── users.go           # User authentication handlers
├── middlewares/
│   └── auth.go            # JWT authentication middleware
└── utils/
    ├── hash.go            # Bcrypt password hashing
    └── jwt.go             # JWT token generation & validation
```

## 🛠️ Technologies Used

- **Go 1.24.4**: Programming language
- **Gin Web Framework**: HTTP router and middleware with route groups
- **SQLite**: Embedded database with foreign key relationships
- **golang.org/x/crypto/bcrypt**: Password hashing (cost factor 14)
- **github.com/golang-jwt/jwt/v5**: JWT token generation and validation
- **database/sql**: Go standard database interface with prepared statements

## 🚦 Getting Started

### Prerequisites

- Go 1.19 or higher
- Git

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/PaulFWatts/rest_api_golang.git
   cd rest_api_golang
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## 📚 API Endpoints

### Public Endpoints (No Authentication Required)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/events` | Get all events |
| GET | `/events/:id` | Get a specific event |
| POST | `/signup` | Register a new user |
| POST | `/login` | User login (returns JWT token) |

### Protected Endpoints (JWT Authentication Required)

| Method | Endpoint | Description | Authorization |
|--------|----------|-------------|---------------|
| POST | `/events` | Create a new event | Any authenticated user |
| PUT | `/events/:id` | Update an event | Event owner only |
| DELETE | `/events/:id` | Delete an event | Event owner only |
| POST | `/events/:id/register` | Register for an event | Any authenticated user (except event owner) |
| DELETE | `/events/:id/register` | Cancel event registration | Any authenticated user with existing registration |

### Authentication Header
For protected endpoints, include the JWT token in the Authorization header:
```
Authorization: your-jwt-token-here
```

## 🧪 Testing

Use the provided HTTP test files in the `api-test/` directory with VS Code REST Client extension or similar tools:

### Basic Workflow
1. **Create a user**: `api-test/signup-user.http`
2. **Login**: `api-test/login-user.http` (saves JWT token)
3. **Create an event**: `api-test/create-event.http` (requires JWT)
4. **Get events**: `api-test/get-events.http`
5. **Register for event**: `api-test/register-for-event.http` (requires JWT)
6. **Update event**: `api-test/update-event.http` (requires JWT + ownership)

### Test Files Available
- `signup-user.http` - User registration
- `login-user.http` - User authentication
- `create-event.http` - Event creation (authenticated)
- `get-events.http` - List all events
- `get-single-event.http` - Get specific event
- `update-event.http` - Update event (authenticated + owner)
- `delete-event.http` - Delete event (authenticated + owner)
- `register-for-event.http` - Register for event (authenticated)

## 🔧 Database Schema

### Users Table
- `id` (INTEGER, PRIMARY KEY, AUTOINCREMENT)
- `email` (TEXT, UNIQUE, NOT NULL)
- `password` (TEXT, NOT NULL) - Bcrypt hashed with cost factor 14

### Events Table
- `id` (INTEGER, PRIMARY KEY, AUTOINCREMENT)
- `name` (TEXT, NOT NULL)
- `description` (TEXT, NOT NULL)
- `location` (TEXT, NOT NULL)
- `dateTime` (DATETIME, NOT NULL)
- `user_id` (INTEGER, FOREIGN KEY REFERENCES users(id))

### Registrations Table (Many-to-Many Relationship)
- `id` (INTEGER, PRIMARY KEY, AUTOINCREMENT)
- `event_id` (INTEGER, FOREIGN KEY REFERENCES events(id))
- `user_id` (INTEGER, FOREIGN KEY REFERENCES users(id))
- UNIQUE constraint on (event_id, user_id) to prevent duplicate registrations

## 🔒 Security Features

### Authentication & Authorization
- JWT tokens with 2-hour expiration time
- Middleware-based authentication for protected routes
- User context injection for authenticated requests
- Event ownership validation (users can only modify their own events)
- Registration validation (users cannot register for their own events)

### Data Security
- Bcrypt password hashing with cost factor 14 (production-grade)
- Email uniqueness enforced at database level
- Comprehensive input validation using Gin binding tags
- Prepared statements to prevent SQL injection attacks
- Secure JWT signing with HMAC-SHA256

## ✅ Project Status: COMPLETED

This REST API project is now **fully implemented** with all planned features:

### ✅ Completed Features
- **User Management**: Registration, login, JWT authentication
- **Event Management**: Complete CRUD with ownership validation
- **Event Registration**: Users can register/unregister for events
- **Authentication**: JWT middleware with 2-hour token expiration
- **Authorization**: Ownership-based access control
- **Database**: Three-table schema with foreign key relationships
- **Security**: Bcrypt hashing, input validation, prepared statements
- **Testing**: Complete HTTP test suite for all endpoints

### 🏆 Architecture Achievements
- Clean separation of concerns across 6 packages
- Middleware-based authentication system
- RESTful API design with proper HTTP status codes
- Database abstraction with model methods
- Comprehensive error handling and validation
- Production-ready security practices
- 🔲 Pagination for event listings
- 🔲 Input sanitization and validation improvements

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is part of an educational course and is intended for learning purposes.

## 🙏 Acknowledgments

- [Maximilian Schwarzmüller](https://www.udemy.com/user/maximilian-schwarzmuller/) for the excellent Udemy course
- The Go community for amazing tools and libraries

