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
		w.Write([]byte("Welcome to the API of Student Gradebook."))
	})

	r.Post("/student", addStudent)
	r.Get("/students", listStudents)
	r.Get("/student/{name}/average", getAverage)
	r.Delete("/student/{name}/delete", deleteStudent)

	fmt.Println("Server is running on exactly http://localhost:8080")

	http.ListenAndServe(":8080", r)
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	var s student.Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	student.AddStudent(s.Name, s.Grades)
	w.Write([]byte("Student added"))
}

func getAverage(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	avg, found := student.GetAverageValue(name)
	if !found {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"average": avg})
}

func listStudents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(student.Gradebook)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	student.DeleteStudent(name)
	w.Write([]byte("Student deleted"))
}
