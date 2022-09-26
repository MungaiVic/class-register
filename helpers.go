package main

import (
	"database/sql"
	"fmt"
	"os/exec"
	"time"

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
		ID:        uuid.New().String(),
		Name:      name,
		Age:       age,
		IsInClass: false,
	}
	return newStudent
}

func AddStudentToClass(student Student, className string) {
	for i := 0; i < len(classes); i++ {
		if classes[i].name == className {
			classes[i].students = append(classes[i].students, student)
			student.IsInClass = true
		}
	}
}

func RemoveStudentFromClass(studentName, className string) {
	// find class
	if classNameExists(className) {
		// find student
		for i := 0; i < len(classes); i++ {
			if classes[i].name == className {
				for j := 0; j < len(classes[i].students); j++ {
					if classes[i].students[j].Name == studentName {
						// remove student
						classes[i].students = append(classes[i].students[:j], classes[i].students[j+1:]...)
					}
				}
			}
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

func LogStartTime(className string) {
	// find class
	if classNameExists(className) {
		// check if class has started
		for i := 0; i < len(classes); i++ {
			if classes[i].name == className {
				if classes[i].startTime == nil {
					// log start time
					currentTime := time.Now()
					classes[i].startTime = &currentTime
				}
			}
		}
	}
}

func classHasStarted(className string) bool {
	// find class
	if classNameExists(className) {
		// check if class has started
		for i := 0; i < len(classes); i++ {
			if classes[i].name == className {
				if classes[i].startTime != nil {
					return true
				}
			}
		}
	}
	return false
}

func classHasEnded(className string) bool {
	//find class
	if classNameExists(className) {
		// check if class has ended
		for i := 0; i < len(classes); i++ {
			if classes[i].name == className {
				if classes[i].endTime != nil {
					return true
				}
			}
		}
	}
	return false
}

func LogEndTime(className string) {
	// find class
	if classNameExists(className) {
		// check if class has ended
		for i := 0; i < len(classes); i++ {
			if classes[i].name == className {
				if classes[i].endTime == nil {
					// log end time
					currentTime := time.Now()
					classes[i].endTime = &currentTime
				}
			}
		}
	}
}

// student queries for the Student with specified ID
func StudentByID(id int64) (Student, error) {
	// a Student to hold the result
	var s Student
	// query the database for the Student with the specified ID
	row := db.QueryRow("SELECT * from student WHERE id = ?", id)
	// unmarshal the row object to Student
	if err := row.Scan(&s.ID, &s.Name, &s.Age, &s.IsInClass); err != nil {
		if err == sql.ErrNoRows {
			return s, fmt.Errorf("Student with ID %d not found", id)
		}
	}
	return s, nil
}
