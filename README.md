# Mini-Indobat Inventory System

A stock management system for pharmaceutical wholesalers (PBF Farmasi) that handles product inventory and purchase transactions with race condition protection for concurrent requests.

## Table of Contents

- [Technology Stack](#technology-stack)
- [Features](#features)
- [Architecture](#architecture)
- [Prerequisites](#prerequisites)
- [Installation & Setup](#installation--setup)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [Testing](#testing)
- [Race Condition Protection](#race-condition-protection)
- [Project Structure](#project-structure)
- [Key Implementation Details](#key-implementation-details)
- [Postman Collection](#postman-collection)

## Technology Stack

**Backend:**
- Go 1.23+
- Gin (HTTP web framework)
- pgx/v5 (PostgreSQL driver)
- golang-migrate (database migrations)

**Frontend:**
- Next.js 15+
- React 19+
- TypeScript
- Tailwind CSS
- Axios (HTTP client)
- React Toastify (notifications)

**Database:**
- PostgreSQL 16+

**Development Tools:**
- Postman (API testing)

## Features

**Product Management:**
- List all products with current stock levels
- Create new products

**Order Management:**
- Create purchase orders with quantity and discount
- Automatic stock reduction
- Real-time price calculation
- Transaction history recording

**Race Condition Protection:**
- Database-level locking prevents concurrent order conflicts
- Ensures stock integrity under high concurrency
- Only one successful order when multiple requests compete for limited stock

**Frontend:**
- Real-time product inventory dashboard
- Order form with live price estimation
- Error handling with user-friendly notifications
- Manual refresh capability

## Architecture

The application follows Clean Architecture principles with clear separation of concerns:

- **Handler Layer**: HTTP request/response handling, input validation
- **Service Layer**: Business logic, transaction coordination, error handling
- **Repository Layer**: Database operations, SQL queries, data access

This layered approach ensures:
- Business logic is isolated from HTTP concerns
- Database operations are abstracted
- Code is testable and maintainable
- Changes in one layer don't affect others

## Prerequisites

Before setting up the application, ensure you have the following installed:

- **Go**: Version 1.23 or higher ([Download](https://golang.org/dl/))
- **Node.js**: Version 18 or higher ([Download](https://nodejs.org/))
- **PostgreSQL**: Version 16 or higher ([Download](https://www.postgresql.org/download/))
- **Git**: For cloning the repository

Verify installations:

```bash
go version
node --version
npm --version
psql --version
```

## Installation & Setup

### Database Setup

1. Create PostgreSQL database:

```bash
psql -U postgres
CREATE DATABASE indobat_inventory;
\q
```

2. Navigate to backend directory and run migrations:

```bash
cd backend
migrate -path ./migrations -database "postgres://postgres:YOUR_PASSWORD@localhost:5432/indobat_inventory?sslmode=disable" up
```

Replace `YOUR_PASSWORD` with your PostgreSQL password.

### Backend Setup

1. Navigate to backend directory:

```bash
cd backend
```

2. Install dependencies:

```bash
go mod download
```

3. Create `.env` file in the `backend` directory:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_postgres_password
DB_NAME=indobat_inventory
DB_SSLMODE=disable
SERVER_PORT=8080
ENV=development
```

4. Update `DB_PASSWORD` with your actual PostgreSQL password.

### Frontend Setup

1. Navigate to frontend directory:

```bash
cd frontend
```

2. Install dependencies:

```bash
npm install
```

3. Create `.env.local` file in the `frontend` directory (optional, defaults to `http://localhost:8080`):

```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

## Running the Application

### Start Backend Server

From the `backend` directory:

```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`. You should see:

```
Mini-Indobat Inventory System - Backend Server
Server starting...
Database connection established successfully
Server listening on port 8080
```

### Start Frontend Application

From the `frontend` directory (in a new terminal):

```bash
npm run dev
```

The frontend will start on `http://localhost:3000`. Open this URL in your browser to access the application.

### Verify Setup

Test the backend health endpoint:

```bash
curl http://localhost:8080/health
```

Expected response:

```json
{"status":"ok"}
```

## API Documentation

The API provides three main endpoints as specified in the requirements.

### GET /products

Retrieves a list of all products.

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "name": "Paracetamol 500mg",
    "stock": 100,
    "price": 5000,
    "created_at": "2026-01-06T12:00:00Z",
    "updated_at": "2026-01-06T12:00:00Z"
  }
]
```

### POST /products

Creates a new product.

**Request Body:**
```json
{
  "name": "Paracetamol 500mg",
  "stock": 100,
  "price": 5000
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "name": "Paracetamol 500mg",
  "stock": 100,
  "price": 5000,
  "created_at": "2026-01-06T12:00:00Z",
  "updated_at": "2026-01-06T12:00:00Z"
}
```

**Error Response (400 Bad Request):**
```json
{
  "error": "Invalid request body",
  "details": "validation error details"
}
```

### POST /order

Creates a purchase order and reduces product stock.

**Request Body:**
```json
{
  "product_id": 1,
  "quantity": 5,
  "discount_percent": 10
}
```

**Response (200 OK):**
```json
{
  "order_id": 123,
  "product_id": 1,
  "quantity": 5,
  "total_price": 22500,
  "message": "Order successful"
}
```

**Error Responses:**

Insufficient Stock (400 Bad Request):
```json
{
  "error": "insufficient stock: available stock is 2, requested 5"
}
```

Product Not Found (404 Not Found):
```json
{
  "error": "product not found"
}
```

Invalid Input (400 Bad Request):
```json
{
  "error": "invalid quantity"
}
```

**Price Calculation Formula:**
```
Total Price = (Product Price × Quantity) × (1 - Discount Percent / 100)
```

For complete API documentation with request/response examples and testing instructions, see the [Postman Collection](#postman-collection) section.

## Testing

### Unit Tests

The backend includes unit tests for core business logic:

**Run all tests:**
```bash
cd backend
go test ./internal/service/... -v -cover
```

**Test Coverage:**
- Price calculation logic (various discount scenarios)
- Stock validation logic (sufficient/insufficient stock)
- Input validation (quantity and discount ranges)
- Edge cases (zero stock, negative values, boundary conditions)

**Expected Output:**
```
=== RUN   TestOrderService_InputValidation
=== RUN   TestCalculateOrderPrice
=== RUN   TestStockValidation
PASS
coverage: 14.3% of statements
```

### Race Condition Testing

Race condition protection is tested using the Postman Collection Runner. See the [Race Condition Protection](#race-condition-protection) section for detailed testing instructions.

### Manual Testing

1. Start backend and frontend servers
2. Access frontend at `http://localhost:3000`
3. Create products using the form or API
4. Place orders and verify stock reduction
5. Test concurrent requests using Postman Collection Runner

## Race Condition Protection

The system implements database-level race condition protection to ensure data integrity when multiple concurrent requests attempt to purchase the same product with limited stock.

### Implementation

**Database Transaction Isolation:**
- Uses PostgreSQL `SERIALIZABLE` isolation level
- Ensures strict serialization of concurrent transactions
- Prevents phantom reads and non-repeatable reads

**Row-Level Locking:**
- Uses `SELECT FOR UPDATE` to lock product rows during order processing
- Prevents other transactions from reading/modifying the same product until the transaction completes
- Ensures only one transaction can process an order for a specific product at a time

**Transaction Flow:**
1. Begin transaction with SERIALIZABLE isolation
2. Lock product row using SELECT FOR UPDATE
3. Validate stock availability
4. Calculate total price
5. Create order record
6. Update product stock
7. Commit transaction (or rollback on error)

### Testing Race Condition Protection

**Scenario:** Product has 1 unit in stock, 10 concurrent requests attempt to purchase it.

**Expected Result:** Only 1 request succeeds, 9 requests fail with "insufficient stock" error.

**Test Steps:**

1. Create a product with stock = 1:

```bash
POST /products
{
  "name": "Test Product",
  "stock": 1,
  "price": 10000
}
```

2. Import the Postman Collection (`backend/postman_collection.json`)

3. Use Postman Collection Runner:
   - Select the "Race Condition Test" request
   - Set iterations to 10
   - Set delay to 0ms (concurrent execution)
   - Run the collection

4. Verify results:
   - Check that only 1 request returned status 200
   - Check that 9 requests returned status 400 with "insufficient stock" error
   - Verify product stock is now 0 (not negative)

**Alternative Test Using curl:**

```bash
# Run 10 concurrent requests
for i in {1..10}; do
  curl -X POST http://localhost:8080/order \
    -H "Content-Type: application/json" \
    -d '{"product_id":1,"quantity":1,"discount_percent":0}' &
done
wait
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
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── config/
│   │   │   ├── config.go
│   │   │   └── database.go
│   │   ├── handler/
│   │   │   ├── product_handler.go
│   │   │   ├── order_handler.go
│   │   │   └── router.go
│   │   ├── service/
│   │   │   ├── product_service.go
│   │   │   ├── order_service.go
│   │   │   ├── price_calculator.go
│   │   │   └── errors.go
│   │   ├── repository/
│   │   │   ├── product_repository.go
│   │   │   └── order_repository.go
│   │   └── model/
│   │       ├── product.go
│   │       └── order.go
│   ├── migrations/
│   │   ├── 000001_create_products_table.up.sql
│   │   ├── 000001_create_products_table.down.sql
│   │   ├── 000002_create_orders_table.up.sql
│   │   ├── 000002_create_orders_table.down.sql
│   │   ├── 000003_add_updated_at_trigger.up.sql
│   │   └── 000003_add_updated_at_trigger.down.sql
│   ├── postman_collection.json
│   ├── go.mod
│   └── go.sum
└── frontend/
    ├── src/
    │   ├── app/
    │   │   ├── layout.tsx
    │   │   ├── page.tsx
    │   │   └── globals.css
    │   ├── components/
    │   │   ├── ProductTable.tsx
    │   │   ├── OrderForm.tsx
    │   │   └── ToastProvider.tsx
    │   └── lib/
    │       ├── api.ts
    │       └── utils.ts
    ├── package.json
    └── tsconfig.json
```

## Key Implementation Details

### Data Integrity

**Stock Validation:**
- Stock is validated before order creation
- Database constraints prevent negative stock values
- Transaction rollback ensures atomicity if stock update fails

**Transaction Management:**
- All order operations (stock check, order creation, stock update) occur within a single database transaction
- Transaction rollback on any error ensures data consistency
- SERIALIZABLE isolation level prevents race conditions

### Security

**SQL Injection Prevention:**
- All database queries use parameterized statements
- No string concatenation in SQL queries
- Input validation at handler and service layers

**Input Validation:**
- Request body validation using Gin binding
- Business logic validation in service layer
- Quantity must be positive integer
- Discount percentage must be between 0 and 100

### Error Handling

**Structured Error Responses:**
- Consistent error format across all endpoints
- Specific error messages for different failure scenarios
- HTTP status codes follow REST conventions

**Service Layer Errors:**
- Custom error types for different failure scenarios
- Error wrapping for context preservation
- Proper error propagation from repository to handler

### Code Organization

**Clean Architecture:**
- Handler layer: HTTP concerns only
- Service layer: Business logic and validation
- Repository layer: Database operations
- Model layer: Data structures and DTOs

**Separation of Concerns:**
- Business logic never in handlers
- Database queries isolated in repositories
- Reusable service functions (price calculation, stock validation)

## Postman Collection

A Postman collection is included for API testing and race condition verification.

**Location:** `backend/postman_collection.json`

### Importing the Collection

1. Open Postman
2. Click "Import" button
3. Select "File" tab
4. Choose `backend/postman_collection.json`
5. Click "Import"

### Collection Contents

The collection includes:

1. **Health Check** - Verify server is running
2. **Get All Products** - List all products
3. **Create Product** - Add a new product
4. **Create Order** - Place an order
5. **Race Condition Test** - Test concurrent order requests

### Running Tests

**Individual Request:**
- Select a request from the collection
- Click "Send"
- Review response

**Collection Runner (Race Condition Test):**
1. Click "Collection Runner" icon
2. Select "Race Condition Test" request
3. Set "Iterations" to 10
4. Set "Delay" to 0ms
5. Click "Run Race Condition Test"
6. Review results - only 1 should succeed, 9 should fail

**Full Collection:**
1. Click "..." next to collection name
2. Select "Run collection"
3. Review all test results

### Environment Variables (Optional)

Create a Postman environment with:
- `base_url`: `http://localhost:8080`
- Update collection requests to use `{{base_url}}` variable
