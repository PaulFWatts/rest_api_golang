# GitHub Copilot Instructions for REST API in Go

## Architecture Overview

This is a **complete 4-layer Go REST API** with event management and user registration functionality:

- **`main.go`**: Entry point - initializes database and starts Gin server on `:8080`
- **`db/`**: Global SQLite connection (`db.DB`) with auto table creation (users, events, registrations)
- **`models/`**: Data models with pointer receiver methods (e.g., `func (e *Event) Save()`)
- **`routes/`**: Gin handlers split by domain (`events.go`, `users.go`, `register.go`) + registration (`routes.go`)
- **`middlewares/`**: Authentication middleware for JWT token validation
- **`utils/`**: Shared utilities (bcrypt password hashing, JWT token management)

## Key Patterns & Conventions

### Authentication Architecture
- **JWT-based authentication**: Users login to receive JWT tokens with 2-hour expiration
- **Middleware pattern**: `middlewares.Authenticate` validates tokens and sets `userId` in context
- **Route groups**: Protected routes use `authenticated.Use(middlewares.Authenticate)`
- **Context extraction**: Handlers get user ID with `context.GetInt64("userId")`

### Authorization & Ownership
- **Event ownership**: Users can only update/delete their own events (validated via `event.UserID != userID`)
- **Registration system**: Users can register/unregister for events they don't own
- **Proper error messages**: Generic "Not authorized" vs specific ownership errors

### Database Operations
- Use **pointer receivers** for model methods that modify structs: `func (e *Event) Save() error`
- **Prepared statements** pattern: `db.DB.Prepare()` → `stmt.Exec()` → `defer stmt.Close()`
- **Auto-ID assignment**: Always set `struct.ID = result.LastInsertId()` in Save methods
- **Three-table schema**: users, events, registrations with proper foreign keys

### HTTP Handlers
- **Gin Context pattern**: `func handlerName(context *gin.Context)`
- **JSON binding**: `context.ShouldBindJSON(&struct)` for request parsing
- **Error responses**: `gin.H{"message": "...", "error": err.Error()}` for debugging
- **ID parsing**: `strconv.ParseInt(context.Param("id"), 10, 64)` for URL params

## Development Workflows

### Running & Testing
```bash
# Start server (creates api.db automatically)
go run .

# Test with HTTP files in api-test/
# Use VS Code REST Client or similar tools
```

### Complete Authentication Flow
1. **Signup**: `POST /signup` → create user with bcrypt-hashed password
2. **Login**: `POST /login` → validate credentials, return JWT token
3. **Protected requests**: Include `Authorization: <token>` header
4. **Middleware**: Validates token, sets `userId` in context for handlers

### Database Schema
- **SQLite file**: `api.db` (ignored in git)
- **Users**: `id`, `email` (unique), `password` (bcrypt hashed)
- **Events**: `id`, `name`, `description`, `location`, `dateTime`, `user_id` (FK)
- **Registrations**: `id`, `event_id` (FK), `user_id` (FK) - many-to-many relationship

### Adding New Features
1. **Models**: Add struct + methods in `models/` (use pointer receivers for mutations)
2. **Routes**: Add handler in `routes/domain.go`, register in `routes/routes.go`
3. **Authentication**: Apply to route groups or individual routes
4. **Testing**: Create `.http` file in `api-test/` with real JWT tokens

## API Endpoints

### Public Endpoints
- `GET /events` - List all events
- `GET /events/:id` - Get specific event
- `POST /signup` - User registration
- `POST /login` - User authentication (returns JWT)

### Protected Endpoints (require JWT)
- `POST /events` - Create event (auto-assigns to authenticated user)
- `PUT /events/:id` - Update event (owner only)
- `DELETE /events/:id` - Delete event (owner only)
- `POST /events/:id/register` - Register for event
- `DELETE /events/:id/register` - Cancel event registration

## Critical Implementation Notes

- **Module path**: `github.com/PaulFWatts/rest_api_golang` (check imports match)
- **Error handling**: Always check `err != nil`, return HTTP 400/500 with descriptive messages
- **JSON comments**: Never include `//` comments in JSON test files - causes parsing errors
- **Resource cleanup**: Use `defer stmt.Close()` and `defer rows.Close()` consistently
- **JWT tokens**: 2-hour expiration, HMAC-SHA256 signing, include `email` and `userId` claims
- **Password security**: Bcrypt cost factor 14 for production-grade hashing

## Current Implementation Status
- ✅ **JWT Authentication**: Complete with middleware-based token validation
- ✅ **User registration/login**: Working with bcrypt password hashing and JWT generation
- ✅ **Protected routes**: Event CRUD requires authentication
- ✅ **User ownership**: Events automatically assigned to authenticated user
- ✅ **Authorization**: Update/Delete events check ownership
- ✅ **Event registration**: Users can register/unregister for events
- ✅ **Complete CRUD**: All operations fully implemented with proper error handling
- ⚠️ **Environment config**: Secret key and database path still hardcoded

## Quick Reference
- **Server port**: `:8080`
- **Test files**: `api-test/*.http` (includes real JWT tokens for testing)
- **Database**: SQLite (`api.db`) with 3 tables
- **Framework**: Gin web framework with route groups
- **Authentication**: JWT middleware with 2-hour expiration
- **Go version**: 1.24.4
- **Dependencies**: Gin, JWT, SQLite, bcrypt
