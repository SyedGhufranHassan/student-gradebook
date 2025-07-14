# 📚 Student Gradebook (Web API using Go + Chi)

A simple RESTful API built in Go using the Chi router to manage student records and grades. Users can add students, view their average grade, list all students, and delete student records—all via HTTP endpoints.

---

## 🚀 Features

- ✅ Add a student and their grades via POST request  
- 📊 Calculate and view a student's average grade  
- 📋 List all students with their grades  
- ❌ Delete a student record by name  
- 🌐 RESTful API with clean routing using Go-Chi  

---

## 🗂️ Project Structure

```
student-gradebook/
├── README.md
├── go.mod                 # Go module definition
├── main.go                # Web server with Chi HTTP routes
├── student/               # Gradebook logic and types
│   ├── student.go         # Core logic for grade management
│   └── student_types.go   # Student struct
└── utils/                 # (Optional) grade parser
    └── parse.go
```

---

## 📦 Installation & Running

### 1. Clone the Repository

```bash
git clone https://github.com/SyedGhufranHassan/student-gradebook.git
cd student-gradebook
```

### 2. Initialize Go Modules & Install Dependencies

```bash
go mod tidy
```

### 3. Run the Server

```bash
go run main.go
```

The server will start at:  
📍 `http://localhost:8080`

---

## 🌐 API Endpoints

| Method | Endpoint                      | Description                    |
|--------|-------------------------------|--------------------------------|
| POST   | `/student`                    | Add a new student              |
| GET    | `/student/{name}/average`     | Get average grade for student  |
| GET    | `/students`                   | List all students              |
| DELETE | `/student/{name}/delete`      | Delete a student               |

---

## 🧪 Example API Usage

### ➕ Add Student
```bash
curl -X POST http://localhost:8080/student \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice", "grades": [90, 85, 78]}' 
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

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Router:** [Chi](https://github.com/go-chi/chi)
- **Standard Libraries:** `net/http`, `encoding/json`, `fmt`, `strings`,`bufio`, `os`, etc.

---

## 📌 Notes

- Data is stored in memory using Go maps. All records are lost when the app restarts.
- Ideal for learning RESTful APIs and Go web development.

