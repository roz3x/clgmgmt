package main

import "net/http"

func main() {
	http.HandleFunc("/create_department", insertDepartments)
	http.HandleFunc("/create_student/", insertStudent)
	http.HandleFunc("/students", students)
	http.Handle("/", http.FileServer(http.Dir("../frontend")))
	println("starting server")
	http.ListenAndServe(":8989", nil)
}
