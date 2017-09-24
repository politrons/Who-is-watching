package main

import (
	"net/http"
	"io/ioutil"
	"time"
	"os"
)

func handler(w http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	username := queryValues.Get("username")
	password := queryValues.Get("password")
	visitors := getVisitors(username, password)
	for visitor, _ := range visitors {
		createVisitorPage(visitor)
		visitor_page := visitor + ".html"
		data, err := ioutil.ReadFile(visitor_page)
		checkError(err)
		w.Write(data)
		time.Sleep(1 * time.Second)
		os.Remove(visitor_page)
	}
	data, err := ioutil.ReadFile("end.html")
	checkError(err)
	w.Write(data)
	//fmt.Fprintf(w, "Hi there, I love %s!", string(data))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}