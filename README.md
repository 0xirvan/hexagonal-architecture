# Go Hexagonal Architecture Template

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Echo](https://img.shields.io/badge/Echo-v4-00ADD8?style=flat)](https://echo.labstack.com/)
[![Wire](https://img.shields.io/badge/Wire-Dependency%20Injection-00ADD8?style=flat)](https://github.com/google/wire)

A Go application template with **Hexagonal Architecture** (Ports & Adapters)

## 🏗️ Architecture

This project uses **Hexagonal Architecture** (also known as Ports & Adapters Pattern) which separates business logic from external dependencies.

```
┌─────────────────────────────────────────────────────────┐
│                    Delivery Layer                       │
│         (HTTP Handlers, Routes, Middleware)             │
└───────────────────┬─────────────────────────────────────┘
                    │
┌───────────────────▼─────────────────────────────────────┐
│                   Core Domain                           │
│          (Business Logic, Use Cases, Entities)          │
└───────────────────┬─────────────────────────────────────┘
                    │
┌───────────────────▼─────────────────────────────────────┐
│                 Persistence Layer                       │
│          (Repositories, Database, Storage)              │
└─────────────────────────────────────────────────────────┘
```

### Layer Explanation

#### 1. **Core Domain** (`internal/core/`)

Contains business logic that is independent from frameworks and external dependencies.

- **Domain**: Entities and business rules
- **Port**: Interfaces for repositories and services
- **UseCase**: Business logic implementation

#### 2. **Adapter** (`internal/adapter/`)

Concrete implementations of ports that interact with the outside world.

- **Delivery**: HTTP handlers, routes, middleware
- **Persistence**: Database implementation (inmemory, sqlite, mysql, etc.)
- **Config**: Application configuration

#### 3. **App** (`internal/app/`)

Application composition using Wire for dependency injection.

## 📁 Struktur Project

```
server/
├── cmd/
│   └── server/
│       └── main.go                    # Application entry point
├── internal/
│   ├── adapter/                       # Adapter Layer
│   │   ├── config/                    # Configuration
│   │   │   ├── app.go
│   │   │   ├── config.go
│   │   │   └── http.go
│   │   ├── delivery/                  # Delivery adapters
│   │   │   └── http/
│   │   │       ├── router.go          # HTTP router setup
│   │   │       ├── todo_handler.go    # Todo HTTP handlers
│   │   │       ├── todo_routes.go     # Todo routes registration
│   │   │       ├── validator.go       # Request validation
│   │   │       ├── dto/               # Data Transfer Objects
│   │   │       ├── helper/            # HTTP helpers
│   │   │       └── middleware/        # HTTP middleware
│   │   └── persistence/               # Persistence adapters
│   │       ├── inmemory/              # In-memory repository
│   │       │   └── todo_repository.go
│   │       └── sqlite/                # SQLite repository (future)
│   ├── app/                           # Application composition
│   │   ├── http_app.go                # HTTP application
│   │   ├── wire.go                    # Wire providers
│   │   └── wire_gen.go                # Wire generated code
│   ├── core/                          # Core Domain Layer
│   │   ├── domain/                    # Domain entities
│   │   │   ├── errors.go              # Domain errors
│   │   │   └── todo.go                # Todo entity
│   │   ├── port/                      # Ports (interfaces)
│   │   │   ├── todo_repository.go     # Repository interface
│   │   │   └── todo_service.go        # Service interface
│   │   └── usecase/                   # Use cases
│   │       └── todo/
│   │           ├── service.go         # Todo service implementation
│   │           ├── create.go
│   │           ├── delete.go
│   │           ├── get.go
│   │           ├── list.go
│   │           ├── list_paginated.go
│   │           ├── mark_done.go
│   │           ├── mark_undone.go
│   │           └── update.go
│   └── shared/                        # Shared utilities
│       └── ptr/
│           └── ptr.go                 # Pointer helpers
├── go.mod                             # Go module definition
├── go.sum                             # Go dependencies
├── README.md                          # This file
└── .env.example                       # env example
```

### 🔄 Architecture Flow

```
Request → Handler → Service Interface → Use Case → Repository Interface → Storage
   ↓          ↓            ↓                 ↓              ↓                 ↓
 Echo     TodoHandler  TodoService      Service Impl   TodoRepository    InMemory
(Adapter)  (Adapter)    (Port)          (Core)          (Port)          (Adapter)
```

**Key Principle:** Adapters depend on Core, Core never depends on Adapters.

## 🛠️ Technologies

- **[Go 1.21+](https://golang.org/)** - Programming language
- **[Echo v4](https://echo.labstack.com/)** - High performance HTTP framework
- **[Wire](https://github.com/google/wire)** - Compile-time dependency injection
- **[Validator](https://github.com/go-playground/validator)** - Struct and field validation

## 📦 Prerequisites

- Go 1.21 or higher
- Wire (for dependency injection)

Install Wire:

```bash
go install github.com/google/wire/cmd/wire@latest
```

## 🚀 Setup & Installation

### 1. Clone Repository

```bash
git clone https://github.com/0xirvan/hexagonal-architecture.git
cd hexagonal-architecture/server
```

### 2. Update Module Name (Optional)

If you want to use your own module name:

```bash
# Update go.mod
go mod edit -module github.com/yourusername/yourproject/server

# Update all imports
find . -type f -name "*.go" -exec sed -i 's|github.com/0xirvan/hexagonal-architecture/server|github.com/yourusername/yourproject/server|g' {} +

# Tidy dependencies
go mod tidy

```

### 3. Install Dependencies

```bash
go mod download
```

### 4. Generate Wire Dependencies

```bash
wire ./internal/app
```

### 5. Configure Environment

```bash
cp .env.example .env
```

### 6. Run Application

```bash
go run cmd/server/main.go
```

Server will run on `http://localhost:8080`

### Patterns Used

- **Hexagonal Architecture** - Core isolation from external dependencies
- **Repository Pattern** - Data access abstraction
- **Service Layer Pattern** - Business logic encapsulation with interface
- **Dependency Injection** - Loose coupling via Wire
- **DTO Pattern** - Data transfer between layers
- **Port & Adapter Pattern** - Interface-based boundaries

## 🔄 Switching Persistence

To switch from in-memory to SQLite:

1. Implement SQLite repository in `internal/adapter/persistence/sqlite/`
2. Update Wire in `internal/app/wire.go`:

```go
var repositorySet = wire.NewSet(
    sqlite.NewTodoRepository,  // Change from inmemory
)
```

3. Regenerate Wire: `wire ./internal/app`

## 📚 Resources

- [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/)
- [Wire Dependency Injection](https://github.com/google/wire)
- [Echo Framework Guide](https://echo.labstack.com/guide/)

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## ‍💻 Author

**0xirvan**

- GitHub: [@0xirvan](https://github.com/0xirvan)

## 🙏 Acknowledgments

- Hexagonal Architecture concept by Alistair Cockburn
- Built with ❤️ using Go and Echo framework
- Special thanks to the Go community for excellent tools and libraries

⭐ If you find this project helpful, please consider giving it a star!
