package http

import (
	"encoding/json"
	"fmt"
	"ikomers-be/http/api"
	"net/http"

	"github.com/graphql-go/graphql"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		RequestString: query,
		Schema:        schema,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
		fmt.Println()
	}
	return result
}

func InitHTTPServer() {
	schema := api.NewGraphQLSchema()

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
