package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.URL)
	fmt.Println(r.Form["url_long"])

	fmt.Fprintf(w, "Hello mrdulin!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, error := template.ParseFiles("login.html")
		if error != nil {
			fmt.Println("")
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		if len(r.Form["username"][0]) == 0 {
			fmt.Fprintf(w, "username is empty")
		}
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func selectFormController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/select", selectFormController)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
