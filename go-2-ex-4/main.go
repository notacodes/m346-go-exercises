package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Students []Student
	}
	// TODO: declare a map of modules being attended by multiple classes
	type ModuleClasses map[string][]Class

	moduleClasses := ModuleClasses{
		"M356": []Class{
			{
				Students: []Student{
					{FirstName: "Alfred", LastName: "Müller"},
					{FirstName: "Bob", LastName: "Andrews"},
				},
			},
		},
		"M357": []Class{
			{
				Students: []Student{
					{FirstName: "Levin", LastName: "Schmidt"},
					{FirstName: "Clara", LastName: "Schneider"},
				},
			},
		},
	}

	fmt.Println(moduleClasses)
}
