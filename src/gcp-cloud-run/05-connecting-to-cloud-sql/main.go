package main

import (
	"encoding/json"
	"fmt"
	"gcp-cloud-run/05-connecting-to-cloud-sql/infra/database"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
)

type user struct {
	UserId           int    `db:"user_id"`
	UserEmailAddress string `db:"user_email_address"`
}

var db *sqlx.DB

func init() {
	var err error
	db, err = database.Connect()
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := `select * from "USER"`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("db.Query: %v", err)
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var user user
		err := rows.Scan(&user)
		if err != nil {
			log.Printf("rows.Scan: %v", err)
			http.Error(w, "Error scanning database", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Printf("json encode: %v", err)
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("HTTP server is listening on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
