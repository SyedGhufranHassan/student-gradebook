# 📚 Student Gradebook (Web API using Go + Chi)

A simple RESTful API built in Go using the Chi router to manage student records and grades. Users can add students, view their average grade, list all students, and delete student records—all via HTTP endpoints.

---

## 🚀 Features

- ✅ Add a student and their grades via POST request  
- 📊 Calculate and view a student's average grade  
- 📋 List all students with their grades  
- ❌ Delete a student record by name  
- 🌐 RESTful API with clean routing using Go-Chi  
- 🛡️ Proper validation for student name and grades  
- 🧠 Grade parsing function (`ParseGrades`) included in logic

---

## 🗂️ Project Structure

```
student-gradebook/
├── README.md
├── go.mod
├── main.go                 # Web server with Chi HTTP routes
├── reader.go              # Optional CLI helper (not used in API)
├── student/               
│   ├── student.go         # Gradebook logic + ParseGrades
│   └── student_types.go   # Student struct
```

---

## 📦 Installation & Running

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/student-gradebook.git
cd student-gradebook
```

### 2. Initialize Go Modules

```bash
go mod tidy
```

### 3. Run the Server

```bash
go run main.go
```

📍 The server will start at:  
http://localhost:8080

---

## 🌐 API Endpoints

| Method | Endpoint                      | Description                    |
|--------|-------------------------------|--------------------------------|
| POST   | `/student`                    | Add a new student              |
| GET    | `/student/{name}/average`     | Get average grade for student  |
| GET    | `/students`                   | List all students              |
| DELETE | `/student/{name}/delete`      | Delete a student               |

---

## 🧪 Example `curl` Commands

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

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Router:** [Chi](https://github.com/go-chi/chi)
- **Standard Libraries:** `net/http`, `encoding/json`, `fmt`, `strings`, `strconv`

---

## 📌 Notes

- Data is stored **in memory** using a Go map. All records are lost when the app restarts.
- This project is great for learning REST APIs, Go structs, slices, and validation.

---

