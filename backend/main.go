package main

import "net/http"

func main() {
	http.HandleFunc("/create_student/", insertStudent)
	http.HandleFunc("/students", students)
	http.HandleFunc("/create_department", insertDepartments)
	http.HandleFunc("/departments", departments)
	http.HandleFunc("/create_enroll", insertEnroll)
	http.HandleFunc("/enrolls", enroll)
	http.HandleFunc("/create_instructor", insertInstructor)
	http.HandleFunc("/instructors", instructors)
	http.HandleFunc("/create_course", insertCourses)
	http.HandleFunc("/courses", course)

	http.Handle("/", http.FileServer(http.Dir("../frontend")))
	println("starting server")
	http.ListenAndServe(":8989", nil)
}
