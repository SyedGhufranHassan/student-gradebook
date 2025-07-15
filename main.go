package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"student-gradebook/student"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Welcome to the API of Student Gradebook."}`))
	})

	r.Post("/student", addStudent)
	r.Get("/students", listStudents)
	r.Get("/student/{name}/average", getAverage)
	r.Delete("/student/{name}/delete", deleteStudent)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var s student.Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if err := student.AddStudent(s.Name, s.Grades); err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Student added successfully"})
}

func getAverage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := chi.URLParam(r, "name")
	avg, found := student.GetAverageValue(name)

	if !found {
		http.Error(w, `{"error": "Student not found or no grades"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]float64{"average": avg})
}

func listStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student.Gradebook)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := chi.URLParam(r, "name")
	if success := student.DeleteStudent(name); success {
		json.NewEncoder(w).Encode(map[string]string{"message": "Student deleted successfully"})
	} else {
		http.Error(w, `{"error": "Student not found"}`, http.StatusNotFound)
	}
}
