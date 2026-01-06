# Mini-Indobat Inventory System - Requirements (Translated)

## Scenario
As a Fullstack Developer, you are asked to create a simple stock management system for a Pharmaceutical Wholesaler (PBF Farmasi). The system must be able to handle incoming and outgoing goods accurately, calculate total prices correctly, and handle concurrent requests (Concurrency Safe).

## Required Tech Stack
- **Backend**: Go - Framework is free choice
- **Frontend**: Next.js
- **Database**: PostgreSQL

## Feature Specifications

### Backend - REST API Features

#### Master Product
- **GET /products**: List products (ID, Drug Name, Stock, Price)
- **POST /products**: Add new product

#### Transaction (Order)
- **POST /order**: Perform purchase (reduce stock)
  - **Input Body**: `product_id`, `quantity`, `discount_percent`
  - **Logic**:
    - Check if stock is sufficient? If not, return error
    - Calculate Total Payment = (Price × Qty) - (Discount %)
    - Reduce stock in database
    - Save transaction history

### Challenge: Race Condition Prevention
Ensure the **POST /order** endpoint is safe from Race Conditions.

**Scenario**: If stock remaining is 1, and 10 concurrent requests come in to purchase that item, only 1 request should succeed. The other 9 must fail. Stock must not go negative.

**Hint**: Use Database Locking

### Frontend
Single Page Application (SPA) - Simple

#### Dashboard
- Display table of drug list with remaining stock in real-time (or manual refresh)

#### Order Form
- Dropdown to select drug
- Input Quantity
- Input Discount (%)
- Display "Estimated Price" before Submit button is pressed

#### Error Handling
- If stock is out or failed to get last stock, display error message (Toast/Alert) to user

## Evaluation Criteria

| Category | Expectation (Weight) | Plus Points (+) | Minus Points (-) |
|----------|---------------------|-----------------|------------------|
| **Data Integrity** | Stock must not go negative | Using DB Transaction & Locking correctly | Stock becomes negative or validation fails |
| **Code Structure** | Clean Architecture / Layered (Handler, Service, Repository) | Clean folder structure, clear variable naming | Business logic placed in Controller/Handler |
| **Database** | Table schema makes sense (Product Table & Transaction Table separated) | Using migration tool (Golang-Migrate/GORM AutoMigrate) | Hardcoded SQL queries vulnerable to SQL Injection |
| **Frontend** | Functional display & User Experience (UX), form validation works | Real-time validation (price changes when typing qty) | No loading state (button can be clicked multiple times) |
| **Nice To Have** (not mandatory) | Unit Testing | Creating simple Unit Tests for price calculation / stock logic | No instructions on how to run the application |

## Submission Instructions

1. Upload code to GitHub (Public)
2. Include README.md file containing:
   - How to setup database & application (Localhost)
   - Brief API documentation (or link to Swagger/Postman Collection)
3. Submit repository link by latest [Day/Date/Time]

