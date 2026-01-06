# Mini-Indobat Inventory System

## Project Overview
A simple stock management system for a Pharmaceutical Wholesaler (PBF Farmasi). The system handles incoming and outgoing goods accurately, calculates total prices correctly, and handles concurrent requests safely (Concurrency Safe).

## Technology Stack
- **Backend**: Go (framework: free choice)
- **Frontend**: Next.js
- **Database**: PostgreSQL

## Features

### Backend - REST API
- **Product Management**
  - `GET /products` - List all products (ID, Drug Name, Stock, Price)
  - `POST /products` - Add new product

- **Transaction Management**
  - `POST /order` - Perform purchase (reduce stock)
    - Input: `product_id`, `quantity`, `discount_percent`
    - Features:
      - Stock validation (prevents negative stock)
      - Price calculation: (Price × Quantity) - (Discount %)
      - Transaction history recording
      - **Race condition protection** using database locking

### Frontend - Single Page Application
- **Dashboard**: Real-time (or manual refresh) product list with stock levels
- **Order Form**:
  - Drug selection dropdown
  - Quantity input
  - Discount percentage input
  - Real-time price estimation before submission
- **Error Handling**: Toast/Alert notifications for errors (stock out, validation failures)

## Prerequisites
- Go (latest stable version)
- Node.js and npm/yarn
- PostgreSQL
- Git

## Installation & Setup
See [SETUP_CHECKLIST.md](./SETUP_CHECKLIST.md) for detailed setup instructions.

### Quick Start

#### Database Setup
```bash
# Create PostgreSQL database
createdb indobat_inventory

# Or using psql
psql -U postgres
CREATE DATABASE indobat_inventory;
```

#### Backend Setup
```bash
cd backend
go mod download
# Configure .env file with database connection
go run main.go
```

#### Frontend Setup
```bash
cd frontend
npm install
# Configure .env file if needed
npm run dev
```

## API Documentation

### Endpoints

#### GET /products
Get list of all products.

**Response:**
```json
[
  {
    "id": 1,
    "name": "Paracetamol 500mg",
    "stock": 100,
    "price": 5000
  }
]
```

#### POST /products
Create a new product.

**Request Body:**
```json
{
  "name": "Paracetamol 500mg",
  "stock": 100,
  "price": 5000
}
```

#### POST /order
Create a purchase order.

**Request Body:**
```json
{
  "product_id": 1,
  "quantity": 5,
  "discount_percent": 10
}
```

**Response:**
```json
{
  "order_id": 123,
  "product_id": 1,
  "quantity": 5,
  "total_price": 22500,
  "message": "Order successful"
}
```

**Error Response:**
```json
{
  "error": "Insufficient stock",
  "available_stock": 2
}
```

## Project Structure
```
Mini-Indobat_Inventory_System/
├── README.md
├── REQUIREMENTS.md
├── SETUP_CHECKLIST.md
├── IMPLEMENTATION_CHECKLIST.md
├── backend/
│   ├── cmd/
│   ├── internal/
│   │   ├── handler/
│   │   ├── service/
│   │   ├── repository/
│   │   └── model/
│   ├── migrations/
│   └── go.mod
└── frontend/
    ├── app/
    ├── components/
    └── package.json
```

## Development
See [IMPLEMENTATION_CHECKLIST.md](./IMPLEMENTATION_CHECKLIST.md) for implementation tasks.

## Key Implementation Notes

### Priority Order (for implementation):
1. **Security** - SQL injection prevention, input validation, race condition handling
2. **Best Practices** - Clean architecture, proper error handling, code organization
3. **Scalability/Generic** - Reusable code, proper abstractions
4. **Structure** - Layered architecture (Handler → Service → Repository)
5. **Sticking to Instructions** - Follow requirements precisely, don't overdo it

### Critical Requirements:
- ✅ Stock must never go negative (data integrity)
- ✅ Race condition protection using database transactions and locking
- ✅ Clean architecture with separated layers
- ✅ Database migrations (Golang-Migrate or GORM AutoMigrate)
- ✅ Frontend with real-time validation and loading states
- ✅ Proper error handling and user feedback

## Testing
- Unit tests for price calculation and stock logic (Nice to Have)
- Integration tests for API endpoints
- Manual testing for concurrent requests

## License
[To be determined]

