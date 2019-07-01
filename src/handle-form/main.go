package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("r.ParseForm: %v", err)
		http.Error(w, "parse form error", http.StatusInternalServerError)
		return
	}
	fmt.Println(r.Form)
	fmt.Println(r.URL)
	fmt.Println(r.Form["url_long"])

	if _, err := fmt.Fprintf(w, "Hello mrdulin!"); err != nil {
		http.Error(w, "send data error", http.StatusInternalServerError)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.html")
		if err != nil {
			log.Printf("template.ParseFiles: %v", err)
			http.Error(w, "parse template file error", http.StatusInternalServerError)
			return
		}
		if err := t.Execute(w, nil); err != nil {
			log.Printf("t.Execute: %v", err)
			http.Error(w, "render template error", http.StatusInternalServerError)
		}
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
