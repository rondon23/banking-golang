# 🏦 Banking API in Golang

A simple RESTful Banking API developed in Go. This project is designed to manage customers, accounts, and financial transactions for a fictional banking service.

---

## 🚀 Features

- ✅ Customer management (CRUD)
- ✅ Account creation
- ✅ Transaction processing (debit/credit)
- ✅ Authentication with JWT
- ✅ RESTful API with JSON responses
- ✅ Example of Hexagonal / Clean Architecture

---

## 🏗️ Project Structure

```plaintext
.
├── app/          # HTTP Handlers, Middlewares, Router
├── domain/       # Domain models and business logic
├── dto/          # Data Transfer Objects for requests and responses
├── main.go       # Entry point of the application
├── go.mod        # Go module definition
└── go.sum        # Go module checksums
```

---

## 🔧 Technologies Used

- 🐹 Go (Golang)
- 🔥 Gin (Web framework)
- 🔐 JWT (Authentication)
- 🗄️ MySQL (Database connection example)
- 🧠 Clean Architecture principles

---

## 📦 Installation and Run

### ✅ Prerequisites

- Go installed (version >= 1.18)
- MySQL running locally or configured with the correct credentials

---

### 🚀 Clone the repository

```bash
git clone https://github.com/your-username/banking-golang.git
cd banking-golang
🔗 Install dependencies
bash
Copy
Edit
go mod tidy
🏗️ Configure Database
Make sure to configure your MySQL credentials in main.go or inside app.go where the DB connection is set.

Example default connection:

go
Copy
Edit
dbUser := "root"
dbPass := "password"
dbHost := "127.0.0.1"
dbPort := "3306"
dbName := "banking"
▶️ Run the API
bash
Copy
Edit
go run main.go
The API will be available at:

arduino
Copy
Edit
http://localhost:8000
📑 API Endpoints
👥 Customers
Method	Endpoint	Description
GET	/customers	Get all customers
GET	/customers/{customer_id}	Get customer by ID

🏦 Accounts
Method	Endpoint	Description
POST	/customers/{customer_id}/account	Create a new account

💰 Transactions
Method	Endpoint	Description
POST	/customers/{customer_id}/account/{account_id}	Create a new transaction

🔐 Authentication
This API uses JWT-based authentication.

Include a valid token in the Authorization header:

makefile
Copy
Edit
Authorization: Bearer <token>
🐳 Build and Run with Docker (Optional)
Coming soon

🤝 Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

📜 License
This project is licensed under the MIT License.

Built with ❤️ in Go.
