package student
package student

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var Gradebook = make(map[string][]int)


func AddStudent(name string, grades []int) error {
	if name == "" {
		return errors.New("student name cannot be empty")
	}
	if len(grades) == 0 {
		return errors.New("grades cannot be empty")
	}
	for _, g := range grades {
		if g < 0 || g > 100 {
			return fmt.Errorf("grade %d is invalid; must be between 0 and 100", g)
		}
	}
	Gradebook[name] = grades
	return nil
}


func DeleteStudent(name string) bool {
	if _, exists := Gradebook[name]; exists {
		delete(Gradebook, name)
		return true
	}
	return false
}

// GetAverageValue returns average grade if student exists and has grades
func GetAverageValue(name string) (float64, bool) {
	if grades, exists := Gradebook[name]; exists {
		if len(grades) == 0 {
			return 0, false
		}
		sum := 0
		for _, g := range grades {
			sum += g
		}
		avg := float64(sum) / float64(len(grades))
		return avg, true
	}
	return 0, false
}


func ParseGrades(input string) ([]int, error) {
	parts := strings.Split(input, ",")
	var grades []int

	for _, part := range parts {
		gradeStr := strings.TrimSpace(part)
		grade, err := strconv.Atoi(gradeStr)
		if err != nil {
			return nil, fmt.Errorf("invalid grade: %s", gradeStr)
		}
		grades = append(grades, grade)
	}
	return grades, nil
}
