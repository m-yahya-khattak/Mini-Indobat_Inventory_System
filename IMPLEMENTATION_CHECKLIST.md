# Implementation Checklist

Use this checklist to track the implementation progress of the Mini-Indobat Inventory System.

**Priority Order for Implementation:**
1. Security (SQL injection prevention, input validation, race condition handling)
2. Best Practices (Clean architecture, proper error handling, code organization)
3. Scalability/Generic (Reusable code, proper abstractions)
4. Structure (Layered architecture: Handler → Service → Repository)
5. Sticking to Instructions (Follow requirements precisely, don't overdo it)

## Phase 1: Project Initialization
- [ ] Create project structure (backend/ and frontend/ directories)
- [ ] Initialize Go module in backend
- [ ] Initialize Next.js project in frontend
- [ ] Set up .gitignore files
- [ ] Initialize Git repository
- [ ] Create .env.example files for configuration reference

## Phase 2: Database Setup & Infrastructure

### Database Schema Design
- [ ] Design Products table schema
  - [ ] id (primary key)
  - [ ] name (drug name)
  - [ ] stock (integer, not null, default 0)
  - [ ] price (decimal/numeric)
  - [ ] created_at, updated_at timestamps
- [ ] Design Transactions/Orders table schema
  - [ ] id (primary key)
  - [ ] product_id (foreign key)
  - [ ] quantity (integer)
  - [ ] discount_percent (decimal)
  - [ ] total_price (calculated)
  - [ ] created_at timestamp
- [ ] Create database migration files
  - [ ] Use golang-migrate or GORM AutoMigrate
  - [ ] Ensure migrations are versioned and reversible

### Backend Infrastructure
- [ ] Set up database connection (PostgreSQL)
- [ ] Configure environment variables (.env)
- [ ] Set up database connection pooling
- [ ] Create repository layer structure
- [ ] Create service layer structure
- [ ] Create handler layer structure
- [ ] Set up HTTP router (Gin, Echo, or standard library)
- [ ] Configure CORS for frontend communication
- [ ] Set up logging system
- [ ] Configure error handling middleware
- [ ] Set up request validation

## Phase 3: Backend - Product Management

### Product Repository Layer
- [ ] Implement GetProducts() - List all products
- [ ] Implement CreateProduct() - Add new product
- [ ] Implement GetProductByID() - Get single product
- [ ] Implement UpdateProductStock() - Update stock (for transactions)
- [ ] Use parameterized queries (prevent SQL injection)
- [ ] Add proper error handling

### Product Service Layer
- [ ] Implement GetProducts() service method
- [ ] Implement CreateProduct() service method
- [ ] Add business logic validation
- [ ] Handle service-level errors

### Product Handler Layer
- [ ] Implement GET /products endpoint
  - [ ] Return: ID, Name, Stock, Price
  - [ ] Proper HTTP status codes
  - [ ] JSON response formatting
- [ ] Implement POST /products endpoint
  - [ ] Request body validation
  - [ ] Proper HTTP status codes
  - [ ] Error response formatting

## Phase 4: Backend - Order/Transaction Management

### Transaction Repository Layer
- [ ] Implement CreateOrder() with database transaction
- [ ] Implement GetProductWithLock() - Use SELECT FOR UPDATE (row-level locking)
- [ ] Implement UpdateStock() within transaction
- [ ] Implement SaveTransactionHistory()
- [ ] Ensure atomic operations (all or nothing)
- [ ] Handle transaction rollback on errors

### Transaction Service Layer
- [ ] Implement CreateOrder() service method
- [ ] Implement stock validation logic
  - [ ] Check if stock is sufficient
  - [ ] Return error if stock insufficient
- [ ] Implement price calculation logic
  - [ ] Calculate: (Price × Quantity) - (Discount %)
  - [ ] Handle edge cases (negative discount, etc.)
- [ ] Coordinate repository calls within transaction
- [ ] Handle race condition scenarios

### Transaction Handler Layer
- [ ] Implement POST /order endpoint
  - [ ] Request body validation (product_id, quantity, discount_percent)
  - [ ] Call service layer
  - [ ] Return success response with order details
  - [ ] Return error response if stock insufficient
  - [ ] Proper HTTP status codes (200, 400, 500)

### Race Condition Protection (CRITICAL)
- [ ] Implement database transaction with proper isolation level
- [ ] Use SELECT FOR UPDATE to lock product row
- [ ] Verify stock before and after lock
- [ ] Test with concurrent requests (10 simultaneous requests for 1 stock)
- [ ] Ensure only 1 request succeeds, 9 fail
- [ ] Verify stock never goes negative
- [ ] Test edge cases (zero stock, negative stock attempts)

## Phase 5: Frontend - Project Setup

### Next.js Configuration
- [ ] Set up Next.js project structure
- [ ] Configure API base URL (environment variable)
- [ ] Set up HTTP client (axios or fetch)
- [ ] Create API service layer
- [ ] Set up error handling utilities
- [ ] Set up toast/notification system (react-toastify or similar)

## Phase 6: Frontend - Dashboard

### Product List Component
- [ ] Create products table component
- [ ] Display: ID, Drug Name, Stock, Price
- [ ] Implement data fetching from GET /products
- [ ] Add loading state
- [ ] Add error handling
- [ ] Implement refresh functionality (manual or real-time)
- [ ] Style table appropriately

## Phase 7: Frontend - Order Form

### Order Form Component
- [ ] Create order form component
- [ ] Implement drug selection dropdown
  - [ ] Fetch products list
  - [ ] Populate dropdown options
- [ ] Implement quantity input field
  - [ ] Number validation
  - [ ] Minimum value validation (>= 1)
- [ ] Implement discount percentage input
  - [ ] Number validation
  - [ ] Range validation (0-100)
- [ ] Implement real-time price estimation
  - [ ] Calculate price as user types
  - [ ] Formula: (Price × Quantity) - (Discount %)
  - [ ] Update display immediately
- [ ] Add submit button
- [ ] Implement form validation
- [ ] Add loading state (disable button during submission)
- [ ] Prevent multiple submissions

### Order Submission
- [ ] Call POST /order API
- [ ] Handle success response
  - [ ] Show success message (toast)
  - [ ] Refresh product list
  - [ ] Reset form
- [ ] Handle error response
  - [ ] Display error message (toast/alert)
  - [ ] Show specific error (stock insufficient, etc.)
  - [ ] Keep form data for retry

## Phase 8: Testing

### Backend Testing
- [ ] Write unit tests for price calculation logic
- [ ] Write unit tests for stock validation logic
- [ ] Write integration tests for GET /products
- [ ] Write integration tests for POST /products
- [ ] Write integration tests for POST /order
- [ ] Write concurrent request tests (race condition)
  - [ ] Simulate 10 concurrent requests
  - [ ] Verify only 1 succeeds
  - [ ] Verify stock doesn't go negative
- [ ] Test edge cases
  - [ ] Zero stock
  - [ ] Negative quantity
  - [ ] Invalid discount percentage
  - [ ] Non-existent product_id

### Frontend Testing
- [ ] Test form validation
- [ ] Test real-time price calculation
- [ ] Test error handling and display
- [ ] Test loading states
- [ ] Manual testing of complete flow

## Phase 9: Code Quality & Best Practices

### Backend Code Quality
- [ ] Ensure Clean Architecture (Handler → Service → Repository)
- [ ] Verify business logic is in Service layer, not Handler
- [ ] Check all SQL queries use parameterized statements
- [ ] Verify proper error handling throughout
- [ ] Add meaningful comments where necessary
- [ ] Ensure consistent code formatting
- [ ] Verify variable naming is clear and descriptive

### Frontend Code Quality
- [ ] Organize components properly
- [ ] Ensure reusable components
- [ ] Verify proper state management
- [ ] Check loading states are implemented
- [ ] Verify error handling is user-friendly
- [ ] Ensure responsive design (if required)

## Phase 10: Documentation

### README.md
- [ ] Complete project overview
- [ ] Add technology stack details
- [ ] Document setup instructions (database & application)
- [ ] Add API documentation
  - [ ] List all endpoints
  - [ ] Request/response examples
  - [ ] Error responses
- [ ] Add link to Swagger/Postman Collection (if created)
- [ ] Add how to run instructions

### Code Documentation
- [ ] Add comments to complex logic
- [ ] Document API handlers
- [ ] Document service methods
- [ ] Document repository methods

### Additional Documentation
- [ ] Create Postman collection or Swagger documentation (optional but recommended)
- [ ] Document environment variables
- [ ] Document database schema

## Phase 11: Final Checks

### Evaluation Criteria Compliance
- [ ] ✅ Data Integrity: Stock never goes negative
- [ ] ✅ Data Integrity: Using DB Transaction & Locking correctly
- [ ] ✅ Code Structure: Clean Architecture / Layered
- [ ] ✅ Code Structure: Clean folder structure, clear variable naming
- [ ] ✅ Database: Table schema makes sense (separated tables)
- [ ] ✅ Database: Using migration tool
- [ ] ✅ Database: No hardcoded SQL (all parameterized)
- [ ] ✅ Frontend: Functional display & UX
- [ ] ✅ Frontend: Real-time validation (price changes)
- [ ] ✅ Frontend: Loading state implemented
- [ ] ✅ Nice to Have: Unit tests created

### Pre-Submission
- [ ] Test complete application flow
- [ ] Verify all requirements are met
- [ ] Check for any console errors
- [ ] Verify README.md is complete
- [ ] Prepare GitHub repository (public)
- [ ] Commit all code
- [ ] Push to GitHub
- [ ] Verify repository is accessible

## Notes
- Check off items as you complete them
- Priority: Security → Best Practices → Scalability → Structure → Instructions
- Focus on making the code stand out while following requirements
- Don't overdo it, but include quality features that demonstrate expertise
- Test race condition handling thoroughly - this is a critical requirement

