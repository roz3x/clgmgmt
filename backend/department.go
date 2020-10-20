package main

import (
	"fmt"
	"net/http"
)

func createDepartment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	fmt.Println(name)
}
