# Backend Setup Guide

Follow these steps to set up the Go backend.

## Step 1: Navigate to Backend Directory
```powershell
cd backend
```

## Step 2: Initialize/Update Go Module
The `go.mod` file already exists, but you can update it:
```powershell
go mod tidy
```

## Step 3: Install Required Dependencies

### Install HTTP Router (Choose one):
**Option A: Gin (Recommended - Popular and easy)**
```powershell
go get github.com/gin-gonic/gin
```

**Option B: Echo**
```powershell
go get github.com/labstack/echo/v4
```

**Option C: Standard library (net/http)**
```powershell
# No installation needed, built-in
```

### Install Database Driver (Choose one):
**Option A: pgx (Recommended - Modern, better performance)**
```powershell
go get github.com/jackc/pgx/v5
go get github.com/jackc/pgx/v5/pgxpool
```

**Option B: lib/pq (Traditional)**
```powershell
go get github.com/lib/pq
```

### Install Environment Variable Loader:
```powershell
go get github.com/joho/godotenv
```

### Install Migration Tool (Choose one):
**Option A: golang-migrate (Recommended)**
```powershell
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

**Option B: GORM (if you prefer ORM)**
```powershell
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

### Install CORS Middleware (if using Gin):
```powershell
go get github.com/gin-contrib/cors
```

## Step 4: Create .env File
Create a `.env` file in the `backend/` directory with the following content:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password_here
DB_NAME=indobat_inventory
DB_SSLMODE=disable

# Server Configuration
SERVER_PORT=8080
SERVER_HOST=localhost

# Environment
ENV=development
```

**Important**: Replace `your_password_here` with your actual PostgreSQL password that you set during installation.

## Step 5: Verify Dependencies
```powershell
go mod download
go mod tidy
```

## Step 6: Verify Setup
Check that all dependencies are installed:
```powershell
go list -m all
```

## Next Steps
1. Create database (if not done): See main SETUP_CHECKLIST.md
2. Create database migrations
3. Set up database connection in code
4. Implement handlers, services, and repositories

## Recommended Tech Stack Choices
For this project, I recommend:
- **Router**: Gin (easy to use, good documentation)
- **Database Driver**: pgx/v5 (modern, performant, good for transactions)
- **Migration Tool**: golang-migrate (explicit migrations, better for production)

