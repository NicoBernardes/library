# Library Management System

A RESTful API for managing library books, users, and loans built with Go and the Gin web framework.

## Overview

This is a library management system that allows you to:
- **Manage Books**: Create, read, update, and delete books in the library catalog
- **Manage Users**: Register and manage library users
- **Track Loans**: Monitor book loans and returns by users

The application follows a clean MVC (Model-View-Controller) architecture with clear separation of concerns using repositories, services, and controllers.

## Tech Stack

- **Language**: Go 1.25.5
- **Framework**: Gin (v1.11.0)
- **Architecture**: MVC with Repository pattern

## Project Structure

```
library/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── books/
│   │   ├── models/                 # Book entity and business logic
│   │   ├── repositories/           # Data access layer for books
│   │   ├── services/               # Business logic layer for books
│   │   └── controllers/            # HTTP handlers for books
│   ├── users/
│   │   ├── models/                 # User entity and business logic
│   │   ├── repositories/           # Data access layer for users
│   │   ├── services/               # Business logic layer for users
│   │   └── controllers/            # HTTP handlers for users
│   └── loans/
│       ├── models/                 # Loan entity and business logic
│       ├── repositories/           # Data access layer for loans
│       ├── services/               # Business logic layer for loans
│       └── controllers/            # HTTP handlers for loans
└── go.mod                          # Go module definition
```

## Architecture

The application uses a **layered architecture**:

1. **Controllers** - Handle HTTP requests and responses
2. **Services** - Contain business logic and orchestration
3. **Repositories** - Handle data persistence and retrieval
4. **Models** - Define data structures and domain entities

### Dependency Flow

```
HTTP Request → Controllers → Services → Repositories → Data
```

## Installation

### Prerequisites

- Go 1.25.5 or higher

### Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd library
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run ./cmd/api/main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Books
- `GET /books` - List all books
- `GET /books/:id` - Get a specific book
- `POST /books` - Create a new book
- `PUT /books/:id` - Update a book
- `DELETE /books/:id` - Delete a book

### Users
- `GET /users` - List all users
- `GET /users/:id` - Get a specific user
- `POST /users` - Create a new user
- `PUT /users/:id` - Update a user
- `DELETE /users/:id` - Delete a user

### Loans
- `GET /loans` - List all loans
- `GET /loans/:id` - Get a specific loan
- `POST /loans` - Create a new loan
- `PUT /loans/:id` - Update a loan
- `DELETE /loans/:id` - Delete a loan

## Key Features

- **Book Management**: Manage library inventory with title, author, and quantity tracking
- **User Management**: Register and maintain user profiles
- **Loan Tracking**: Track which users have borrowed which books and manage returns
- **Validation**: Input validation on required fields and constraints

## Development

### Running Tests

```bash
go test ./...
```

### Building for Production

```bash
go build -o library ./cmd/api
```

## Future Enhancements

- Database integration (currently in-memory)
- Authentication and authorization
- Loan due dates and overdue tracking
- Book availability checking
- User reservation system
- API documentation (Swagger/OpenAPI)

## License

🏆 MIT License