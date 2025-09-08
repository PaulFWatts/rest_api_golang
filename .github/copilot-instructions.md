# GitHub Copilot Instructions for REST API in Go

## Architecture Overview

This is a **3-layer Go REST API** with event management functionality:

- **`main.go`**: Entry point - initializes database and starts Gin server on `:8080`
- **`db/`**: Global SQLite connection (`db.DB`) with auto table creation
- **`models/`**: Data models with pointer receiver methods (e.g., `func (e *Event) Save()`)
- **`routes/`**: Gin handlers split by domain (`events.go`, `users.go`) + registration (`routes.go`)
- **`middlewares/`**: Authentication middleware for JWT token validation
- **`utils/`**: Shared utilities (bcrypt password hashing, JWT token management)

## Key Patterns & Conventions

### Authentication Architecture
- **JWT-based authentication**: Users login to receive JWT tokens
- **Middleware pattern**: `middlewares.Authenticate` validates tokens and sets `userId` in context
- **Protected routes**: Use middleware like `server.POST("/events", middlewares.Authenticate, createEvent)`
- **Context extraction**: Handlers get user ID with `context.GetInt64("userId")`

### Database Operations
- Use **pointer receivers** for model methods that modify structs: `func (e *Event) Save() error`
- **Prepared statements** pattern: `db.DB.Prepare()` → `stmt.Exec()` → `defer stmt.Close()`
- **Auto-ID assignment**: Always set `struct.ID = result.LastInsertId()` in Save methods
- Tables auto-created in `db.createTables()` with foreign key relationships

### HTTP Handlers
- **Gin Context pattern**: `func handlerName(context *gin.Context)`
- **JSON binding**: `context.ShouldBindJSON(&struct)` for request parsing
- **Error responses**: `gin.H{"message": "...", "error": err.Error()}` for debugging
- **ID parsing**: `strconv.ParseInt(context.Param("id"), 10, 64)` for URL params

## Development Workflows

### Running & Testing
```bash
# Start server (creates api.db automatically)
go run main.go

# Test with HTTP files in api-test/
# Use VS Code REST Client or similar tools
```

### Authentication Flow
1. **Signup**: `POST /signup` → create user with hashed password
2. **Login**: `POST /login` → validate credentials, return JWT token
3. **Protected requests**: Include `Authorization: <token>` header
4. **Middleware**: Validates token, sets `userId` in context for handlers

### Database Schema
- **SQLite file**: `api.db` (ignored in git)
- **Users**: `id`, `email` (unique), `password` (hashed)
- **Events**: `id`, `name`, `description`, `location`, `dateTime`, `user_id` (FK)

### Adding New Features
1. **Models**: Add struct + methods in `models/` (use pointer receivers for mutations)
2. **Routes**: Add handler in `routes/domain.go`, register in `routes/routes.go`
3. **Authentication**: Apply `middlewares.Authenticate` to protected routes
4. **Testing**: Create `.http` file in `api-test/`

## Critical Implementation Notes

- **Module path**: `github.com/PaulFWatts/rest_api_golang` (check imports match)
- **Error handling**: Always check `err != nil`, return HTTP 400/500 with descriptive messages
- **JSON comments**: Never include `//` comments in JSON test files - causes parsing errors
- **Resource cleanup**: Use `defer stmt.Close()` and `defer rows.Close()` consistently
- **JWT tokens**: 2-hour expiration, HMAC-SHA256 signing, include `email` and `userId` claims

## Current Implementation Status
- ✅ **JWT Authentication**: Complete with middleware-based token validation
- ✅ **User registration/login**: Working with bcrypt password hashing
- ✅ **Protected routes**: Event creation requires authentication
- ✅ **User ownership**: Events automatically assigned to authenticated user
- ⚠️ **Authorization**: Update/Delete events don't check ownership yet
- ⚠️ **Environment config**: Secret key and database path still hardcoded

## Quick Reference
- **Server port**: `:8080`
- **Test files**: `api-test/*.http`
- **Database**: SQLite (`api.db`)
- **Framework**: Gin web framework
- **Authentication**: JWT middleware with 2-hour expiration
- **Go version**: 1.24.4
