package main

import (
	"collegedbms/db"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func instructors(w http.ResponseWriter, r *http.Request) {
	s, err := DB.SelectInstructor(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	_s, _ := json.Marshal(s)
	fmt.Fprintf(w, "%s", _s)
}

func insertInstructor(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	_s := &db.InsertInstructorParams{}
	_ = json.Unmarshal(b, _s)
	if _, e := DB.InsertInstructor(context.Background(), *_s); e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println("student added ", _s)
	}
}
