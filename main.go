package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var classes = make([]Class, 0) // making a list of maps. The 0 is the initial size of the list

type Student struct {
	ID        string
	Name      string
	Age       uint8
	IsInClass bool
}

type Class struct {
	id      string
	name    string
	maxSize uint8
}

var db *sql.DB

func main() {
	dotenv := godotenv.Load(".env")
	if dotenv != nil {
		log.Fatal("Failed to load .env file")
	}

	// Capture Connection properties
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DBNAME"),
	}
	// Get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Welcome to Class Register")
	for {
		fmt.Println("0. Get student by id")
		fmt.Println("1. Create a new class")
		fmt.Println("2. Add a student to a class")
		fmt.Println("3. Remove a student from a class")
		fmt.Println("4. Print all students in a class")
		fmt.Println("5. Log start time of class")
		fmt.Println("6. Log end time of class")
		fmt.Println("7. Exit")
		fmt.Print("Please enter your choice: ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 0:
			fmt.Print("Enter student id: ")
			var id uint64
			fmt.Scanln(&id)
			student, err, _ := StudentByID(id)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Student found: ", student)

		case 1:
			fmt.Println("Create a new class")
			var className string
			var maxSize uint8
			fmt.Print("Please enter the name of the class: ")
			fmt.Scanln(&className)
			fmt.Print("Please enter the maximum size of class: ")
			fmt.Scanln(&maxSize)
			// check if class name already exists
			classExists, class_id := classNameExists(className)
			if classExists {
				fmt.Printf("\nClass name already exists. It has id = %v\n", class_id)
			} else {
				newClass, err := CreateClass(Class{
					name:    className,
					maxSize: maxSize,
				})
				if err != nil {
					fmt.Println("Something went wrong")
					continue
				}
				fmt.Printf("New class ID: %v\n", newClass)
			}

			fmt.Println()
			json.MarshalIndent(classes, "", "  ")
		case 2:
			fmt.Println("Add a student to a class")
			fmt.Print("Please enter the ID number of the student: ")
			var studentID uint64
			fmt.Scanln(&studentID)
			// Check if student exists
			_, err, studentExists := StudentByID(studentID)
			if err != nil {
				fmt.Printf("Something went wrong in this way:\n		%v\n", err)
				continue
			}
			if studentExists {
				fmt.Print("Enter class name: ")
				var className string
				fmt.Scan(&className)
				classExists,_ := classNameExists(className)
				if classExists {
					// Check if class has started
					fmt.Println("Checking if class has started.")
					if classHasStarted(className) {
						// Update the current class register
						fmt.Println("Adding student to class")
					} else {
						fmt.Println("Class has not yet started. Please start the class first.")
						continue
					}
				}
			} else {
				fmt.Println("Student ID does not exist. Please try again.")
				continue
			}
		case 3:
			fmt.Println("Remove a student from a class")
			fmt.Print("Enter name of class: ")
			var className string
			fmt.Scanln(&className)
			fmt.Print("Enter name of student: ")
			var studentName string
			fmt.Scanln(&studentName)
			RemoveStudentFromClass(studentName, className)

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
			classExists, _ := classNameExists(className)
			if classExists {
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
			fmt.Print("Enter name of class: ")
			var className string
			fmt.Scanln(&className)
			// check if class exists
			classExists, _ := classNameExists(className)
			if classExists {
				// check if class has already ended
				if classHasEnded(className) {
					fmt.Println("Class has already ended")
				} else {
					// log end time
					LogEndTime(className)
				}
			} else {
				fmt.Println("Class does not exist")
			}
		case 7:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid choice")
			continue
		}
	}
}
