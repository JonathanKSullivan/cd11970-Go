package main

import "fmt"

func main() {
	courses := map[int]string{
		0: "Calculus",
		1: "Biology",
		2: "Chemistry",
		3: "Computer Science",
		4: "Communications",
		5: "English",
		6: "Cantonese",
	}

	for course_number, course_name := range courses {
		if course_name[0] == 'C' {
			fmt.Println(course_number, course_name)
		}
	}

	courses[3] = "Algorithms"
	courses[7] = "Spanish"
	delete(courses, 0)
	fmt.Println(courses)
}
