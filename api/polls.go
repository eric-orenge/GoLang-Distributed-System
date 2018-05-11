package main

import (
	"errors"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type poll struct {
	ID      bson.ObjectId  `bson:"_id" json:"id"`
	Title   string         `json":"title""`
	Options []string       `json:"options"`
	Results map[string]int `json:"results,omitempty"`
}

func handlePolls(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlePollsGet(w, r)
		return
	case "POST":
		handlePollsPost(w, r)
		return
	case "DELETE":
		handlePollsDelete(w, r)
		return
	}
	// not found
	respondHTTPErr(w, r, http.StatusNotFound)
}
func handlePollsGet(w http.ResponseWriter, r *http.Request) {
	db := GetVar(r, "db").(*mgo.Database)
	c := db.C("polls")
	var q mgo.Query
	path := NewPath(r.URL.Path)
	if path.HasID() {
		q = c.FindId(bson.ObjectIdHex(path.ID))
	} else {
		q = c.Find(nil)
	}
	var results *[]poll

	respondErr(w, r, http.StatusInternalServerError, errors.New("not implemented"))
}
func handlePollsPost(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("not implemented"))
}
func handlePollsDelete(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("not implemented"))
}
