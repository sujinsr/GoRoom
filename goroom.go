package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHanlder(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("html/home.html")
		t.Execute(w, "")
	}
}

func main() {
	fmt.Println("Bachelor Room Management System")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHanlder).Methods("GET")

	fmt.Println("Application started on port 9090")
	http.ListenAndServe(":9090", router)

}
