# Project Structure

```
Mini-Indobat_Inventory_System/
├── backend/                          # Go backend application
│   ├── cmd/
│   │   └── server/
│   │       └── main.go               # Application entry point
│   ├── internal/                     # Internal application code
│   │   ├── handler/                  # HTTP handlers (API endpoints)
│   │   │   ├── product_handler.go
│   │   │   └── order_handler.go
│   │   ├── service/                  # Business logic layer
│   │   │   ├── product_service.go
│   │   │   └── order_service.go
│   │   ├── repository/               # Data access layer
│   │   │   ├── product_repository.go
│   │   │   └── order_repository.go
│   │   ├── model/                    # Data models/entities
│   │   │   ├── product.go
│   │   │   └── order.go
│   │   └── config/                   # Configuration
│   │       └── config.go
│   ├── migrations/                   # Database migration files
│   │   ├── 000001_create_products.up.sql
│   │   ├── 000001_create_products.down.sql
│   │   ├── 000002_create_orders.up.sql
│   │   └── 000002_create_orders.down.sql
│   ├── go.mod                        # Go module file
│   ├── go.sum                        # Go dependencies checksum
│   └── .env.example                  # Environment variables example
│
├── frontend/                         # Next.js frontend application
│   ├── app/                          # Next.js App Router
│   │   ├── layout.tsx                # Root layout
│   │   ├── page.tsx                  # Dashboard page
│   │   └── api/                      # API routes (if needed)
│   ├── components/                   # React components
│   │   ├── ProductTable.tsx          # Product list table
│   │   ├── OrderForm.tsx             # Order form component
│   │   └── ui/                       # Reusable UI components
│   ├── lib/                          # Utility libraries
│   │   ├── api.ts                    # API client
│   │   └── utils.ts                  # Utility functions
│   ├── public/                       # Static assets
│   ├── package.json                  # Node.js dependencies
│   ├── next.config.js                # Next.js configuration
│   └── .env.local.example            # Frontend environment variables example
│
├── README.md                         # Project documentation
├── REQUIREMENTS.md                   # Project requirements
├── SETUP_CHECKLIST.md                # Setup instructions
├── IMPLEMENTATION_CHECKLIST.md       # Implementation tasks
├── PROJECT_STRUCTURE.md              # This file
└── .gitignore                        # Git ignore rules
```

## Directory Descriptions

### Backend (`backend/`)
- **cmd/server/**: Application entry point
- **internal/handler/**: HTTP request handlers (REST API endpoints)
- **internal/service/**: Business logic (price calculation, stock validation)
- **internal/repository/**: Database operations (queries, transactions)
- **internal/model/**: Data structures/models
- **internal/config/**: Configuration management
- **migrations/**: Database schema migration files

### Frontend (`frontend/`)
- **app/**: Next.js App Router pages and layouts
- **components/**: React components (ProductTable, OrderForm, etc.)
- **lib/**: API client and utility functions
- **public/**: Static files (images, icons, etc.)

## Architecture Layers

Following Clean Architecture principles:

1. **Handler Layer** → Receives HTTP requests, validates input, calls service
2. **Service Layer** → Contains business logic (calculations, validations)
3. **Repository Layer** → Handles database operations (queries, transactions)

This separation ensures:
- Business logic is not in handlers
- Database operations are isolated
- Code is testable and maintainable
- Easy to swap implementations

