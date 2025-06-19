
# 🏦 Banking API in Golang

![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?logo=go)
![Gin](https://img.shields.io/badge/Gin_Framework-Yes-blue?logo=gin)
![License](https://img.shields.io/badge/License-MIT-green)
![Architecture](https://img.shields.io/badge/Architecture-Hexagonal%20%2F%20Clean-informational)

A robust **RESTful Banking API** built with **Golang** using **Hexagonal (Clean) Architecture**, ensuring separation of concerns, scalability, and easy maintainability.

---

## 🏛️ Architecture

This project follows the principles of **Hexagonal Architecture (Ports and Adapters)**, also known as **Clean Architecture**.

> ✅ **Core Business Logic is isolated from frameworks, databases, and external dependencies.**  
> ✅ **Adapters (HTTP, DB) communicate with the core domain through interfaces (ports).**

---

## 🚀 Features

- 👥 Customer management (CRUD)
- 🏦 Account creation
- 💸 Transaction processing (debit/credit)
- 🔐 JWT Authentication
- 🌐 RESTful API with JSON responses
- 🏛️ Clean & Hexagonal Architecture applied

---

## 🏗️ Project Structure

```plaintext
.
├── app/          # Application layer: HTTP Handlers, Middlewares, Router (Adapters)
├── domain/       # Domain layer: Entities, Business Rules, Interfaces (Ports)
├── dto/          # Data Transfer Objects (Input/Output Models)
├── main.go       # Application entry point (Composition root)
├── go.mod        # Go module definition
└── go.sum        # Go module checksums
```

> 📌 **`domain/` contains all business rules, completely independent of frameworks and external services.**

---

## 🔧 Technologies Used

- 🐹 **Golang**
- 🔥 **Gin** (HTTP Framework)
- 🔐 **JWT** (Authentication)
- 🗄️ **MySQL** (Database)
- 🏛️ **Hexagonal / Clean Architecture**

---

## 📦 Getting Started

### ✅ Prerequisites

- Go **>= 1.18**
- MySQL running locally or remotely

---

### 🚀 Clone the repository

```bash
git clone https://github.com/rondon23/banking-golang.git
cd banking-golang
```

---

### 🔗 Install dependencies

```bash
go mod tidy
```

---

### ⚙️ Configure Database

Update your database credentials in `main.go` (or wherever your DB connection is handled):

```go
dbUser := "root"
dbPass := "password"
dbHost := "127.0.0.1"
dbPort := "3306"
dbName := "banking"
```

---

### ▶️ Run the API

```bash
go run main.go
```

The API will be running at:

```
http://localhost:8000
```

---

## 📑 API Endpoints

### 👥 Customers

| Method | Endpoint                    | Description        |
|--------|------------------------------|--------------------|
| GET    | `/customers`                 | Get all customers  |
| GET    | `/customers/{customer_id}`   | Get customer by ID |

---

### 🏦 Accounts

| Method | Endpoint                                          | Description            |
|--------|---------------------------------------------------|------------------------|
| POST   | `/customers/{customer_id}/account`               | Create a new account   |

---

### 💰 Transactions

| Method | Endpoint                                                          | Description             |
|--------|---------------------------------------------------------------------|-------------------------|
| POST   | `/customers/{customer_id}/account/{account_id}`                   | Create a new transaction (debit/credit) |

---

### 🔐 Authentication

This API uses **JWT-based authentication**.

Include the token in the **Authorization** header:

```http
Authorization: Bearer <token>
```

---

## 🐳 Docker (Optional)

> 🚧 **Coming soon...**

---

## 🤝 Contributing

Contributions are welcome!  
Feel free to open issues or submit pull requests.

---

## 📜 License

This project is licensed under the **MIT License**.

---

## ❤️ Built with Go and passion for Clean Code & Hexagonal Architecture
