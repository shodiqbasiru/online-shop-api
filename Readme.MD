# Online Shop API

This is a RESTful API for an online shop, built with Go.

## Features

- Category management (Create, Read, Update, Delete)
- Product management (Create, Read, Update, Delete)
- User authentication (Register, Login)
- Customer management (Read, Update, Delete)
- Order management (Create, Read, Update)

## Technologies Used

- Go
- MySQL
- Go Playground Validator
- Go Co-op Gocron Scheduler
- HTTP Router
- JWT Authentication

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.22 or later)
- MySQL

### Installing

1. Clone the repository
    ```bash
    git clone https://github.com/shodiqbasiru/online-shop-api.git
    ```

2. Change into the project directory
    ```bash
    cd online-shop-api
    ```

3. Run the application
    ```bash
    go run cmd/main.go
    ```

## API Endpoints

### Categories
- `GET /api/categories` - Get all categories
- `GET /api/categories/:categoryId` - Get a category by ID
- `POST /api/categories` - Create a new category
- `PUT /api/categories/:categoryId` - Update a category
- `DELETE /api/categories/:categoryId` - Delete a category

### Products
- `GET /api/products` - Get all products
- `GET /api/products/:productId` - Get a product by ID
- `POST /api/products` - Create a new product
- `PUT /api/products/:productId` - Update a product
- `DELETE /api/products/:productId` - Delete a product

### Auth
- `POST /api/users/register` - Register a new user
- `POST /api/users/register-admin` - Register a new admin user
- `POST /api/users/login` - Login a user

### Customers
- `GET /api/customers` - Get all customers
- `GET /api/customers/:customerId` - Get a customer by ID
- `PUT /api/customers/:customerId` - Update a customer
- `DELETE /api/customers/:customerId` - Delete a customer

### Orders
- `POST /api/orders` - Create a new order
- `GET /api/orders/:orderId` - Get an order by ID
- `GET /api/orders/:orderId/status` - Update order status
- `GET /api/orders` - Get all orders

