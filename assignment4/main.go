package main

import (
	"fmt"
)

func calculateGrade(score int) string {
	switch {
	case score >= 90 && score <= 100:
		return "A"
	case score >= 80 && score <= 89:
		return "B"
	case score >= 70 && score <= 79:
		return "C"
	case score >= 60 && score <= 69:
		return "D"
	case score >= 50 && score <= 59:
		return "E"
	default:
		return "F"
	}
}

func main() {
	students := []struct {
		name  string
		score int
	}{
		{"A", 90},
		{"B", 48},
		{"C", 63},
		{"D", 75},
		{"E", 70},
	}

	fmt.Println("Team, Grade")

	for _, student := range students {
		grade := calculateGrade(student.score)
		fmt.Printf("%s, %s\n", student.name, grade)
	}
}
