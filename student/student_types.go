package student

type Student struct {
	Name   string `json:"name"`
	Grades []int  `json:"grades"`
}