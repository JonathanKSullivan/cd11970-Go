package main

import "fmt"

type Student struct {
	id   int
	name string
}

type Classroom struct {
	id          int
	capacity    int
	subject     string
	studentList []Student
}

func main() {
	c1 := Classroom{
		id:       1,
		capacity: 30,
		subject:  "Math",
		studentList: []Student{
			Student{
				id:   1,
				name: "John Doe",
			},
			Student{
				id:   2,
				name: "Jane Doe",
			},
		},
	}

	c2 := new(Classroom)
	c2.id = 2
	c2.capacity = 30
	c2.subject = "Scientific Method"
	c2.studentList = []Student{
		Student{id: 1, name: "John Smith"},
		Student{id: 2, name: "Jane Smith"},
	}

	fmt.Println(c1)
	fmt.Println(*c2)
}
