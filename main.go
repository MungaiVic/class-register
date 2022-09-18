package main

import (
	"encoding/json"
	"fmt"
	"time"
)

var classes = make([]Class, 0) // making a list of maps. The 0 is the initial size of the list
var students = make([]Student, 0)

type Student struct {
	id        string
	name      string
	age       int
	isInClass bool
}

type Class struct {
	id        string
	name      string
	students  []Student
	startTime *time.Time
	endTime   *time.Time
}

func main() {
	fmt.Println("Welcome to Class Register")
	for {
		fmt.Println("1. Create a new class")
		fmt.Println("2. Add a student to a class")
		fmt.Println("3. Remove a student from a class")
		fmt.Println("4. Print all students in a class")
		fmt.Println("5. Log start time of class")
		fmt.Println("6. Log end time of class")
		fmt.Print("Please enter your choice: ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("Create a new class")
			fmt.Print("Please enter the name of the class: ")
			var className string
			fmt.Scanln(&className)
			// check if class name already exists
			if classNameExists(className) {
				fmt.Println("Class name already exists")
			} else {
				CreateClass(className)
			}
			fmt.Println(classes)
			fmt.Println("Class created successfully")
			fmt.Println()
			// fmt.Println(classes)
			json.MarshalIndent(classes, "", "  ")
		case 2:
			fmt.Println("Add a student to a class")
			fmt.Print("Please enter the name of the student: ")
			var studentName string
			fmt.Scanln(&studentName)
			fmt.Print("Please enter the age of the student: ")
			var studentAge int
			fmt.Scanln(&studentAge)
			// Create Student
			studentData := CreateStudent(studentName, studentAge)
			// Add student to list of students
			students = append(students, studentData)
			fmt.Print("Please enter the name of the class: ")
			var className string
			fmt.Scanln(&className)
			AddStudentToClass(studentData, className)
		case 3:
			fmt.Println("Remove a student from a class")
		case 4:
			fmt.Println("Print all students in a class")
			var className string
			if len(classes) == 0 {
				fmt.Println("No classes have been created yet")
				fmt.Println()
			} else {
				fmt.Print("Please enter the name of the class: ")
				fmt.Scanln(&className)
				PrintStudentsInClass(className)
			}

		case 5:
			fmt.Println("Log start time of class")
			fmt.Print("Enter name of class: ")
			var className string
			fmt.Scanln(&className)
			// check if class exists
			if classNameExists(className) {
				// check if class has already started
				if classHasStarted(className) {
					fmt.Println("Class has already started")
				} else {
					// log start time
					LogStartTime(className)
				}
			} else {
				fmt.Println("Class does not exist")
			}

		case 6:
			fmt.Println("Log end time of class")
		case 7:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid choice")
			continue
		}
	}
}
