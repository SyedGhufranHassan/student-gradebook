# 📚 Student Gradebook (Web API using Go + Chi + Swagger)

A simple RESTful API built in Go using the Chi router to manage student records and grades. Enhanced with Swagger documentation using **Swaggo**.

---

## 🚀 Features

- ✅ Add a student and their grades via POST request  
- 📊 Calculate and view a student's average grade  
- 📋 List all students with their grades  
- ❌ Delete a student record by name  
- 🌐 RESTful API with clean routing using Go-Chi  
- 📄 Swagger-based API documentation with interactive UI

---

## 🗂️ Project Structure

```
student-gradebook/
├── README.md
├── go.mod
├── go.sum
├── main.go                 # Web server and routes
├── docs/                  # Auto-generated Swagger docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── student/
│   ├── student.go         # Business logic and validation
│   └── student_types.go   # Student struct
└── student-gradebook
```

---

## 🧰 Tech Stack

- **Language**: Go (Golang)
- **Router**: [Chi](https://github.com/go-chi/chi)
- **Docs**: [Swaggo/Swagger](https://github.com/swaggo/swag)
- **Format**: JSON APIs

---

## 🔧 Setup & Installation

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/student-gradebook.git
cd student-gradebook
```

### 2. Initialize Go Modules

```bash
go mod tidy
```

### 3. Install Swag CLI (if not installed)

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 4. Generate Swagger Documentation

```bash
swag init
```

### 5. Run the Server

```bash
go run main.go
```

Server will start at:

```
📍 http://localhost:8080
```

---

## 🌐 Swagger API Docs

Visit Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

Here you can interact with the API (try requests, view schema, etc.).

---

## 🧪 Example API Usage (cURL)

### ➕ Add Student

```bash
curl -X POST http://localhost:8080/student \
  -H "Content-Type: application/json" \
  -d '{"name": "Ghufran", "grades": [90, 85, 78]}'
```

### 📊 Get Average Grade

```bash
curl http://localhost:8080/student/Ghufran/average
```

### 📋 List All Students

```bash
curl http://localhost:8080/students
```

### ❌ Delete Student

```bash
curl -X DELETE http://localhost:8080/student/Ghufran/delete
```

---

## 📌 Notes

- ⚠️ Data is stored in-memory. Restarting the app clears all student records.
- 🧪 Swagger documentation is auto-generated using Go code annotations.
- This is a lightweight project meant for learning REST APIs in Go.

