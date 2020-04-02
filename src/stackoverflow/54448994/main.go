package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

type region struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type user struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Region region `json:"region"`
}

var aRegion = region{ID: 34, Name: "San jose"}
var aUser = user{ID: "34", Name: "something", Region: aRegion}

var regionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Region",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.Int},
		"name": &graphql.Field{Type: graphql.String},
	},
})

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"region": &graphql.Field{
			Type: regionType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				source, ok := p.Source.(user)
				fmt.Printf("%#v\n", source)
				fmt.Println("ok:", ok)
				return source.Region, nil
			},
		},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return aUser, nil
			},
		},
	},
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})
	fmt.Println("Now server is running on port 8080")
	fmt.Println("Access the web app via browser at 'http://localhost:8080'")
	http.ListenAndServe(":8080", nil)
}
