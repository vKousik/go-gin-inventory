🗃️ Inventory Management API (Go + Gin + PostgreSQL)

This project is a production-style REST API built using Go (Gin framework) following Clean Architecture principles.
It demonstrates database integration, layered architecture, concurrency, and proper error handling, as required in Week-5.

🚀 Features Implemented
✅ Clean Architecture

Domain – Entities & business errors

Repository – Database access (PostgreSQL via GORM)

Usecase – Business logic & validation

Delivery (HTTP) – Gin handlers & routing

Database package – Centralized DB connection

✅ Database Integration

PostgreSQL with proper PK / FK constraints

Tables:

products

categories

orders

UUIDs used as primary keys

Environment-based configuration using .env

✅ API Endpoints
Method	Endpoint	Description
GET	/products	Fetch all products
POST	/products	Create a product
GET	/categories	Fetch all categories
GET	/orders	Fetch all orders
GET	/tasks	All data from all tables
✅ Concurrency (Go Feature)

Uses a goroutine to simulate sending a notification email

Runs asynchronously without blocking the HTTP response

Dummy implementation as explicitly stated in the assignment


✅ Error Handling

Validation errors → 400 Bad Request

Not found errors → 404 Not Found

Infrastructure / DB errors → 500 Internal Server Error

Centralized domain-level errors for consistency

✅ CORS Middleware

Implemented using gin-contrib/cors

Handles cross-origin requests correctly


