package main

import "net/http"

func main() {
	http.HandleFunc("/create_department", createDepartment)
	http.HandleFunc("/students", students)
	println("starting server")
	http.ListenAndServe(":8080", nil)
}
