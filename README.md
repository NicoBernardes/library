# Library Management System

A full-stack library management application built with Go and the Gin web framework, featuring both a RESTful API and a web-based UI.

## Features

- **Book Management** - Full CRUD with title, author, and quantity tracking
- **User Management** - Register and manage library members with email validation
- **Loan Tracking** - Borrow and return books with automatic availability updates
- **Search & Filtering** - Search books by title/author, users by name/email, filter loans by status
- **Dashboard** - Overview with statistics (total books, users, active loans, etc.)

## Tech Stack

- **Language**: Go 1.25.5
- **Framework**: Gin v1.11.0
- **Storage**: In-memory (concurrent-safe with RWMutex)
- **Frontend**: Server-rendered HTML templates with embedded CSS
- **Architecture**: MVC with Repository and Service layers

## Project Structure

```
library/
├── cmd/api/
│   └── main.go                          # Entry point and route setup
├── internal/
│   ├── books/
│   │   ├── models/                      # Entity and interfaces
│   │   ├── repositories/               # Data access layer
│   │   ├── services/                    # Business logic
│   │   └── controllers/                # API handlers
│   ├── users/                           # Same layered structure
│   ├── loans/                           # Same layered structure
│   └── web/
│       └── controllers/                 # Web UI handlers
├── templates/                           # HTML templates
│   ├── layout.html                      # Master layout with navigation
│   ├── dashboard.html                   # Home page with stats
│   ├── books.html                       # Book management
│   ├── users.html                       # User management
│   ├── loans.html                       # Loan management
│   ├── modals.html                      # Reusable modal forms
│   └── styles.html                      # Embedded CSS
└── go.mod
```

## Getting Started

### Prerequisites

- Go 1.25.5 or higher

### Run

```bash
git clone <repository-url>
cd library
go mod download
go run ./cmd/api/main.go
```

The application starts at `http://localhost:8080`.

### Build

```bash
go build -o library ./cmd/api
```

## Web Interface

| Page | Path | Description |
|------|------|-------------|
| Dashboard | `/` | Statistics and quick actions |
| Books | `/books` | Manage book catalog |
| Users | `/users` | Manage library members |
| Loans | `/loans` | Track borrows and returns |

## API Endpoints

### Books (`/api/books`)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/books` | List all books |
| GET | `/api/books/:id` | Get a book |
| POST | `/api/books` | Create a book |
| PUT | `/api/books/:id` | Update a book |
| DELETE | `/api/books/:id` | Delete a book |

### Users (`/api/users`)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/users` | List all users |
| GET | `/api/users/:id` | Get a user |
| POST | `/api/users` | Create a user |
| PUT | `/api/users/:id` | Update a user |
| DELETE | `/api/users/:id` | Delete a user |

### Loans (`/api/loans`)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/loans` | List all loans |
| GET | `/api/loans/:id` | Get a loan |
| POST | `/api/loans` | Create a loan (borrow) |
| PUT | `/api/loans/:id/return` | Return a book |
| GET | `/api/loans/users/:userId/loans` | Get loans by user |

## Business Rules

- Books require a title and author (min 5 characters) and quantity (min 1)
- Users require a name and valid email
- A user cannot borrow a new book while they have an active loan
- Book quantity decreases on borrow and increases on return

## Architecture

```
HTTP Request -> Controllers -> Services -> Repositories -> In-Memory Store
```

Each domain (books, users, loans) follows interface-driven design with dependency injection, keeping layers decoupled and testable.

## License

MIT License
