# GitHub Copilot Instructions for REST API in Go

## Architecture Overview

This is a **3-layer Go REST API** with event management functionality:

- **`main.go`**: Entry point - initializes database and starts Gin server on `:8080`
- **`db/`**: Global SQLite connection (`db.DB`) with auto table creation
- **`models/`**: Data models with pointer receiver methods (e.g., `func (e *Event) Save()`)
- **`routes/`**: Gin handlers split by domain (`events.go`, `users.go`) + registration (`routes.go`)
- **`utils/`**: Shared utilities (currently bcrypt password hashing)

## Key Patterns & Conventions

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

### Authentication (In Progress)
- **Password hashing**: Use `utils.HashPassword()` before saving users
- **Validation**: `user.ValidateCredentials()` method compares hashed passwords
- **Routes**: `/signup` and `/login` endpoints exist but JWT implementation pending

## Development Workflows

### Running & Testing
```bash
# Start server (creates api.db automatically)
go run main.go

# Test with HTTP files in api-test/
# Use VS Code REST Client or similar tools
```

### Database Schema
- **SQLite file**: `api.db` (ignored in git)
- **Users**: `id`, `email` (unique), `password` (hashed)
- **Events**: `id`, `name`, `description`, `location`, `dateTime`, `user_id` (FK)

### Adding New Features
1. **Models**: Add struct + methods in `models/` (use pointer receivers for mutations)
2. **Routes**: Add handler in `routes/domain.go`, register in `routes/routes.go`
3. **Testing**: Create `.http` file in `api-test/`

## Critical Implementation Notes

- **Module path**: `github.com/PaulFWatts/rest_api_golang` (check imports match)
- **Error handling**: Always check `err != nil`, return HTTP 400/500 with descriptive messages
- **JSON comments**: Never include `//` comments in JSON test files - causes parsing errors
- **Resource cleanup**: Use `defer stmt.Close()` and `defer rows.Close()` consistently

## Current Limitations
- UserID hardcoded to `1` in event creation (no auth middleware yet)
- No input validation beyond Gin's `binding:"required"` tags  
- No pagination or filtering on GET endpoints
- JWT authentication system incomplete (login endpoint exists but no token generation)

## Quick Reference
- **Server port**: `:8080`
- **Test files**: `api-test/*.http`
- **Database**: SQLite (`api.db`)
- **Framework**: Gin web framework
- **Go version**: 1.24.4
