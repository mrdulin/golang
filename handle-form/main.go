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
		tpl, err := template.ParseFiles("select-form.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if err := tpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), 500)
		}
	} else if r.Method == "POST" {

	}
}

func main() {
	const PORT int = 9090
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/select", selectFormController)
	log.Println(fmt.Sprintf("Staring http server on http://localhost:%d", PORT))
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
