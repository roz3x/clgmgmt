package main

import "net/http"

func main() {
	http.HandleFunc("/create_department", createDepartment)
	println("starting server")
	http.ListenAndServe(":8080", nil)
}
