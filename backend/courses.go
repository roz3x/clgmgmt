package main

import (
	"collegedbms/db"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func course(w http.ResponseWriter, r *http.Request) {
	s, err := DB.SelectStudents(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	_s, _ := json.Marshal(s)
	fmt.Fprintf(w, "%s", _s)
}

func insertCourses(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	_s := &db.InsertStudentParams{}
	_ = json.Unmarshal(b, _s)
	if _, e := DB.InsertStudent(context.Background(), *_s); e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println("student added ", _s)
	}
}
