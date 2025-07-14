package utils

import (
	"fmt"
	"strconv"
	"strings"
)

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
