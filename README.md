# Class Register Console App

- This app is built as a beginner project in an effort to learn more about the Go programming language.
- This app allows one to add classes and students to those classes.
- This document will be updated as development continues.
- :rotating_light: **To run this program, it is assumed that you have a valid installation of Go** :rotating_light:

## Key things to note

- At present, the app only works in memory. This means that all data created will NOT persist when the app is stopped or interupted.
- The app is slowly being migrated to using MySQL database for persistence of data. One can:
    1. Find a student by ID
    2. Check if a class exists
    3. Create a class.
- Other features will be ported to storing data in a DB over time.

## App functionality

### How to run the app

-> change directory into the folder like so: `cd class-register`.
-> Next, run the command

```bash
go run .
```

#### OR

```bash
go run main.go helpers.go
```

- This will present you with a menu that looks like this:

    ```bash
    Welcome to Class Register
    1. Create a new class
    2. Add a student to a class
    3. Remove a student from a class
    4. Print all students in a class
    5. Log start time of class
    6. Log end time of class
    7. Exit
    ```
