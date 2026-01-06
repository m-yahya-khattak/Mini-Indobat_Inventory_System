# Setup Checklist

Use this checklist to set up the development environment for the Mini-Indobat Inventory System.

## Pre-Setup Requirements
- [ ] Review project requirements in [REQUIREMENTS.md](./REQUIREMENTS.md)
- [ ] Ensure you have administrator/sudo access if needed
- [ ] Check available disk space (recommended: at least 2GB free)

## Environment Setup

### Install Go
- [ ] Download and install Go (latest stable version)
  ```bash
  # Download from https://golang.org/dl/
  # Or using package manager (Windows: Chocolatey, Linux: apt/yum, Mac: Homebrew)
  ```
- [ ] Verify Go installation
  ```bash
  go version
  ```
- [ ] Set up GOPATH and GOROOT if needed
  ```bash
  go env GOPATH
  go env GOROOT
  ```

### Install Node.js and npm
- [ ] Download and install Node.js (LTS version recommended)
  ```bash
  # Download from https://nodejs.org/
  # Or using package manager
  ```
- [ ] Verify Node.js and npm installation
  ```bash
  node --version
  npm --version
  ```

### Install PostgreSQL
- [ ] Download and install PostgreSQL
  ```bash
  # Download from https://www.postgresql.org/download/
  # Or using package manager
  ```
- [ ] Verify PostgreSQL installation
  ```bash
  psql --version
  ```
- [ ] Start PostgreSQL service
  ```bash
  # Windows: Services → Start PostgreSQL
  # Linux: sudo systemctl start postgresql
  # Mac: brew services start postgresql
  ```

## Project Setup

### Navigate to Project Directory
- [ ] Navigate to project directory
  ```bash
  cd D:\codilityTest\Mini-Indobat_Inventory_System
  ```

### Database Setup
- [ ] Create PostgreSQL database
  ```bash
  # Option 1: Using createdb command
  createdb indobat_inventory
  
  # Option 2: Using psql
  psql -U postgres
  CREATE DATABASE indobat_inventory;
  \q
  ```
- [ ] Note database connection details (host, port, user, password, database name)

### Backend Setup
- [ ] Navigate to backend directory
  ```bash
  cd backend
  ```
- [ ] Initialize Go module
  ```bash
  go mod init mini-indobat-inventory
  ```
- [ ] Install Go dependencies
  ```bash
  go mod download
  ```
- [ ] Install database driver (e.g., pgx or lib/pq)
  ```bash
  go get github.com/lib/pq
  # or
  go get github.com/jackc/pgx/v5
  ```
- [ ] Install migration tool (Golang-Migrate or GORM)
  ```bash
  # Option 1: Golang-Migrate
  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  
  # Option 2: GORM (if using GORM)
  go get gorm.io/gorm
  go get gorm.io/driver/postgres
  ```
- [ ] Create .env file for backend configuration
  ```bash
  # Create .env file with:
  # DB_HOST=localhost
  # DB_PORT=5432
  # DB_USER=postgres
  # DB_PASSWORD=your_password
  # DB_NAME=indobat_inventory
  # SERVER_PORT=8080
  ```
- [ ] Run database migrations
  ```bash
  # If using golang-migrate
  migrate -path ./migrations -database "postgres://user:password@localhost/indobat_inventory?sslmode=disable" up
  
  # If using GORM AutoMigrate, it will run on application start
  ```

### Frontend Setup
- [ ] Navigate to frontend directory
  ```bash
  cd frontend
  ```
- [ ] Initialize Next.js project (if not already initialized)
  ```bash
  npx create-next-app@latest . --yes
  ```
- [ ] Install dependencies
  ```bash
  npm install
  # or
  yarn install
  ```
- [ ] Install additional dependencies (if needed)
  ```bash
  npm install axios
  npm install react-toastify
  # or other UI libraries as needed
  ```
- [ ] Create .env.local file for frontend configuration (if needed)
  ```bash
  # Create .env.local file with:
  # NEXT_PUBLIC_API_URL=http://localhost:8080
  ```

## Development Tools
- [ ] Install code editor/IDE (VS Code recommended)
- [ ] Install recommended VS Code extensions:
  - Go extension
  - ESLint
  - Prettier
  - PostgreSQL extension
- [ ] Set up version control (Git)
  ```bash
  git init
  ```
- [ ] Create .gitignore file
  ```bash
  # .gitignore should include:
  # node_modules/
  # .env
  # .env.local
  # *.log
  # dist/
  # build/
  # .DS_Store
  # etc.
  ```

## Verification
- [ ] Verify backend can connect to database
  ```bash
  cd backend
  go run main.go
  # Check if server starts without errors
  ```
- [ ] Verify frontend can start
  ```bash
  cd frontend
  npm run dev
  # Check if development server starts on http://localhost:3000
  ```
- [ ] Test API endpoints (using curl, Postman, or browser)
  ```bash
  # Test GET /products
  curl http://localhost:8080/products
  
  # Test POST /products
  curl -X POST http://localhost:8080/products \
    -H "Content-Type: application/json" \
    -d '{"name":"Test Drug","stock":100,"price":5000}'
  ```
- [ ] Verify all dependencies are installed correctly
- [ ] Test basic project functionality

## Notes
- All terminal commands should be run manually by the developer
- Make sure PostgreSQL service is running before starting the backend
- Update .env files with your actual database credentials
- Keep .env files in .gitignore and never commit them

