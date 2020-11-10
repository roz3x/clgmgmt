package main

import (
	"collegedbms/db"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func enroll(w http.ResponseWriter, r *http.Request) {
	s, err := DB.SelectEnroll(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	_s, _ := json.Marshal(s)
	fmt.Fprintf(w, "%s", _s)
}

func insertEnroll(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	_s := &db.InsertEnrollParams{}
	_ = json.Unmarshal(b, _s)
	if _, e := DB.InsertEnroll(context.Background(), *_s); e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println("student enrolled ", _s)
	}
}
