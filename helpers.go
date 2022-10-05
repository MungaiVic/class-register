package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func CreateClass(class Class) (int64, error) {
	// Create a new class
	// newClass := Class{
	// 	id:   newClassId,
	// 	name: name,
	// }
	result, err := db.Exec("INSERT INTO class (className, maxSize) VALUES (?, ?)", &class.name, &class.maxSize)
	if err != nil {
		return 0, fmt.Errorf("CreateClass: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreateClass: %v", err)
	}
	return id, nil
}

func CreateStudent(name string, age uint8) Student {
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
			// classes[i].students = append(classes[i].students, student)
			student.IsInClass = true
		}
	}
}

func RemoveStudentFromClass(studentName, className string) {
	// find class
	if classNameExists(className) {
		// find student
		// 	for i := 0; i < len(classes); i++ {
		// 		if classes[i].name == className {
		// 			for j := 0; j < len(classes[i].students); j++ {
		// 				if classes[i].students[j].Name == studentName {
		// 					// remove student
		// 					classes[i].students = append(classes[i].students[:j], classes[i].students[j+1:]...)
		// 				}
		// 			}
		// 		}
		// 	}
	}
}

func PrintStudentsInClass(className string) {
	for i := 0; i < len(classes); i++ {
		if classes[i].name == className {
			// fmt.Println(classes[i].students)
		}
		// fmt.Println(classes[i].name, ", id: ", classes[i].id)
	}
}

func classNameExists(className string) bool {
	var class Class

	row := db.QueryRow("SELECT * FROM class WHERE className = ?", className)
	// unmarshall the row object to Class
	if err := row.Scan(&class.id, &class.name, &class.maxSize); err != nil {
		if err == sql.ErrNoRows {
			// Do nothing is duplicate class name is not found
			return false
		}
	}
	return true
}

func LogStartTime(className string) {
	// find class
	if classNameExists(className) {
		// check if class has started
		for i := 0; i < len(classes); i++ {
			if classes[i].name == className {
				// if classes[i].startTime == nil {
				// 	// log start time
				// 	currentTime := time.Now()
				// 	classes[i].startTime = &currentTime
				// }
			}
		}
	}
}

func classHasStarted(className string) bool {
	// find class
	if classNameExists(className) {
		// check if class has started
		fmt.Println("Checking")
	}
	return false
}

func classHasEnded(className string) bool {
	//find class
	if classNameExists(className) {
		// check if class has ended
		for i := 0; i < len(classes); i++ {
			if classes[i].name == className {
				// if classes[i].endTime != nil {
				// 	return true
				// }
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
				if classes == nil {
					// log end time
					// currentTime := time.Now()
					// classes[i].endTime = &currentTime
				}
			}
		}
	}
}

/*
This function is used in two ways.
*/
func StudentByID(id uint64) (Student, error, bool) {
	// a Student to hold the result
	var student Student
	// query the database for the Student with the specified ID
	row := db.QueryRow("SELECT * from student WHERE id = ?", id)
	// unmarshal the row object to Student
	if err := row.Scan(&student.ID, &student.Name, &student.Age, &student.IsInClass); err != nil {
		if err == sql.ErrNoRows {
			return student, fmt.Errorf("Student with ID %d not found", id), false
		}
	}
	return student, nil, true
}
