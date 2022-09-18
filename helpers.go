package main

import (
	"fmt"
	"os/exec"

	"github.com/google/uuid"
)

func CreateClass(name string) {
	// Generate a new class id
	newUUID, _ := exec.Command("uuidgen").Output()
	newClassId := string(newUUID)

	// Create a new class
	newClass := Class{
		id:        newClassId,
		name:      name,
		students:  []Student{},
		startTime: nil,
		endTime:   nil,
	}
	classes = append(classes, newClass)

}

func CreateStudent(name string, age int) Student {
	newStudent := Student{
		id:        uuid.New().String(),
		name:      name,
		age:       age,
		isInClass: false,
	}
	return newStudent
}

func AddStudentToClass(student Student, className string) {
	for i := 0; i < len(classes); i++ {
		if classes[i].name == className {
			classes[i].students = append(classes[i].students, student)
		}
	}
}

func PrintStudentsInClass(className string) {
	for i := 0; i < len(classes); i++ {
		if classes[i].name == className {
			fmt.Println(classes[i].students)
		}
		// fmt.Println(classes[i].name, ", id: ", classes[i].id)
	}
}

func classNameExists(className string) bool {
	for i := 0; i < len(classes); i++ {
		if classes[i].name == className {
			return true
		}
	}
	return false
}
