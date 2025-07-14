package student

import "fmt"

var Gradebook = make(map[string][]int)

func AddStudent(name string, grades []int) {
	Gradebook[name] = grades
	fmt.Println("Student added successfully!")
}

func DeleteStudent(name string) {
	if _, exists := Gradebook[name]; exists {
		delete(Gradebook, name)
		fmt.Println("Student deleted successfully!")
	} else {
		fmt.Println("Student not found.")
	}
}

func GetAverageValue(name string) (float64, bool) {
	if grades, exists := Gradebook[name]; exists {
		sum := 0
		for _, g := range grades {
			sum += g
		}
		avg := float64(sum) / float64(len(grades))
		return avg, true
	}
	return 0, false
}
