# 🛒 Rest API for Food Ordering Service with Echo Framework in Golang and Clean Architecture

## 📋 Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Environment Variables](#environment-variables)
  - [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Authentication](#authentication)
- [Contributing](#contributing)
- [License](#license)

---

## 🌟 Introduction
This project is a **Food Ordering Service** built with **Golang**, using the **Echo framework** and adhering to **Clean Architecture** principles. The goal of this project is to provide a robust and scalable backend service for managing food orders, including features like user authentication, menu management, cart handling, order processing, and payments.

---

## ✨ Features
- **User Authentication**: Secure user registration and login using JWT tokens.
- **Menu Management**: Browse and manage different categories of food items.
- **Cart Functionality**: Add, update, or delete items in the cart.
- **Order Processing**: Place orders and handle the checkout process.
- **Payment Integration**: Manage payments securely.
- **Clean Architecture**: Ensures high modularity and separation of concerns.

---

## 🛠️ Tech Stack
- **Language**: Go (Golang)
- **Web Framework**: Echo
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **Architecture**: Clean Architecture
- **Package Management**: Go Modules

---

## 📂 Project Structure
The project is organized based on **Clean Architecture**, following the principles of separation of concerns and modular design.

```
.
├── cmd
│   └── main.go                # Entry point of the application
├── internal
│   ├── application            # Application layer
│   │   ├── usecase            
│   │   │   └── usecase.go     # usecase flow business logic
│   ├── config
│   │   └── config.go          # Configuration setup
│   ├── domain                 # Business logic interfaces and entities
│   │   ├── repository         
│   │   ├── usecase             
│   ├── infrastructure
│   │   └── halper       
│   │   └── repository       
│   ├── presentation           # Authentication and request handling
│   │   └── handler       
│   │   └── middleware       
├── pkg
│   ├── tools            
│   │   └── migrations         # database migrations
│   └── utils            
│   │   └── jwt.go             # JWT token generation & validation
│   │   └── utils.go           # Utility functions (ID generation, helpers)
└── .env                       # Environment variables
```

---

## 🚀 Getting Started

### 📋 Prerequisites
Ensure you have the following installed on your machine:
- [Go](https://golang.org/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/) (optional, for containerized deployment)

### 📦 Installation
Clone the repository:
```bash
git clone https://github.com/yourusername/food-ordering-service.git
cd food-ordering-service
```

Install dependencies:
```bash
go mod tidy
```

### ⚙️ Environment Variables
Create a `.env` file in the root directory and configure it as follows:

```
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=yourusername
DB_PASSWORD=yourpassword
DB_NAME=food_ordering
JWT_SECRET=your_jwt_secret
```

### ▶️ Running the Application
Run the application using the following command:
```bash
go run cmd/main.go
```

Access the service at `http://localhost:8080`.

---

## 📌 API Endpoints
Here is a list of the key endpoints available in the service:

### 🔑 Authentication
- **POST** `/api/v1/register` - Register a new user
- **POST** `/api/v1/login` - User login and get JWT token

### 📋 Menu Management
- **GET** `/api/v1/menu` - Get list of menu items
- **POST** `/api/v1/menu` - Add a new menu item (Admin only)

### 🛒 Cart
- **GET** `/api/v1/cart` - Get items in the cart
- **POST** `/api/v1/cart` - Add items to the cart
- **DELETE** `/api/v1/cart/:item_id` - Remove item from the cart

### 🛍️ Orders
- **POST** `/api/v1/order` - Place an order
- **GET** `/api/v1/order/:order_id` - Get order details

### 💳 Payments
- **POST** `/api/v1/payment` - Process a payment

---

## 🔒 Authentication
- The application uses **JWT (JSON Web Tokens)** for secure authentication.
- After a successful login, include the token in the `Authorization` header as:
  ```
  Authorization: Bearer <your_token>
  ```

---

## 🤝 Contributing
Contributions are welcome! Please open an issue or a pull request if you have ideas for improvements or bug fixes.

### Steps to Contribute:
1. Fork the repository.
2. Create a new feature branch:
   ```bash
   git checkout -b feature/your-feature
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add your feature"
   ```
4. Push to your branch:
   ```bash
   git push origin feature/your-feature
   ```
5. Open a pull request.

---

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## ✨ Acknowledgments
Special thanks to the Go, Echo, and Clean Architecture communities for their valuable resources and inspiration.

---

Feel free to modify this template as needed to better align with your project's details and requirements! and sorry if there are differences in viewpoints regarding clean architecture 🙏
