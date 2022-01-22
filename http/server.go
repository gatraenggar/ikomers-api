package http

import (
	"encoding/json"
	"fmt"
	"ikomers-be/http/api"
	"net/http"

	"github.com/graphql-go/graphql"
)

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
		fmt.Println()
	}
	return result
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	schema := api.NewGraphQLSchema()
	result := ExecuteQuery(r.URL.Query().Get("query"), schema)
	json.NewEncoder(w).Encode(result)
}

func InitHTTPServer() {

	http.HandleFunc("/graphql", serverHandler)

	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
