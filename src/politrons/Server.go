package main

import (
	"net/http"
	"io/ioutil"
	"time"
	"os"
)

type VisitorAction interface {
	renderProfile(w http.ResponseWriter)
}

type Person struct {
	profileId string
}

// Go use composition instead of inheritance, now VisitorProfile just contains profileId
type VisitorProfile struct {
	Person
}

func handler(w http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	username := queryValues.Get("username")
	password := queryValues.Get("password")
	visitors := getVisitors(username, password)
	for profileId, _ := range visitors {
		visitor := VisitorProfile{Person{profileId}}
		actions := [...]VisitorAction{visitor}
		actions[0].renderProfile(w)
	}
	data, err := ioutil.ReadFile("end.html")
	checkError(err)
	w.Write(data)
}

// We implement a interface method just defining the method with same name and arguments
// and passing the class that implement as first argument
func (visitor VisitorProfile) renderProfile(w http.ResponseWriter) {
	createVisitorPage(visitor.profileId)
	visitor_page := visitor.profileId + ".html"
	data, err := ioutil.ReadFile(visitor_page)
	checkError(err)
	w.Write(data)
	time.Sleep(1 * time.Second)
	os.Remove(visitor_page)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}