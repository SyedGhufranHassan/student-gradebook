package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"student-gradebook/student"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swaggo/http-swagger"
	_ "student-gradebook/docs"
)

// @title Student Gradebook API
// @version 1.0
// @description A RESTful API to manage student grades using Go and Chi
// @host localhost:8080
// @BasePath /
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

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

// @Summary Add a new student
// @Tags student
// @Accept json
// @Produce json
// @Param student body student.Student true "Student object"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /student [post]
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

// @Summary Get average grade of a student
// @Tags student
// @Produce json
// @Param name path string true "Student name"
// @Success 200 {object} map[string]float64
// @Failure 404 {object} map[string]string
// @Router /student/{name}/average [get]
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

// @Summary List all students
// @Tags student
// @Produce json
// @Success 200 {object} map[string][]int
// @Router /students [get]
func listStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student.Gradebook)
}

// @Summary Delete a student
// @Tags student
// @Produce json
// @Param name path string true "Student name"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /student/{name}/delete [delete]
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := chi.URLParam(r, "name")
	if success := student.DeleteStudent(name); success {
		json.NewEncoder(w).Encode(map[string]string{"message": "Student deleted successfully"})
	} else {
		http.Error(w, `{"error": "Student not found"}`, http.StatusNotFound)
	}
}
