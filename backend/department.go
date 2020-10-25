package main

import (
	"collegedbms/db"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func createDepartment(w http.ResponseWriter, r *http.Request) {
	p := db.InsertDepartmentParams{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, err := DB.InsertDepartment(context.Background(), p)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(b)
	return
}
