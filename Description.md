# REST API in Go - Project Description

## 📋 Project Overview
This is a **RESTful API built with Go** following a course by Maximilian Schwarzmüller on Udemy. It's an event management system with full CRUD operations for managing events with persistent data storage.

## 🏗️ Architecture & Technologies

### Framework & Database
- **Gin Web Framework** - A lightweight, fast HTTP web framework for Go
- **SQLite3** - Local database for data persistence using `github.com/mattn/go-sqlite3`
- **Go 1.24.4** - Latest Go version

### Project Structure
```
rest_api_golang/
├── main.go                 # Entry point & server setup
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── README.md               # Basic project info
├── Description.md          # This file - detailed documentation
├── db/
│   └── db.go              # Database configuration & connection
├── models/
│   └── event.go           # Event data model & business logic
├── routes/
│   ├── routes.go          # Route registration
│   └── events.go          # Event route handlers
├── api-test/
│   ├── create-event.http  # POST request test
│   └── get-events.http    # GET request test
└── Resources/             # Learning materials & documentation links
```

## 🚀 API Endpoints

The API provides complete CRUD operations for events:

| Method | Endpoint | Description | Example |
|--------|----------|-------------|---------|
| `GET` | `/events` | Get all events | Returns JSON array of all events |
| `GET` | `/events/:id` | Get a specific event | `/events/1` returns event with ID 1 |
| `POST` | `/events` | Create a new event | Requires JSON body with event data |
| `PUT` | `/events/:id` | Update an existing event | Updates event with specified ID |
| `DELETE` | `/events/:id` | Delete an event | Removes event with specified ID |

## 📊 Data Model

### Event Structure
The `Event` struct includes the following fields:

```go
type Event struct {
    ID          int64     // Auto-incrementing primary key
    Name        string    // Required event name
    Description string    // Required event description
    Location    string    // Required event location
    DateTime    time.Time // Required event date/time
    UserID      int       // Associated user ID
}
```

### Database Schema
```sql
CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    location TEXT NOT NULL,
    dateTime DATETIME NOT NULL,
    user_id INTEGER
);
```

## 🔧 Key Features

1. **Database Integration**
   - SQLite3 with prepared statements for security
   - Automatic table creation on startup
   - Connection pooling (max 10 open, 5 idle connections)

2. **JSON Handling**
   - Automatic request/response JSON binding
   - Struct tag validation for required fields
   - Clean error responses

3. **Error Handling**
   - Comprehensive error responses with appropriate HTTP status codes
   - Database error handling with meaningful messages
   - Input validation error handling

4. **HTTP Testing**
   - Ready-made `.http` files for API testing
   - Example requests for all endpoints

5. **Security**
   - SQL injection protection via prepared statements
   - Input validation using Gin's binding features

## 🔍 Current Implementation Details

### Server Configuration
- **Port**: 8080
- **Framework**: Gin with default middleware (logger and recovery)
- **Database**: SQLite file (`api.db`) created automatically

### Sample API Usage

#### Create Event
```http
POST http://localhost:8080/events
Content-Type: application/json

{
  "name": "Test event",
  "description": "A test event",
  "location": "A test location",
  "dateTime": "2025-01-01T15:30:00.000Z"
}
```

#### Get All Events
```http
GET http://localhost:8080/events
```

### Response Examples

**Success Response (Create Event):**
```json
{
  "message": "Event created!",
  "event": {
    "ID": 1,
    "Name": "Test event",
    "Description": "A test event",
    "Location": "A test location",
    "DateTime": "2025-01-01T15:30:00Z",
    "UserID": 1
  }
}
```

**Error Response:**
```json
{
  "message": "Could not parse request data."
}
```

## 🎯 Technical Highlights

### Database Operations
- **Create**: Insert new events with auto-generated IDs
- **Read**: Query all events or specific events by ID
- **Update**: Modify existing events while preserving ID
- **Delete**: Remove events by ID

### Validation
- Required field validation using struct tags (`binding:"required"`)
- Type validation for date/time fields
- ID parsing validation for route parameters

### Error Handling Patterns
- Database connection errors with panic for critical failures
- Request parsing errors with 400 Bad Request
- Not found errors with 500 Internal Server Error
- Successful operations with appropriate 2xx status codes

## 🚧 Current Limitations & Future Improvements

### Current State
- User authentication is not implemented (UserID hardcoded to 1)
- No pagination for large event lists
- Basic error messages without detailed validation feedback
- No API documentation (Swagger/OpenAPI)

### Potential Enhancements
1. **Authentication & Authorization**
   - JWT token-based authentication
   - User registration and login endpoints
   - Protected routes with middleware

2. **API Improvements**
   - Pagination and filtering for event lists
   - Search functionality
   - Sorting options

3. **Data Validation**
   - More sophisticated input validation
   - Custom validation rules
   - Better error messages

4. **Testing & Documentation**
   - Unit tests for all handlers and models
   - Integration tests
   - Swagger/OpenAPI documentation
   - Postman collection

5. **Configuration**
   - Environment-based configuration
   - Database connection string from environment
   - Configurable server port

6. **Performance**
   - Database indexing
   - Caching layer
   - Connection pooling optimization

## 🔗 Learning Resources

The project includes links to various learning resources in the `Resources/` folder covering:
- HTTP fundamentals
- Web development concepts
- Go-specific libraries (Gin, SQLite driver)
- Internet and web protocols

## 🏃 Getting Started

1. **Prerequisites**: Go 1.24.4 or later
2. **Run the application**: `go run main.go`
3. **Test the API**: Use the provided `.http` files in `api-test/` folder
4. **Database**: SQLite file `api.db` will be created automatically

The server will start on `http://localhost:8080` and create the necessary database tables on first run.
