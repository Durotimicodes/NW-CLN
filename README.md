# NatWest Boxed Mettle Microservices API - Clone

This repository contains a set of microservices for the Mettle financial platform, developed using **Golang (Gin Gonic)** in a **monorepo** structure. The system is designed to support various financial services such as invoicing, expenses, accounting, tax management, savings, payments, lending (BNPL), user management, and customer support.

## Project Structure

```
mettle-microservices/
│── cmd/                         # Entry points for each service
│   ├── invoicing/
│   ├── expenses/
│   ├── accounting/
│   ├── tax/
│   ├── savings/
│   ├── payments/
│   ├── lending/                  # BNPL Service
│   ├── user-management/
│   ├── support/
│
│── services/                     # Core business logic for each microservice
│   ├── invoicing/
│   │   ├── api/                   # gRPC / REST Handlers
│   │   ├── service/               # Business logic
│   │   ├── repository/            # Data persistence layer
│   │   ├── models/                # Data models
│   │   ├── proto/                 # gRPC contracts (if used)
│   │   ├── config/                # Config files
│   │   ├── main.go                # Service entry point
│   │   └── Dockerfile
│   ├── payments/
│   ├── lending/                   # BNPL service
│   ├── user-management/
│
│── pkg/                          # Shared packages
│   ├── logger/                    # Logging utility
│   ├── middleware/                # Common HTTP middleware
│   ├── database/                  # PostgreSQL connection helpers
│   ├── queue/                     # Kafka / RabbitMQ event bus
│   ├── auth/                      # JWT, OAuth handlers
│
│── deployments/                   # Kubernetes manifests
│   ├── invoicing.yaml
│   ├── payments.yaml
│   ├── lending.yaml
│
│── configs/                       # Configuration files (per environment)
│   ├── dev/
│   ├── prod/
│
│── scripts/                       # DevOps and CI/CD scripts
│
│── Makefile                        # Build automation
│── docker-compose.yaml             # Local testing
│── README.md                       # Documentation
│── .github/                         # GitHub Actions for CI/CD
```

## Services & Endpoints

### **User Management Service**

- `POST /api/v1/users/register` - Register a new user
- `POST /api/v1/users/login` - Authenticate user & generate token
- `GET /api/v1/users/profile` - Get user profile
- `PUT /api/v1/users/profile` - Update user profile
- `DELETE /api/v1/users/:id` - Delete user (admin only)

### **Lending (BNPL) Service**

- `POST /api/v1/bnpl/apply` - Apply for BNPL
- `GET /api/v1/bnpl/:id` - Get BNPL details
- `PATCH /api/v1/bnpl/:id/repay` - Repay an installment

### **Payments Service**

- `POST /api/v1/payments/process` - Process a payment
- `GET /api/v1/payments/:id` - Get payment details

### **Expenses Service**

- `POST /api/v1/expenses` - Add an expense
- `GET /api/v1/expenses/:id` - Get expense details

### **Support Service**

- `POST /api/v1/support/ticket` - Create a support ticket
- `GET /api/v1/support/ticket/:id` - Get ticket details

## Tech Stack

- **Golang (Gin Gonic)** - Backend framework
- **PostgreSQL** - Database
- **Docker & Kubernetes** - Containerization and orchestration
- **gRPC & REST** - Communication between services
- **JWT Authentication** - Secure authentication

## Getting Started

### Prerequisites

- Install **Go**
- Install **Docker & Docker Compose**

### Running Locally

1. Clone the repository:
   ```sh
   git clone https://github.com/Durotimicodes/NW-CLN.git
   cd mettle-microservices
   ```
2. Start the services using Docker:
   ```sh
   docker-compose up --build
   ```
3. Access API documentation (if available):
   ```sh
   http://localhost:8080/docs
   ```

## Contribution

1. Fork the repository
2. Create a new branch (`feature/your-feature`)
3. Commit your changes
4. Create a pull request





