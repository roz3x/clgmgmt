package main

import (
	"collegedbms/db"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func departments(w http.ResponseWriter, r *http.Request) {
	s, err := DB.SelectDepartment(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	_s, _ := json.Marshal(s)
	fmt.Fprintf(w, "%s", _s)
	fmt.Println(string(_s))
}

func insertDepartments(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	_s := &db.InsertDepartmentParams{}
	_ = json.Unmarshal(b, _s)
	if _, e := DB.InsertDepartment(context.Background(), *_s); e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println("department added ", _s)
	}
}
