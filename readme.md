# Task Manager API (Golang + Gin + GORM + PostgreSQL)

A simple task management API with user authentication and task CRUD operations. Built with Golang, Gin, GORM, and PostgreSQL. Uses JWT for secure user authentication and only allows authorized users to manage their own tasks.

---

## 🚀 Features

- **User Registration & Login**  
- **JWT Authentication (Bearer Token)**  
- **CRUD for Tasks (Only for Authenticated Users)**  
- **PostgreSQL Integration with GORM**  
- **Status Validation for Tasks ("pending" or "done")**  

---

## 🛠️ Installation

### 1. Clone the Repository  
```bash
git clone https://github.com/yourusername/task-manager-api.git
cd task-manager-api
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Set up environtment variable (.env)

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=taskmanager
JWT_SECRET=yoursecretkey
```

### 3. Run

```bash
go run main.go
```

Congratulation your application is run in `localhost:8080`

For documentaion about how to use the API, you can check `Task Manager.postman_collection.json` file

Let me know if you want me to tweak anything! 🚀✨ 