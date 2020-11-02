package main

import "net/http"

func main() {
	http.HandleFunc("/create_department", createDepartment)
	http.HandleFunc("/students", students)
	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/fs", fs)
	println("starting server")
	http.ListenAndServe(":8080", nil)
}
