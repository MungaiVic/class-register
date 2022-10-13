package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func CreateClass(class Class) (int64, error) {
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
	fmt.Println("Removing student from class")
}

func PrintStudentsInClass(className string) {
	fmt.Println("Students are printed here.")
}

func classNameExists(className string) (bool, string) {
	var class Class
	// var classID uint8

	row := db.QueryRow("SELECT * FROM class WHERE className = ?", className)
	// unmarshall the row object to Class
	if err := row.Scan(&class.id, &class.name, &class.maxSize); err != nil {
		if err == sql.ErrNoRows {
			// Do nothing if duplicate class name is not found
			return false, class.id
		}
	}
	// classID = class.id

	return true, class.id
}

func LogStartTime(className string) {
	// find class
	classExists, class_id := classNameExists(className)
	if classExists {
		//TODO: check if class has started
		fmt.Printf("Class Exists with id = %v\n", class_id)
		// Logging start time
		result, err := db.Exec("INSERT INTO classtime (classId, startTime) VALUES (?, ?)", class_id, time.Now())
		if err != nil {
			fmt.Printf("LogStartTime: %v", err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			fmt.Printf("Log Start time: %v", err)
		}
		fmt.Printf("\nClass %v has started. Entry ID = %v\n\n", className, id)
	} else {
		fmt.Println("Class doesn't exist")
	}
}

func classHasStarted(className string) bool {
	// find class
	classExists, _ := classNameExists(className)
	if classExists {
		// check if class has started
		fmt.Println("Checking")

	}
	return false
}

func classHasEnded(className string) bool {
	//find class
	classExists, _ := classNameExists(className)
	if classExists {
		// check if class has ended
		fmt.Println("Class exists")
	}
	return false
}

func LogEndTime(className string) {
	// find class
	classExists, class_id := classNameExists(className)
	if classExists {
		//TODO: check if class has started
		fmt.Printf("Class Exists with id = %v\n", class_id)
		// Logging end time
		result, err := db.Exec("UPDATE classtime SET endTime = ? WHERE classId = ?", time.Now(), class_id)
		if err != nil {
			fmt.Printf("LogEndTime: %v", err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			fmt.Printf("Log End time: %v", err)
		}
		fmt.Printf("\nClass %v has ended. Entry ID = %v\n\n", className, id)
	} else {
		fmt.Println("Class doesn't exist")
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
