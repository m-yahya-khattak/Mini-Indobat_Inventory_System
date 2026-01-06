# Implementation Checklist

Use this checklist to track the implementation progress of the Mini-Indobat Inventory System.

**Priority Order for Implementation:**
1. Security (SQL injection prevention, input validation, race condition handling)
2. Best Practices (Clean architecture, proper error handling, code organization)
3. Scalability/Generic (Reusable code, proper abstractions)
4. Structure (Layered architecture: Handler → Service → Repository)
5. Sticking to Instructions (Follow requirements precisely, don't overdo it)

## Phase 1: Project Initialization
- [x] Create project structure (backend/ and frontend/ directories) ✅
- [x] Initialize Go module in backend ✅
- [x] Initialize Next.js project in frontend ✅
- [x] Set up .gitignore files ✅
- [ ] Initialize Git repository
- [ ] Create .env.example files for configuration reference

## Phase 2: Database Setup & Infrastructure

### Database Schema Design
- [x] Design Products table schema ✅
  - [x] id (primary key) ✅
  - [x] name (drug name) ✅
  - [x] stock (integer, not null, default 0) ✅
  - [x] price (decimal/numeric) ✅
  - [x] created_at, updated_at timestamps ✅
- [x] Design Transactions/Orders table schema ✅
  - [x] id (primary key) ✅
  - [x] product_id (foreign key) ✅
  - [x] quantity (integer) ✅
  - [x] discount_percent (decimal) ✅
  - [x] total_price (calculated) ✅
  - [x] created_at timestamp ✅
- [x] Create database migration files ✅
  - [x] Use golang-migrate ✅
  - [x] Ensure migrations are versioned and reversible ✅

### Backend Infrastructure
- [x] Set up database connection (PostgreSQL) ✅
- [x] Configure environment variables (.env) ✅
- [x] Set up database connection pooling ✅
- [x] Create repository layer structure ✅
- [x] Create service layer structure ✅
- [x] Create handler layer structure ✅
- [x] Set up HTTP router (Gin) ✅
- [x] Configure CORS for frontend communication ✅
- [x] Set up logging system ✅
- [x] Configure error handling middleware ✅
- [x] Set up request validation ✅

## Phase 3: Backend - Product Management

### Product Repository Layer
- [x] Implement GetProducts() - List all products ✅
- [x] Implement CreateProduct() - Add new product ✅
- [x] Implement GetProductByID() - Get single product ✅
- [x] Implement UpdateProductStock() - Update stock (for transactions) ✅
- [x] Use parameterized queries (prevent SQL injection) ✅
- [x] Add proper error handling ✅

### Product Service Layer
- [x] Implement GetProducts() service method ✅
- [x] Implement CreateProduct() service method ✅
- [x] Add business logic validation ✅
- [x] Handle service-level errors ✅

### Product Handler Layer
- [x] Implement GET /products endpoint ✅
  - [x] Return: ID, Name, Stock, Price ✅
  - [x] Proper HTTP status codes ✅
  - [x] JSON response formatting ✅
- [x] Implement POST /products endpoint ✅
  - [x] Request body validation ✅
  - [x] Proper HTTP status codes ✅
  - [x] Error response formatting ✅

## Phase 4: Backend - Order/Transaction Management

### Transaction Repository Layer
- [x] Implement CreateOrder() with database transaction ✅
- [x] Implement GetProductWithLock() - Use SELECT FOR UPDATE (row-level locking) ✅
- [x] Implement UpdateStock() within transaction ✅
- [x] Implement SaveTransactionHistory() ✅
- [x] Ensure atomic operations (all or nothing) ✅
- [x] Handle transaction rollback on errors ✅

### Transaction Service Layer
- [x] Implement CreateOrder() service method ✅
- [x] Implement stock validation logic ✅
  - [x] Check if stock is sufficient ✅
  - [x] Return error if stock insufficient ✅
- [x] Implement price calculation logic ✅
  - [x] Calculate: (Price × Quantity) - (Discount %) ✅
  - [x] Handle edge cases (negative discount, etc.) ✅
- [x] Coordinate repository calls within transaction ✅
- [x] Handle race condition scenarios ✅

### Transaction Handler Layer
- [x] Implement POST /order endpoint ✅
  - [x] Request body validation (product_id, quantity, discount_percent) ✅
  - [x] Call service layer ✅
  - [x] Return success response with order details ✅
  - [x] Return error response if stock insufficient ✅
  - [x] Proper HTTP status codes (200, 400, 500) ✅

### Race Condition Protection (CRITICAL)
- [x] Implement database transaction with proper isolation level ✅ (SERIALIZABLE)
- [x] Use SELECT FOR UPDATE to lock product row ✅
- [x] Verify stock before and after lock ✅
- [ ] Test with concurrent requests (10 simultaneous requests for 1 stock) ⏳ **PENDING TESTING**
- [ ] Ensure only 1 request succeeds, 9 fail ⏳ **PENDING TESTING**
- [x] Verify stock never goes negative ✅ (Database constraint + validation)
- [ ] Test edge cases (zero stock, negative stock attempts) ⏳ **PENDING TESTING**

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
- [x] Ensure Clean Architecture (Handler → Service → Repository) ✅
- [x] Verify business logic is in Service layer, not Handler ✅
- [x] Check all SQL queries use parameterized statements ✅
- [x] Verify proper error handling throughout ✅
- [x] Add meaningful comments where necessary ✅
- [x] Ensure consistent code formatting ✅
- [x] Verify variable naming is clear and descriptive ✅

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
- [x] ✅ Data Integrity: Stock never goes negative ✅ (DB constraint + validation)
- [x] ✅ Data Integrity: Using DB Transaction & Locking correctly ✅ (SERIALIZABLE + SELECT FOR UPDATE)
- [x] ✅ Code Structure: Clean Architecture / Layered ✅
- [x] ✅ Code Structure: Clean folder structure, clear variable naming ✅
- [x] ✅ Database: Table schema makes sense (separated tables) ✅
- [x] ✅ Database: Using migration tool ✅ (golang-migrate)
- [x] ✅ Database: No hardcoded SQL (all parameterized) ✅
- [ ] ✅ Frontend: Functional display & UX ⏳ **IN PROGRESS**
- [ ] ✅ Frontend: Real-time validation (price changes) ⏳ **IN PROGRESS**
- [ ] ✅ Frontend: Loading state implemented ⏳ **IN PROGRESS**
- [ ] ✅ Nice to Have: Unit tests created ⏳ **PENDING**

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

