# REST API in Go

A comprehensive REST API built with Go, featuring event management with user authentication and SQLite database integration.

Based on the Udemy course by [Maximilian Schwarzmüller](https://www.udemy.com/user/maximilian-schwarzmuller/)

## 🚀 Features

- **Event Management**: Create, read, update, and delete events
- **User Authentication**: User registration and login with password hashing
- **Database Integration**: SQLite database with automatic table creation
- **RESTful Design**: Clean API endpoints following REST principles
- **Input Validation**: Request validation using Gin binding tags
- **Password Security**: Bcrypt hashing for secure password storage

## 🏗️ Architecture

This is a **3-layer Go REST API** with the following structure:

- **`main.go`**: Entry point - initializes database and starts Gin server on `:8080`
- **`db/`**: Database connection and table management
- **`models/`**: Data models with business logic methods
- **`routes/`**: HTTP handlers organized by domain (events, users)
- **`utils/`**: Shared utilities (password hashing)

## 📁 Project Structure

```
rest_api_golang/
├── main.go                 # Application entry point
├── api.db                  # SQLite database (auto-generated)
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── api-test/               # HTTP test files
│   ├── create-event.http
│   ├── create-user.http
│   ├── delete-event.http
│   ├── get-events.http
│   ├── get-single-event.http
│   └── update-event.http
├── db/
│   └── db.go              # Database connection & table creation
├── models/
│   ├── event.go           # Event model & methods
│   └── user.go            # User model & methods
├── routes/
│   ├── routes.go          # Route registration
│   ├── events.go          # Event-related handlers
│   └── users.go           # User-related handlers
└── utils/
    └── hash.go            # Password hashing utilities
```

## 🛠️ Technologies Used

- **Go 1.24.4**: Programming language
- **Gin Web Framework**: HTTP router and middleware
- **SQLite**: Embedded database
- **golang.org/x/crypto/bcrypt**: Password hashing
- **database/sql**: Go standard database interface

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

### Events

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/events` | Get all events |
| GET | `/events/:id` | Get a specific event |
| POST | `/events` | Create a new event |
| PUT | `/events/:id` | Update an event |
| DELETE | `/events/:id` | Delete an event |

### Users

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/signup` | Register a new user |
| POST | `/login` | User login |

## 🧪 Testing

Use the provided HTTP test files in the `api-test/` directory with VS Code REST Client extension or similar tools:

1. Create a user: `api-test/create-user.http`
2. Login: `api-test/login-user.http` (if implemented)
3. Create an event: `api-test/create-event.http`
4. Get events: `api-test/get-events.http`

## 🔧 Database Schema

### Users Table
- `id` (INTEGER, PRIMARY KEY)
- `email` (TEXT, UNIQUE, NOT NULL)
- `password` (TEXT, NOT NULL) - Bcrypt hashed

### Events Table
- `id` (INTEGER, PRIMARY KEY)
- `name` (TEXT, NOT NULL)
- `description` (TEXT, NOT NULL)
- `location` (TEXT, NOT NULL)
- `dateTime` (DATETIME, NOT NULL)
- `user_id` (INTEGER, FOREIGN KEY)

## 🔒 Security

- Passwords are hashed using bcrypt with cost factor 14
- Email uniqueness enforced at database level
- Input validation using Gin binding tags
- Prepared statements used to prevent SQL injection

## 🚧 Development Status

### Completed Features
- ✅ Event CRUD operations
- ✅ User registration with password hashing
- ✅ User login validation
- ✅ SQLite database integration
- ✅ Input validation

### Planned Features
- 🔲 JWT authentication middleware
- 🔲 Protected routes requiring authentication
- 🔲 Enhanced error handling
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

